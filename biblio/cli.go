package biblio

import (
	"github.com/integrii/flaggy"
)

var (
	files []string
)

func AddCommands(root_cmd *flaggy.Subcommand) {
	root_cmd.StringSlice(&files, "f", "filenames", "filenames of generated books")
}

func Execute() {
	if len(files) != 0 {
		for _, name := range files {
			book(name)
		}
	}
}
