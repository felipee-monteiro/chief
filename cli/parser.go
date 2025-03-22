package cli

import (
	"chief/utils"
	"os"
	"path"
	"time"
)

type CLIParser struct {
	Strict bool
}

type CLIOptions struct {
	create          bool
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
}

func (p *CLIParser) ParseAndCreateBaseDir(migrationsDir, migrationName string) (bool, string) {
	if !utils.IsValidString(migrationsDir) {
		return false, "Please specify a valid migrations dir"
	}

	if !utils.IsValidString(migrationName) {
		return false, "Please specify a valid migration name"
	}

	baseDir := path.Clean(migrationsDir + "/" + time.Now().Format(time.RFC3339Nano) + "_" + migrationName)

	if _, err := os.Stat(baseDir); err != nil {
		if os.IsExist(err) {
			return true, baseDir
		}

		if os.IsNotExist(err) {
			if err := os.MkdirAll(baseDir, 0o755); err != nil {
				return false, err.Error()
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

func (p *CLIParser) Parse(c *CLIOptions) {
	if c.create {
		ok, message := p.ParseAndCreateBaseDir(c.migrationsDir, c.migrationName)

		if !ok {
			panic(message)
		}
	}
}
