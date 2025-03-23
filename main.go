package main

import (
	"chief/cli"
	"os"
)

func main() {
	// db := db.DatabaseConnection{}
	// db.Connect("sa", "Epilefac57#$!$24042002", "host.docker.internal", "sigma", 1433)

	ci := cli.CLIParser{Strict: true}
	ci.Parse(ci.Setup())

	os.Exit(0)
}
