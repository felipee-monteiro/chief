package cli

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

var cliOptions = CLIOptions{}

// Setup configures the CLI flags and sets the time zone to "America/Sao_Paulo".
// It returns a *CLIOptions object containing the parsed flags.
func (c *CLIParser) Setup() *CLIOptions {

	timeerr := os.Setenv("TZ", "America/Sao_Paulo")

	if timeerr != nil {
		os.Exit(1)
		panic(timeerr.Error())
	}

	if _, err := time.LoadLocation("America/Sao_Paulo"); err != nil {
		log.Fatal("Cannot set location env variable")
		return nil
	}

	flag.BoolVar(&cliOptions.create, "create", false, "Creates a migration")
	flag.BoolVar(&cliOptions.migrate, "migrate", false, "Execute all migrations")
	flag.BoolVar(&cliOptions.history, "history", false, "Shows the entire operations history")
	flag.StringVar(&cliOptions.migrationsDir, "migrations-dir", "migrations", "Sets the migrations dir path")
	flag.StringVar(&cliOptions.migrationName, "name", "migration", "Customize the default migrate name. The value will be truncated with the prefix.")
	flag.StringVar(&cliOptions.datatabseOptions.host, "host", "localhost", "Sets the database host")
	flag.StringVar(&cliOptions.datatabseOptions.database, "database", "", "Sets the database name")
	flag.Int64Var(&cliOptions.datatabseOptions.port, "port", 1433, "Sets the database port")
	flag.StringVar(&cliOptions.datatabseOptions.user, "user", os.Getenv("SQLCMDUSER"), "Sets the database user (https://learn.microsoft.com/en-us/sql/tools/sqlcmd/sqlcmd-use-scripting-variables?view=sql-server-ver16#sqlcmd-scripting-variables)")
	flag.StringVar(&cliOptions.datatabseOptions.password, "password", os.Getenv("SQLCMDPASSWORD"), "Sets the database password. [WARNING] If the password contains special chars, use single quotes (https://learn.microsoft.com/en-us/sql/tools/sqlcmd/sqlcmd-use-scripting-variables?view=sql-server-ver16#sqlcmd-scripting-variables)")

	flag.Parse()

	if len(flag.Args()) > 0 {
		for _, arg := range flag.Args() {
			fmt.Println("Unknown argument: " + arg)
		}
		os.Exit(1)
	}

	return &cliOptions
}
