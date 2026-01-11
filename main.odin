package main

import "core:fmt"
import "core:os"

import "./geomancy"
import "./iching"
import "./solo"

main :: proc() {
	if len(os.args) == 1 || os.args[1] == "help" {
		fmt.print(USAGE)
		os.exit(0)
	}

	switch os.args[1] {
	case "geomancy":
		geomancy.execute(os.args[1:])
	case "iching":
		iching.execute(os.args[1:])
	case "solo":
		solo.execute(os.args[1:])
	case "help":
		fmt.print(USAGE)
		os.exit(0)
	case:
		fmt.eprintfln("Error: No command named \"%s\"", os.args[1])
		os.exit(1)
	}
}

USAGE :: "The help message"
