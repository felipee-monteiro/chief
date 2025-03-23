package cli

import (
	"chief/utils"
	"fmt"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"time"
)

type CLIParser struct {
	Strict bool
}

type CLIOptions struct {
	create          bool
	migrate         bool
	migrationsDir   string
	migrationName   string
	history         bool
	databaseOptions struct {
		server   string
		user     string
		password string
		port     int32
	}
}

type CLIParsedValues struct {
	migrationsDirParsed string
	fsys                fs.FS
	migrationsFiles     []fs.DirEntry
}

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

func (p *CLIParser) Execute(baseDir string) (bool, string) {
	err := filepath.WalkDir(baseDir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !d.IsDir() {
			if d.Name() == "up.sql" {
				fmt.Println("Arquivo de up encontrado")
			}
		}

		return nil
	})

	if err != nil {
		return false, err.Error()
	}

	return true, ""
}

func (p *CLIParser) Parse(c *CLIOptions) {
	if c.create {
		ok, message := p.ParseAndCreateBaseDir(c.migrationsDir, c.migrationName)

		if !ok {
			os.Exit(1)
			panic(message)
		}

		os.Exit(0)

		return
	}

	if c.migrate {
		if utils.IsValidString(c.migrationName) {
			// TODO: implementar execução de migration individual
		}

		ok, message := p.Execute(path.Clean(c.migrationsDir))

		if !ok && utils.IsValidString(message) {
			panic(message)
		}
	}
}
