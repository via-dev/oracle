package iching

import "core:encoding/json"
import "core:fmt"
import "core:log"
import os "core:os/os2"
import "core:path/filepath"


Info :: struct {
	author:    string,
	hexagrams: []struct {
		name:      string,
		image:     string,
		judgement: string,
		lines:     []string,
	},
	trigrams:  struct {
		heaven:   string,
		lake:     string,
		fire:     string,
		thunder:  string,
		wind:     string,
		water:    string,
		mountain: string,
		earth:    string,
	},
}


ReadJson :: proc(filename: string) -> Info {
	cfg_dir, dir_err := os.user_config_dir(context.allocator)
	if dir_err != nil {
		log.fatal(dir_err)
	}

	trans_dir := filepath.join({cfg_dir, "oracle", "iching"})
	file_path := filepath.join({trans_dir, fmt.aprint(filename, ".json", sep = "")})

	data, file_error := os.read_entire_file(file_path, context.allocator)
	if file_error != nil && !os.exists(file_path) {
		fmt.eprintf("Could not find file \"%s\".\n", file_path)
		fmt.eprintln(
			"Make sure that you typed it's name correctly and that the file is located at:",
		)
		fmt.eprintln(trans_dir)
		os.exit(1)
	}

	defer delete(data) // Free the memory at the end

	// Load data from the json bytes directly to the struct
	info: Info
	unmarshal_err := json.unmarshal(data, &info, spec = .JSON5)
	if unmarshal_err != nil {
		fmt.eprintln("Failed to unmarshal the file!")
		fmt.eprintln(unmarshal_err)
		os.exit(1)
	}
	return info
}
