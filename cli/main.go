package cli

import (
	"flag"
	"os"
	"time"
)

var cliOptions = CLIOptions{}

func (c *CLIParser) Setup() *CLIOptions {

	timeerr := os.Setenv("TZ", "America/Sao_Paulo")

	if timeerr != nil {
		panic(timeerr.Error())
	}

	_, err := time.LoadLocation("America/Sao_Paulo")

	if err != nil {
		return nil
	}

	flag.BoolVar(&cliOptions.create, "create", false, "Creates a migration")
	flag.BoolVar(&cliOptions.history, "history", false, "Shows the entire operations history")
	flag.StringVar(&cliOptions.migrationsDir, "migrations-dir", "migrations", "Sets the migrations file path")
	flag.StringVar(&cliOptions.migrationName, "name", "migration", "Customize the default migrate name. The value will be truncated with the prefix.")

	flag.Parse()

	return &cliOptions
}
