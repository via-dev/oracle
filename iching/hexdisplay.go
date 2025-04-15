package iching

import "fmt"

var hexlines = [2][2]string{
	{"===   ===", "=== X ==="},
	{"=========", "====O===="},
}

func HexDisplay(hexagram Hexagram, info Info) {
	lines, changes := hexagram.lines, hexagram.changes
	upper, lower := NameTrigrams(lines, info.Trigrams)
	hexinfo := info.Hexagrams[hexnums[lines]]

	fmt.Println(hexinfo.Name)

	for i := range 6 {
		idx := 5 - i
		line := Index(lines, idx)
		change := Index(changes, idx)
		str := hexlines[line][change]
		if idx == 4 {
			str += fmt.Sprint(" ", upper)
		}
		if idx == 1 {
			str += fmt.Sprint(" ", lower)
		}
		fmt.Println(str)
	}

	fmt.Print("\n", hexinfo.Judgement, "\n")

	for i := range 6 {
		if Index(changes, i) == 1 {
			fmt.Println(hexinfo.Lines[i])
		}
	}

	if changes == 63 && (lines == 0 || lines == 63) {
		fmt.Println(hexinfo.Lines[6])
	}
}

func NameTrigrams(lines uint8, tri TriInfo) (upper, lower string) {
	upper_num := (lines >> 3)
	lower_num := (lines << 5) >> 5

	trinames := map[uint8]string{
		0b000: tri.Earth,
		0b001: tri.Thunder,
		0b010: tri.Water,
		0b011: tri.Lake,
		0b100: tri.Mountain,
		0b101: tri.Fire,
		0b110: tri.Wind,
		0b111: tri.Heaven,
	}

	upper, lower = trinames[upper_num], trinames[lower_num]
	return
}

func ExtendedDisplay(primary, secondary, reversed, anti Hexagram) {
	if primary != secondary {
		changing := [2]string{" ", ">"}
		fmt.Println(" ANTIHEX  |  PRIMARY  | SECONDARY")
		for i := range 6 {
			idx := (5 - i)
			pline := Index(primary.lines, idx)
			pchange := Index(primary.changes, idx)
			sline := Index(secondary.lines, idx)
			schange := Index(secondary.changes, idx)
			aline := Index(anti.lines, idx)
			achange := Index(anti.changes, idx)
			fmt.Printf(
				"%s %s %s %s %s\n",
				hexlines[aline][0],
				changing[achange],
				hexlines[pline][pchange],
				changing[pchange],
				hexlines[sline][schange])
		}
		fmt.Println("")

		for i := range 6 {
			idx := 5 - i
			rline := Index(reversed.lines, idx)
			rchange := Index(reversed.changes, idx)
			fmt.Printf("            %s\n", hexlines[rline][rchange])
		}
	} else {
		fmt.Println(" ANTIHEX  |  PRIMARY")
		for i := range 6 {
			pline := Index(primary.lines, i)
			pchange := Index(primary.changes, i)
			aline := Index(anti.lines, i)
			fmt.Printf("%s   %s\n", hexlines[aline][0], hexlines[pline][pchange])
		}
	}
}

func AscendingDisplay(primary, secondary Hexagram) {
	fmt.Println("Ascending Hexagrams")
	fmt.Printf("%s > %s > %s > %s > %s\n",
		hexlines[Index(secondary.lines, 0)][0],
		hexlines[Index(secondary.lines, 1)][0],
		hexlines[Index(secondary.lines, 2)][0],
		hexlines[Index(secondary.lines, 3)][0],
		hexlines[Index(secondary.lines, 4)][0])
	fmt.Printf("%s > %s > %s > %s > %s\n",
		hexlines[Index(primary.lines, 5)][0],
		hexlines[Index(secondary.lines, 0)][0],
		hexlines[Index(secondary.lines, 1)][0],
		hexlines[Index(secondary.lines, 2)][0],
		hexlines[Index(secondary.lines, 3)][0])
	fmt.Printf("%s > %s > %s > %s > %s\n",
		hexlines[Index(primary.lines, 4)][0],
		hexlines[Index(primary.lines, 5)][0],
		hexlines[Index(secondary.lines, 0)][0],
		hexlines[Index(secondary.lines, 1)][0],
		hexlines[Index(secondary.lines, 2)][0])
	fmt.Printf("%s > %s > %s > %s > %s\n",
		hexlines[Index(primary.lines, 3)][0],
		hexlines[Index(primary.lines, 4)][0],
		hexlines[Index(primary.lines, 5)][0],
		hexlines[Index(secondary.lines, 0)][0],
		hexlines[Index(secondary.lines, 1)][0])
	fmt.Printf("%s > %s > %s > %s > %s\n",
		hexlines[Index(primary.lines, 2)][0],
		hexlines[Index(primary.lines, 3)][0],
		hexlines[Index(primary.lines, 4)][0],
		hexlines[Index(primary.lines, 5)][0],
		hexlines[Index(secondary.lines, 0)][0])
	fmt.Printf("%s > %s > %s > %s > %s\n",
		hexlines[Index(primary.lines, 1)][0],
		hexlines[Index(primary.lines, 2)][0],
		hexlines[Index(primary.lines, 3)][0],
		hexlines[Index(primary.lines, 4)][0],
		hexlines[Index(primary.lines, 5)][0])
}

/*
The conventional way of ordering the I Ching
hexagrams is through the King Wei sequence.
Users (myself included) will want to organize
their hexagrams using this sequence because it
makes the most intuitive sense from a text copying
standpoint.

However, the King Wei sequence makes no orderly sense
from a purely binary standpoint. This means
I cannot use a hexagrams uint number to index
into the slice of hexagrams returned by the lua
file, because the order of elements will not match.

Therefore, the large map below is used to turn
the hexagram's uint number into the actual slice
index and thus get the correct info dump to be
displayed.

Key: Actual hexagram binary
Value: Corresponding slice index
*/

var hexnums = map[uint8]int{
	0b111111: 1 - 1,
	0b000000: 2 - 1,
	0b010001: 3 - 1,
	0b100010: 4 - 1,
	0b010111: 5 - 1,
	0b111010: 6 - 1,
	0b000010: 7 - 1,
	0b010000: 8 - 1,
	0b110111: 9 - 1,
	0b111011: 10 - 1,
	0b000111: 11 - 1,
	0b111000: 12 - 1,
	0b111101: 13 - 1,
	0b101111: 14 - 1,
	0b000100: 15 - 1,
	0b001000: 16 - 1,
	0b011001: 17 - 1,
	0b100110: 18 - 1,
	0b000011: 19 - 1,
	0b110000: 20 - 1,
	0b101001: 21 - 1,
	0b100101: 22 - 1,
	0b100000: 23 - 1,
	0b000001: 24 - 1,
	0b111001: 25 - 1,
	0b100111: 26 - 1,
	0b100001: 27 - 1,
	0b011110: 28 - 1,
	0b010010: 29 - 1,
	0b101101: 30 - 1,
	0b011100: 31 - 1,
	0b001110: 32 - 1,
	0b111100: 33 - 1,
	0b001111: 34 - 1,
	0b101000: 35 - 1,
	0b000101: 36 - 1,
	0b110101: 37 - 1,
	0b101011: 38 - 1,
	0b010100: 39 - 1,
	0b001010: 40 - 1,
	0b100011: 41 - 1,
	0b110001: 42 - 1,
	0b011111: 43 - 1,
	0b111110: 44 - 1,
	0b011000: 45 - 1,
	0b000110: 46 - 1,
	0b011010: 47 - 1,
	0b010110: 48 - 1,
	0b011101: 49 - 1,
	0b101110: 50 - 1,
	0b001001: 51 - 1,
	0b100100: 52 - 1,
	0b110100: 53 - 1,
	0b001011: 54 - 1,
	0b001101: 55 - 1,
	0b101100: 56 - 1,
	0b110110: 57 - 1,
	0b011011: 58 - 1,
	0b110010: 59 - 1,
	0b010011: 60 - 1,
	0b110011: 61 - 1,
	0b001100: 62 - 1,
	0b010101: 63 - 1,
	0b101010: 64 - 1,
}
