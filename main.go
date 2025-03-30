package main

import (
	"os"

	"github.com/felipee-monteiro/chief/cli"
)

func main() {
	ci := cli.CLIParser{}
	ci.Parse(ci.Setup())

	os.Exit(0)
}
