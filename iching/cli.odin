package iching

import "core:flags"
import "core:fmt"
import "core:os"

chingFlags :: struct {
	hexarg:      string `args:"pos=0" usage:"Hexagram code."`,
	m:           string `args:"hidden"`,
	method:      string `usage:"Method for generating the primary hexagram."`,
	t:           string `args:"hidden"`,
	translation: string `usage:"Translation file to be used in displaying hexagram information."`,
	e:           bool `args:"hidden"`,
	extended:    bool `usage:"Show the anti-hexagram and reversed hexagram."`,
	a:           bool `args:"hidden"`,
	ascending:   bool `usage:"Show the five ascending hexagrams."`,
}

main :: proc() {
	args := os.args
	execute(args)
}

execute :: proc(args: []string) {
	cflags: chingFlags
	flags.parse_or_exit(&cflags, args)

	hexagram, secondary: Hexagram

	if cflags.hexarg != "" {

		if len(cflags.hexarg) != 6 {
			fmt.eprintfln("Invalid syntax in \"%s\": must be six digits long.", cflags.hexarg)
			os.exit(1)
		}

		for r, i in cflags.hexarg {
			linmap := make(map[rune]u8)
			linmap['6'] = 0
			linmap['7'] = 1
			linmap['8'] = 0
			linmap['9'] = 1

			cmap := make(map[rune]u8)
			cmap['6'] = 1
			cmap['7'] = 0
			cmap['8'] = 0
			cmap['9'] = 1

			num, valid := linmap[r]

			if !valid {
				fmt.eprintfln(
					"Invalid character in \"%s\": only 6, 7, 8 or 9 allowed.",
					cflags.hexarg,
				)
				os.exit(1)
			}

			hexagram.lines |= num << u8(i)
			hexagram.changes |= cmap[r] << u8(i)
		}
		secondary = secondary_hexagram(hexagram)
	} else {
		methods := make(map[string]proc() -> Hexagram, 4)
		methods[""] = yarrow_stalk
		methods["yarrow"] = yarrow_stalk
		methods["coins"] = three_coins
		methods["oneline"] = single_line

		chosen: #type proc() -> Hexagram
		valid: bool

		if cflags.method != "" && cflags.m == "" {
			chosen, valid = methods[cflags.method]
		}
		if cflags.method == "" && cflags.m != "" {
			chosen, valid = methods[cflags.m]
		}
		if cflags.method == "" && cflags.m == "" {
			chosen, valid = methods[""]
		}
		if cflags.method != "" && cflags.m != "" {
			chosen, valid = methods[cflags.m]
		}

		if !valid {
			fmt.eprintfln(
				"Unknown method \"%s\": valid methods are \"yarrow\", \"coins\" and \"oneline\".",
				cflags.hexarg,
			)
			os.exit(1)
		}

		hexagram = chosen()
		secondary = secondary_hexagram(hexagram)
	}

	info: Info

	info.author = "NO TRANSLATION"
	info.hexagrams = make([]struct {
			name:      string,
			image:     string,
			judgement: string,
			lines:     []string,
		}, 64)

	for i in 0 ..< 64 {
		info.hexagrams[i].lines = make([]string, 7)
	}

	if cflags.translation == "" && cflags.t == "" {
		info = read_json("default")
	}
	if cflags.translation != "" || cflags.t != "" {
		if cflags.translation != "" && cflags.t == "" {
			info = read_json(cflags.translation)
		}
		if cflags.translation == "" && cflags.t != "" {
			info = read_json(cflags.t)
		}
	}

	fmt.printf("H%0d", hexnums[hexagram.lines] + 1)
	if hexagram.changes != 0 {
		fmt.print(":")
		for i in 0 ..< 6 {
			if index(hexagram.changes, i) == 1 {
				fmt.printf("%d ", i + 1)
			}
		}
	}

	fmt.print("\n\n")

	if cflags.extended || cflags.e {
		extended_display(hexagram, secondary, reversed_hexagram(hexagram), anti_hexagram(hexagram))
	} else {
		hex_display(hexagram, info)
		if hexagram != secondary {
			fmt.println("")
			hex_display(secondary, info)
		}
	}

	if cflags.ascending || cflags.a {
		ascending_display(hexagram, secondary)
	}
}
