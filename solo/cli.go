package solo

import (
	// "flag"
	"fmt"
	"github.com/integrii/flaggy"
	"strconv"
	"strings"
)

var (
	diceCmd  *flaggy.Subcommand
	yesnoCmd *flaggy.Subcommand
	motifCmd *flaggy.Subcommand
	tableCmd *flaggy.Subcommand
	die      string
	odds     string
	motifs   []string
	table    string
)

func AddCommands(root_cmd *flaggy.Subcommand) {
	diceCmd = flaggy.NewSubcommand("dice")
	root_cmd.AttachSubcommand(diceCmd, 1)
	diceCmd.AddPositionalValue(&die, "die", 1, true, "Die to be rolled")

	yesnoCmd = flaggy.NewSubcommand("yesno")
	root_cmd.AttachSubcommand(yesnoCmd, 1)
	yesnoCmd.AddPositionalValue(&odds, "odds", 1, false, "Odds for the question")

	motifCmd = flaggy.NewSubcommand("motif")
	root_cmd.AttachSubcommand(motifCmd, 1)
	motifCmd.StringSlice(&motifs, "f", "flavors", "Flavors for Motif")

	tableCmd = flaggy.NewSubcommand("table")
	root_cmd.AttachSubcommand(tableCmd, 1)
	tableCmd.AddPositionalValue(&table, "table", 1, true, "Table to be rolled on")
	tableCmd.AddPositionalValue(&die, "dice", 2, false, "Dice used on the table.")
}

func Execute() {

	if diceCmd.Used {
		dice, roll := parseDice(die)
		res := fmt.Sprintf("%s: ", die)

		if len(dice) > 0 {
			res += "["
			for i := range dice {
				res += fmt.Sprintf("%d", dice[i])
				if i != len(dice)-1 {
					res += ","
				}
			}
			res += "] = "
		}

		res += fmt.Sprintf("%d", roll)

		fmt.Println(res)
	}

	if motifCmd.Used {
		Motif(motifs)
	}

	if tableCmd.Used {
		ReadTable(table)
	}

	if yesnoCmd.Used {
		Recluse(odds)
	}
}

func parseDice(die string) (dice []int, rolled int) {
	diearg1, diearg2, ok := strings.Cut(die, "d")

	var num int64 = 1
	var sides int64 = 0

	if ok {
		if diearg1 != "" {
			num, _ = strconv.ParseInt(diearg1, 10, 0)
		} else {
			num = 1
		}
		sides, _ = strconv.ParseInt(diearg2, 10, 0)
	} else {
		sides, _ = strconv.ParseInt(diearg1, 10, 0)
	}

	dice, rolled = roll(int(num), int(sides))
	return
}
