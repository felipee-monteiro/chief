package cli

import (
	"os"
	"path"
	"strings"
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

func (p *CLIParser) ParseAndCreateBaseDir(migrationsDir string) (*CLIParsedValues, string) {
	if len(strings.TrimSpace(migrationsDir)) == 0 {
		return nil, "Please specify a valid migrations dir"
	}

	_, err := os.Stat(migrationsDir)

	if err != nil {
		if os.IsExist(err) {
			return nil, ""
		}

		if os.IsNotExist(err) {
			if err := os.Mkdir(migrationsDir, 0o755); err != nil {
				return nil, "Theres some errors while trying to create the migrations dir. Please check it or try again"
			}
		}

		return nil, "Something went wrong. Try again later"
	}

	return &CLIParsedValues{
		migrationsDirParsed: path.Clean(migrationsDir),
	}, ""
}

func (p *CLIParser) Parse(c *CLIOptions) *CLIParsedValues {

	values, err := p.ParseAndCreateBaseDir(c.migrationsDir)

	if len(err) > 0 {
		return nil
	}

	return values
}
