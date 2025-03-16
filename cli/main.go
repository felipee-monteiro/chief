package cli

import (
	"flag"
)

var cliOptions = CLIOptions{}

func (c *CLIParser) Setup() *CLIOptions {

	flag.BoolVar(&cliOptions.create, "create", false, "Creates a migration")
	flag.BoolVar(&cliOptions.history, "history", false, "Shows the entire operations history")
	flag.StringVar(&cliOptions.migrationsDir, "migrations-dir", "migrations", "Sets the migrations file path")

	flag.Parse()

	return &cliOptions
}
