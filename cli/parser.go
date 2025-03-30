package cli

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/felipee-monteiro/chief/db"
	"github.com/felipee-monteiro/chief/utils"
)

type CLIParser struct {
	Strict bool
}

type CLIOptions struct {
	create           bool
	migrate          bool
	migrationsDir    string
	migrationName    string
	history          bool
	datatabseOptions struct {
		host     string
		port     int64
		user     string
		database string
		password string
	}
}

var connection = db.Connect()

// ParseAndCreateBaseDir parses the migrations dir and migration name from the CLI args, then creates a new migration base dir
// if it doesn't exist. It creates the up.sql and down.sql files in the new dir.
//
// If the migrations dir is invalid, it returns false and an error message.
// If the migration name is invalid, it returns false and an error message.
// If the base dir already exists, it returns true and the base dir path.
// If the base dir does not exist, it creates it and returns true and the base dir path.
func (p *CLIParser) ParseAndCreateBaseDir(migrationsDir, migrationName string) (bool, string) {
	if !utils.IsValidString(migrationsDir) {
		return false, "Please specify a valid migrations dir"
	}

	if !utils.IsValidString(migrationName) {
		return false, "Please specify a valid migration name"
	}

	baseDir := path.Clean(migrationsDir + "/" + time.Now().Format(time.RFC3339Nano) + "_" + migrationName)

	if _, e := os.Stat(baseDir); e != nil {
		if os.IsExist(e) {
			return true, baseDir
		}

		if os.IsNotExist(e) {
			if e := os.MkdirAll(baseDir, 0o755); e != nil {
				return false, e.Error()
			}

			if _, e := os.Create(baseDir + "/up.sql"); e != nil {
				return false, e.Error()
			}

			if _, e := os.Create(baseDir + "/down.sql"); e != nil {
				return false, e.Error()
			}
		}
	}

	return true, baseDir
}

// ExecuteMigration executes a SQL migration script using the "sqlcmd" utility.
// It takes the path to the SQL file as an argument and attempts to execute it
// on the "sigma" database using the specified connection parameters.
// If "sqlcmd" is not installed, the function will print an error message and exit.
// It also captures and logs any errors or output from the execution process.
func (p *CLIParser) ExecuteMigration(path string, c *CLIOptions) {

	filename := strings.Split(path, "/")[1]
	migrationName := strings.Split(filename, "_")[1]

	if ok, _ := db.IsExecuted(connection, migrationName, c.datatabseOptions.database); !ok {
		otp := exec.Command("sqlcmd", "-S", c.datatabseOptions.host, "-d", c.datatabseOptions.database, "-U", c.datatabseOptions.user, "-P", c.datatabseOptions.password, "-i", path, "-C")

		fmt.Println("Executing " + path + "...")

		stderr, err := otp.StderrPipe()

		if err != nil {
			log.Fatal(err)
		}

		if err := otp.Start(); err != nil {
			log.Fatal(err)
		}

		scanner := bufio.NewScanner(stderr)

		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

		if err := otp.Wait(); err != nil {
			log.Fatal(err)
		}

		db.CreateMigration(connection, migrationName, path+"/up.sql", path+"/down.sql", c.datatabseOptions.database)

		fmt.Println("Migration executed successfully")
	}
}

// Execute traverses the given base directory to find and execute "up.sql"
// migration files. It uses the ExecuteMigration method to perform the SQL
// execution. If any error occurs during the directory traversal or execution,
// it returns false and the error message. On success, it returns true and an
// empty string.
func (p *CLIParser) Execute(baseDir string, c *CLIOptions) (bool, string) {
	if _, err := exec.LookPath("sqlcmd"); err != nil {
		return false, "The \"sqlcmd\" utility MUST be installed and available in $PATH"
	}

	db.Migrate(connection)

	err := filepath.WalkDir(baseDir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !d.IsDir() {
			if d.Name() == "up.sql" {
				p.ExecuteMigration(path, c)
			}
		}

		return nil
	})

	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return false, "migration not found. did you create it?"
		}

		return false, err.Error()
	}

	db.Close(connection)

	return true, ""
}

// Parse is responsible for parsing the CLI options and executing the desired action.
// If the -create flag is specified, it will attempt to create a new migration base directory.
// If the -migrate flag is specified, it will execute all migrations in the specified directory.
// If the migration name is specified, it will execute the migration with the given name individually.
func (p *CLIParser) Parse(c *CLIOptions) {
	if c.create {
		ok, message := p.ParseAndCreateBaseDir(c.migrationsDir, c.migrationName)

		if !ok && utils.IsValidString(message) {
			os.Exit(1)
			panic(message)
		}

		os.Exit(0)

		return
	}

	if c.migrate {
		if utils.IsValidString(c.migrationName) && c.migrationName != "migration" {
			p.ExecuteMigration(path.Clean(c.migrationsDir+"/"+c.migrationName+"/up.sql"), c)
			os.Exit(0)
			return
		}

		if !utils.IsValidString(c.datatabseOptions.database) {
			fmt.Println("Database name is required")
			os.Exit(1)
			return
		}

		ok, message := p.Execute(path.Clean(c.migrationsDir), c)

		if !ok && utils.IsValidString(message) {
			log.Fatal(message)
		}
	}
}
