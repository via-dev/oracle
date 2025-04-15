package main

import (
	"github.com/integrii/flaggy"
	geo "github.com/via-dev/oracle/geomancy"
	"log"
)

var geo_cli = flaggy.NewParser("geomancer")

func main() {
	geo_cli.Version = "0.1.0"

	geo.AddCommands(&geo_cli.Subcommand)
	err := geo_cli.Parse()

	if err != nil {
		log.Fatal(err)
	}

	geo.Execute()
}
