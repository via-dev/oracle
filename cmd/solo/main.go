package main

import (
	"github.com/integrii/flaggy"
	"github.com/via-dev/oracle/solo"
	"log"
)

var solo_cli = flaggy.NewParser("solo")

func main() {
	solo_cli.Version = "0.1.0"

	solo.AddCommands(&solo_cli.Subcommand)
	err := solo_cli.Parse()

	if err != nil {
		log.Fatal(err)
	}

	solo.Execute()
}
