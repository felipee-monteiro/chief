package cli

import (
	"os"
	"path"
	"strings"
	"time"
)

type CLIParser struct {
	Strict bool
}

type CLIOptions struct {
	create          bool
	migrationsDir   string
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

func (p *CLIParser) ParseAndCreateBaseDir(migrationsDir string) (bool, string) {
	if len(strings.TrimSpace(migrationsDir)) == 0 {
		return false, "Please specify a valid migrations dir"
	}

	baseDir := migrationsDir + "/" + time.Now().String()
	_, err := os.Stat(baseDir)

	if err != nil {
		if os.IsExist(err) {
			return true, path.Clean(baseDir)
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

	return true, ""
}

func (p *CLIParser) Parse(c *CLIOptions) {
	if c.create {
		ok, message := p.ParseAndCreateBaseDir(c.migrationsDir)

		if !ok {
			panic(message)
		}
	}
}
