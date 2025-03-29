package main

import (
	"os"

	"github.com/felipee-monteiro/chief/cli"
	"github.com/felipee-monteiro/chief/db"
)

func main() {
	db.Connect()

	ci := cli.CLIParser{}
	ci.Parse(ci.Setup())

	os.Exit(0)
}
