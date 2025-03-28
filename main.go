package main

import (
	"chief/cli"
	"os"
)

func main() {
	ci := cli.CLIParser{Strict: true}
	ci.Parse(ci.Setup())

	os.Exit(0)
}
