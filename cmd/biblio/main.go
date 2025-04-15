package main

import (
	"github.com/integrii/flaggy"
	"github.com/via-dev/oracle/biblio"
	"log"
)

var biblio_cli = flaggy.NewParser("biblio")

func main() {
	biblio_cli.Version = "0.1.0"

	biblio.AddCommands(&biblio_cli.Subcommand)
	err := biblio_cli.Parse()

	if err != nil {
		log.Fatal(err)
	}

	biblio.Execute()
}
