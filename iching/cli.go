package iching

import (
	"fmt"
	"github.com/integrii/flaggy"
	"log"
)

var (
	hexarg      string
	method      string
	translation string
	extended    bool
	ascending   bool
)

func AddCommands(root_cmd *flaggy.Subcommand) {
	root_cmd.AddPositionalValue(&hexarg, "hexcode", 1, false, "Hexagram code")
	root_cmd.String(&method, "m", "method", "Method for generating the primary hexagram.")
	root_cmd.String(&translation, "t", "translation", "Translation file to be used in displaying hexagram information.")
	root_cmd.Bool(&extended, "e", "extended", "Show the anti-hexagram and reversed hexagram.")
	root_cmd.Bool(&ascending, "a", "ascending", "Show the fice ascending hexagrams.")
}

func Execute() {
	var hexagram, secondary Hexagram

	if hexarg != "" {

		if len(hexarg) != 6 {
			log.Fatalf("Invalid syntax in \"%s\": must be six digits long.\n", hexarg)
		}

		for i, r := range hexarg {
			linmap := map[rune]uint8{
				'6': 0, '7': 1,
				'8': 0, '9': 1,
			}
			cmap := map[rune]uint8{
				'6': 1, '7': 0,
				'8': 0, '9': 1,
			}

			num, valid := linmap[r]

			if !valid {
				log.Fatalf("Invalid character in \"%s\": only 6, 7, 8 or 9 allowed.\n", hexarg)
			}

			hexagram.lines |= num << i
			hexagram.changes |= cmap[r] << i
		}
		secondary = SecondaryHexagram(hexagram)
	} else {
		methods := map[string]func() Hexagram{
			"":        YarrowStalk,
			"yarrow":  YarrowStalk,
			"coins":   ThreeCoins,
			"oneline": SingleLine,
		}

		chosen, valid := methods[method]

		if !valid {
			log.Fatalf("Unknown method \"%s\": valid methods are \"yarrow\", \"coins\" and \"oneline\".\n", hexarg)
		}

		hexagram = chosen()
		secondary = SecondaryHexagram(hexagram)
	}

	var info Info

	info.Author = "NO TRANSLATION"
	info.Hexagrams = make([]HexInfo, 64)
	for i := range 64 {
		info.Hexagrams[i].Lines = make([]string, 7)
	}

	if translation != "" {
		info = ReadLua(translation)
	}

	fmt.Printf("H%0d", hexnums[hexagram.lines]+1)
	if hexagram.changes != 0 {
		fmt.Print(":")
		for i := range 6 {
			if Index(hexagram.changes, i) == 1 {
				fmt.Printf("%d ", i+1)
			}
		}
	}

	fmt.Print("\n\n")

	if extended {
		ExtendedDisplay(hexagram, secondary, ReversedHexagram(hexagram), AntiHexagram(hexagram))
	} else {
		HexDisplay(hexagram, info)
		if hexagram != secondary {
			fmt.Println("")
			HexDisplay(secondary, info)
		}
	}

	if ascending {
		AscendingDisplay(hexagram, secondary)
	}
}
