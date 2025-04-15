package main

import (
	"log"
	"os"

	"github.com/integrii/flaggy"
	bib "github.com/via-dev/oracle/biblio"
	geo "github.com/via-dev/oracle/geomancy"
	chi "github.com/via-dev/oracle/iching"
	sol "github.com/via-dev/oracle/solo"
)

var (
	cli = flaggy.NewParser("oracle")

	Biblio   *flaggy.Subcommand
	Ching    *flaggy.Subcommand
	Geomancy *flaggy.Subcommand
	Solo     *flaggy.Subcommand
)

func main() {
	cli.Version = "0.1.0"

	Biblio = flaggy.NewSubcommand("biblio")
	bib.AddCommands(Biblio)
	cli.AttachSubcommand(Biblio, 1)

	Ching = flaggy.NewSubcommand("iching")
	chi.AddCommands(Ching)
	cli.AttachSubcommand(Ching, 1)

	Geomancy = flaggy.NewSubcommand("geomancer")
	geo.AddCommands(Geomancy)
	cli.AttachSubcommand(Geomancy, 1)

	Solo = flaggy.NewSubcommand("solo")
	sol.AddCommands(Solo)
	cli.AttachSubcommand(Solo, 1)

	err := cli.Parse()

	if err != nil {
		log.Fatal(err)
	}

	if Biblio.Used {
		bib.Execute()
		os.Exit(0)
	}

	if Geomancy.Used {
		geo.Execute()
		os.Exit(0)
	}

	if Ching.Used {
		chi.Execute()
		os.Exit(0)
	}

	if Solo.Used {
		sol.Execute()
		os.Exit(0)
	}

	cli.ShowHelp()
	os.Exit(0)
}
