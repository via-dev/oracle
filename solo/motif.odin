package solo

import "core:fmt"

motif :: proc(args: []string) {
	yesnomaybe := [6]string{"No", "No", "Mixed, maybe", "Mixed, maybe", "Yes", "Yes"}

	impact := [6]string {
		"Little to none",
		"Feeble, weak or limited",
		"Feeble, weak or limited",
		"Firm, notable or full",
		"Firm, notable or full",
		"Extreme,lasting or truly severe",
	}

	favorability := [6]string {
		"Completely Unfavorable",
		"Mildly Unfavorable",
		"Narrowly Favorable",
		"Broadly Favorable",
		"Completely Favorable",
		"Extremely Favorable",
	}

	for arg in args {
		roll := d6(1) - 1
		switch arg {
		case "answer":
			fmt.printfln("answer: %s", yesnomaybe[roll])

		case "favor":
			fmt.printfln("favor: %s", favorability[roll])

		case "impact":
			fmt.printfln("impact: %s", impact[roll])

		case:
			fmt.printfln("%s: %d", arg, roll + 1)
		}
	}
}
