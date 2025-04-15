package main

import (
	"github.com/integrii/flaggy"
	"github.com/via-dev/oracle/iching"
	"log"
)

var ching_cli = flaggy.NewParser("iching")

func main() {
	ching_cli.Version = "0.1.0"

	iching.AddCommands(&ching_cli.Subcommand)
	err := ching_cli.Parse()

	if err != nil {
		log.Fatal(err)
	}

	iching.Execute()
}
