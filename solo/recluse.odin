package solo

import "core:fmt"

Recluse :: proc(odds: string) {
	yesdie, nodie, odddie, verydie := d6(1), d6(1), d6(1), d6(1)

	switch odds {
	case "likely", "l":
		yesdie = max(yesdie, odddie)
	case "verylikely", "vl":
		yesdie = max(yesdie, odddie, verydie)
	case "unlikely", "u":
		nodie = max(nodie, odddie)
	case "veryunlikely", "vu":
		nodie = max(nodie, odddie, verydie)
	}

	answer: string

	switch {
	case yesdie > nodie:
		answer = "Yes"
		if yesdie > 3 && nodie > 3 {
			answer = fmt.aprintf("%s, and...", answer)
		} else if yesdie < 4 && nodie < 4 {
			answer = fmt.aprintf("%s, but...", answer)
		}

	case yesdie < nodie:
		answer = "No"
		if yesdie > 3 && nodie > 3 {
			answer = fmt.aprintf("%s, and...", answer)
		} else if yesdie < 4 && nodie < 4 {
			answer = fmt.aprintf("%s, but...", answer)
		}

	case yesdie == nodie:
		answer = "Twist / Wrong Question"
	}

	fmt.println(answer)
}
