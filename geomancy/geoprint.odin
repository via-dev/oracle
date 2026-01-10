package geomancy

import "core:fmt"
import "core:strings"

linstr := [2]string{"O   O", "  O  "}

print_shield :: proc(chart: Chart) {
	shield: strings.Builder
	str := "Shield Chart\n"
	strings.write_string(&shield, str)
	str = "+-------+-------+-------+-------+-------+-------+-------+-------+\n"
	strings.write_string(&shield, str)

	for i in 0 ..< 4 {
		str = fmt.aprintf(
			"| %s | %s | %s | %s ",
			linstr[index(chart.Daughters[3], i)],
			linstr[index(chart.Daughters[2], i)],
			linstr[index(chart.Daughters[1], i)],
			linstr[index(chart.Daughters[0], i)],
		)
		strings.write_string(&shield, str)

		str = fmt.aprintf(
			"| %s | %s | %s | %s |\n",
			linstr[index(chart.Mothers[3], i)],
			linstr[index(chart.Mothers[2], i)],
			linstr[index(chart.Mothers[1], i)],
			linstr[index(chart.Mothers[0], i)],
		)
		strings.write_string(&shield, str)
	}
	str = "+---------------+---------------+---------------+---------------+\n"
	strings.write_string(&shield, str)
	for i in 0 ..< 4 {
		str = fmt.aprintf(
			"|     %s     |     %s     |     %s     |     %s     |\n",
			linstr[index(chart.Nieces[3], i)],
			linstr[index(chart.Nieces[2], i)],
			linstr[index(chart.Nieces[1], i)],
			linstr[index(chart.Nieces[0], i)],
		)
		strings.write_string(&shield, str)
	}
	str = "+-------------------------------+-------------------------------+\n"
	strings.write_string(&shield, str)
	for i in 0 ..< 4 {
		str = fmt.aprintf(
			"|             %s             |             %s             |\n",
			linstr[index(chart.Lwitness, i)],
			linstr[index(chart.Rwitness, i)],
		)
		strings.write_string(&shield, str)
	}
	str = "+---------------------------------------------------------------+\n"
	strings.write_string(&shield, str)
	for i in 0 ..< 4 {
		str = fmt.aprintf(
			"|                             %s                             |\n",
			linstr[index(chart.Judge, i)],
		)
		strings.write_string(&shield, str)
	}
	str = "+---------------------------------------------------------------+\n"
	strings.write_string(&shield, str)
	fmt.print(strings.to_string(shield))
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
