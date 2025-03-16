package main

import (
	"chief/cli"
	"fmt"
	"os"
)

func main() {
	// db := db.DatabaseConnection{}
	// db.Connect("sa", "Epilefac57#$!$24042002", "host.docker.internal", "sigma", 1433)

	ci := cli.CLIParser{Strict: true}
	optParsed, error := ci.Parse(ci.Setup())

	if len(error) != 0 {
		fmt.Println(error)
	}

	fmt.Println(optParsed)

	os.Exit(1)
}
