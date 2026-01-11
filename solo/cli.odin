package solo

import "core:flags"
import "core:fmt"
import "core:os"
import "core:strconv"
import "core:strings"

diceFlags :: struct {
	dice:     string `args:"name=die_to_roll,pos=1,required" usage:"Dice to be rolled"`,
	overflow: [dynamic]string `hidden`,
}

motifFlags :: struct {
	flavors:  string `args:"pos=1,required" usage:"Comma-separated list of Motif flavors"`,
	overflow: [dynamic]string `hidden`,
}

recluseFlags :: struct {
	odds:     string `args:"pos=1" usage:"Odds for how likely the answer is."`,
	overflow: [dynamic]string `hidden`,
}

tableFlags :: struct {
	table_name: string `args:"pos=1,required" usage:"Name of the table."`,
	overflow:   [dynamic]string `hidden`,
}

execute :: proc(args: []string) {
	switch args[1] {
	case "dice":
		dflags: diceFlags
		flags.parse_or_exit(&dflags, args)
		die := dflags.dice
		dice, roll := parse_dice(die)
		res: string = fmt.aprintf("%s: ", die)
		if len(dice) > 0 {
			res = strings.concatenate({res, "["})
			for _, i in dice {
				res = strings.concatenate({res, fmt.aprint(dice[i])})
				if i != len(dice) - 1 {
					res = strings.concatenate({res, ","})
				}
			}
			res = strings.concatenate({res, "] = "})
		}

		res = strings.concatenate({res, fmt.aprint(roll)})

		fmt.println(res)
		os.exit(0)

	case "motif":
		mflags: motifFlags
		flags.parse_or_exit(&mflags, args)
		flavors := strings.split(mflags.flavors, ",")
		motif(flavors)
		os.exit(0)

	case "table":
		tflags: tableFlags
		flags.parse_or_exit(&tflags, args)
		read_table(tflags.table_name)
		os.exit(0)

	case "yesno":
		rflags: recluseFlags
		flags.parse_or_exit(&rflags, args)
		odds := rflags.odds
		recluse(odds)
		os.exit(0)
	}
}

main :: proc() {
	args := os.args

	if len(args) > 1 {
		execute(os.args)
	}
}

parse_dice :: proc(die: string) -> (dice: [dynamic]int, rolled: int) {
	dieargs, err := strings.split(die, "d")

	num := 1
	sides := 0

	if err == nil {
		if dieargs[0] != "" {
			num, _ = strconv.parse_int(dieargs[0], 10)
		}
		sides, _ = strconv.parse_int(dieargs[1], 10)
	} else {
		sides, _ = strconv.parse_int(dieargs[1], 10)
	}
	for _ in 0 ..< num {
		this_roll := roll(num, sides)
		rolled += this_roll
		append(&dice, this_roll)
	}
	return
}
