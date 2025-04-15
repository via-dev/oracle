package solo

import "fmt"

func Recluse(odds string) {
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

	var answer string

	if yesdie > nodie {
		answer = "Yes"
	}
	if yesdie < nodie {
		answer = "No"
	}
	if yesdie > 3 && nodie > 3 {
		answer += ", and..."
	}
	if yesdie < 4 && nodie < 4 {
		answer += ", but..."
	}
	if yesdie == nodie {
		answer = "Twist / Wrong Question"
	}

	fmt.Println(answer)
}
