package geomancy

import "core:flags"
import "core:fmt"
import "core:os"

geoFlags :: struct {
	mother1: string `args:"pos=0"`,
	mother2: string `args:"pos=1"`,
	mother3: string `args:"pos=2"`,
	mother4: string `args:"pos=3"`,
}

execute :: proc(args: []string) {
	gflags: geoFlags
	flags.parse_or_exit(&gflags, args)
	chart: Chart

	margs := [4]string{gflags.mother1, gflags.mother2, gflags.mother3, gflags.mother4}
	mothers: [4]u8

	for i in 0 ..< 4 {
		if margs[i] != "" {
			mothers[i] = parse_figure(margs[i])
		} else {
			mothers[i] = get_figure()
		}
	}

	chart = gen_chart(mothers)
	print_shield(chart)
}

main :: proc() {
	args := os.args
	execute(args)
}

parse_figure :: proc(fig: string) -> u8 {
	figure: u8

	if len(fig) != 4 {
		fmt.eprintfln("Invalid size for \"%s\": must be 4 characters in length", fig)
		os.exit(1)
	}

	for run, i in fig {
		smap := make(map[rune]u8, 4)
		smap['0'] = 0
		smap['1'] = 1

		if run not_in smap {
			fmt.eprintfln(
				"Invalid character when parsing \"%s\": must only contain \"0\" and \"1\"",
				fig,
			)
			os.exit(1)
		}

		num := smap[run]
		figure |= num << u8(i)
	}
	return figure
}
