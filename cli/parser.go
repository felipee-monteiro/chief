package cli

import (
	"fmt"
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

type FSCache struct {
	fileInfo os.FileInfo
}

type CLIParsedValues struct {
	migrationsDirParsed string
}

func (p *CLIParser) ParseAndCreateDir(migrationsDir string) (*CLIParsedValues, string) {
	if len(strings.TrimSpace(migrationsDir)) == 0 {
		fmt.Println(migrationsDir)
		return nil, "Please specify a valid migrations dir"
	}

	migrationsPathStats, err := os.Stat(migrationsDir)

	if err != nil {
		if os.IsExist(err) {
			return nil, ""
		}

		if os.IsNotExist(err) {
			if err := os.Mkdir(migrationsDir, 0755); err != nil {
				return nil, "Theres some errors while trying to create the migrations dir. Please check it or try again"
			}
		}

		return nil, "Something went wrong. Try again later"
	}

	cp := CLIParsedValues{}

	migrationsDir = path.Clean(migrationsDir)
	fsCache := FSCache{}
	fsCache.fileInfo = migrationsPathStats

	cp.migrationsDirParsed = migrationsDir

	return &cp, ""
}

func (p *CLIParser) Parse(c *CLIOptions) *CLIParsedValues {

	values, err := p.ParseAndCreateDir(c.migrationsDir)

	if len(err) > 0 {
		return nil
	}

	return values
}
