package geomancy

import (
	"fmt"
)

var linstr = [2]string{"O   O", "  O  "}

func PrintShield(chart Chart) {
	str := "Shield Chart\n"
	str += "+-------+-------+-------+-------+-------+-------+-------+-------+\n"
	for i := range 4 {
		str += fmt.Sprintf("| %s | %s | %s | %s ",
			linstr[Index(chart.Daughters[3], i)],
			linstr[Index(chart.Daughters[2], i)],
			linstr[Index(chart.Daughters[1], i)],
			linstr[Index(chart.Daughters[0], i)])
		str += fmt.Sprintf("| %s | %s | %s | %s |\n",
			linstr[Index(chart.Mothers[3], i)],
			linstr[Index(chart.Mothers[2], i)],
			linstr[Index(chart.Mothers[1], i)],
			linstr[Index(chart.Mothers[0], i)])
	}
	str += "+---------------+---------------+---------------+---------------+\n"
	for i := range 4 {
		str += fmt.Sprintf("|     %s     |     %s     |     %s     |     %s     |\n",
			linstr[Index(chart.Nieces[3], i)],
			linstr[Index(chart.Nieces[2], i)],
			linstr[Index(chart.Nieces[1], i)],
			linstr[Index(chart.Nieces[0], i)])
	}
	str += "+-------------------------------+-------------------------------+\n"
	for i := range 4 {
		str += fmt.Sprintf("|             %s             |             %s             |\n",
			linstr[Index(chart.Lwitness, i)],
			linstr[Index(chart.Rwitness, i)])
	}
	str += "+---------------------------------------------------------------+\n"
	for i := range 4 {
		str += fmt.Sprintf("|                             %s                             |\n",
			linstr[Index(chart.Judge, i)])
	}
	str += "+---------------------------------------------------------------+\n"
	fmt.Print(str)
}

// For later
// const house = `
// 	       XI        X         IX        VIII
//     +---------+---------+---------+---------+
//     |  %s  |  %s  |  %s  |  %s  |
//  XI |  %s  |  %s  |  %s  |  %s  | VIII
//     |  %s  |  %s  |  %s  |  %s  |
//     |  %s  |  %s  |  %s  |  %s  |
//     +---------+---------+---------+---------+
//     |  %s  |  %s     %s  |  %s  |
// XII |  %s  |  %s     %s  |  %s  | VII
//     |  %s  |  %s     %s  |  %s  |
//     |  %s  |  %s     %s  |  %s  |
//     +---------+                   +---------+
//     |  %s  |       %s       |  %s  |
//     |  %s  |       %s       |  %s  |
//   I |  %s  |       %s       |  %s  | VI
//     |  %s  |       %s       |  %s  |
//     +---------+---------+---------+---------+
//     |  %s  |  %s  |  %s  |  %s  |
//     |  %s  |  %s  |  %s  |  %s  |
//  II |  %s  |  %s  |  %s  |  %s  | V
//     |  %s  |  %s  |  %s  |  %s  |
//     +---------+---------+---------+---------+
// 	      II        III       IV         V`
