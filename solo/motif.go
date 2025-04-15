package solo

import "fmt"

func Motif(args []string) {
	yesnomaybe := map[int]string{
		1: "No",
		2: "No",
		3: "Mixed, maybe",
		4: "Mixed, maybe",
		5: "Yes",
		6: "Yes",
	}

	impact := map[int]string{
		1: "Little to none",
		2: "Feeble, weak or limited",
		3: "Feeble, weak or limited",
		4: "Firm, notable or full",
		5: "Firm, notable or full",
		6: "Extreme,lasting or truly severe",
	}

	favorability := map[int]string{
		1: "Completely Unfavorable",
		2: "Mildly Unfavorable",
		3: "Narrowly Favorable",
		4: "Broadly Favorable",
		5: "Completely Favorable",
		6: "Extremely Favorable",
	}

	motifs := map[string]map[int]string{
		"answer": yesnomaybe,
		"impact": impact,
		"favor":  favorability,
	}

	for _, arg := range args {
		roll := d6(1)
		motif := motifs[arg]
		text, ok := motif[roll]
		if !ok {
			text = fmt.Sprintf("%s: %d", arg, roll)
		} else {
			text = fmt.Sprintf("%s: %s", arg, text)
		}
		fmt.Println(text)
	}
}
