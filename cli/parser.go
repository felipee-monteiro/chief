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

type FSCache struct {
	BaseDirStat os.FileInfo
}

func (p *CLIParser) ParseAndCreateBaseDir(migrationsDir string) (bool, string) {
	if len(strings.TrimSpace(migrationsDir)) == 0 {
		return false, "Please specify a valid migrations dir"
	}

	baseDir := migrationsDir + "/" + time.Now().String()
	baseDirStat, err := os.Stat(baseDir)

	_ = FSCache{BaseDirStat: baseDirStat}

	if err != nil {
		if os.IsExist(err) {
			return true, path.Clean(baseDir)
		}

		if os.IsNotExist(err) {

			if err := os.MkdirAll(baseDir, 0o755); err != nil {
				return false, "Theres some errors while trying to create the migrations dir. Please check it or try again"
			}
		}

		return false, "Something went wrong. Try again later"
	}

	upFile, err := os.Create(path.Clean(baseDir + "up.sql"))

	if err != nil {
		panic("Error while trying to create the migration files")
	}

	downFile, err := os.Create(path.Clean(baseDir + "down.sql"))

	if err != nil {
		panic("Error while trying to create the migration files")
	}

	upFile.Close()
	downFile.Close()

	return true, path.Clean(baseDir)
}

func (p *CLIParser) Parse(c *CLIOptions) {
	if c.create {
		ok, _ := p.ParseAndCreateBaseDir(c.migrationsDir)

		if !ok {
			panic("Something wrong happened while trying to create the migration")
		}
	}
}
