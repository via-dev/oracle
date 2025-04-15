package geomancy

import (
	"log"

	"github.com/integrii/flaggy"
)

var (
	mother1 string
	mother2 string
	mother3 string
	mother4 string
)

func AddCommands(root_cmd *flaggy.Subcommand) {
	root_cmd.AddPositionalValue(&mother1, "m1", 1, false, "1st mother")
	root_cmd.AddPositionalValue(&mother2, "m2", 2, false, "2nd mother")
	root_cmd.AddPositionalValue(&mother3, "m3", 3, false, "3rd mother")
	root_cmd.AddPositionalValue(&mother4, "m4", 4, false, "4th mother")
}

func Execute() {
	var chart Chart

	margs := [4]string{mother1, mother2, mother3, mother4}
	var mothers [4]uint8

	for i := range margs {
		if margs[i] != "" {
			figure := parseFigure(margs[i])
			mothers[i] = figure
		} else {
			mothers[i] = GetFigure()
		}
	}

	chart = GenChart(mothers)
	PrintShield(chart)
}

func parseFigure(fig string) uint8 {
	var figure uint8

	if len(fig) != 4 {
		log.Fatalf("Invalid size for \"%s\": must be 4 characters in length", fig)
	}

	for i, run := range fig {
		smap := map[rune]uint8{
			':': 0, '.': 1,
			'0': 0, '1': 1}
		num, ok := smap[run]
		if !ok {
			log.Fatalf("Invalid character when parsing \"%s\": must only contain \":\" and \".\" or \"0\" and \"1\"\n", fig)
		}
		figure |= num << i
	}
	return figure
}
