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

type CLIParsedValues struct {
	migrationsDirParsed string
}

func (p *CLIParser) Parse(c *CLIOptions) (*CLIParsedValues, string) {
	if len(strings.TrimSpace(c.migrationsDir)) == 0 {
		fmt.Println(c.migrationsDir)
		return nil, "Please specify a valid migrations dir"
	}

	migrationsDirStat, err := os.Stat(c.migrationsDir)

	if err != nil {
		return nil, "Please Specify a valid migrations dir"
	}

	if !migrationsDirStat.IsDir() {
		return nil, "Please specify a valid migrations dir"
	}

	cp := CLIParsedValues{}

	c.migrationsDir = path.Clean(c.migrationsDir)

	if err := os.Mkdir(c.migrationsDir, 0755); err != nil {
		return nil, "Something wrong happens, please try again later"
	}

	cp.migrationsDirParsed = c.migrationsDir

	return &cp, ""
}
