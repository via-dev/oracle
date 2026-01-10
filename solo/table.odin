package solo


import "base:runtime"
import "core:fmt"
import "core:log"
import os "core:os/os2"
import "core:path/filepath"
import "core:strings"

import lua "vendor:lua/5.4"

read_table :: proc(argument: string) {
	cfg_dir, dir_err := os.user_config_dir(context.allocator)
	if dir_err != nil {
		log.fatal(dir_err)
	}

	filename, table_fn: string
	args := strings.split_n(argument, ":", 2)

	if len(args) == 1 {
		filename = args[0]
	}
	if len(args) == 2 {
		filename = args[0]
		table_fn = args[1]
	}

	func_name := "main"
	if table_fn != "" {
		func_name = table_fn
	}

	tables_dir := filepath.join({cfg_dir, "oracle", "tables"})
	file_path := filepath.join({tables_dir, strings.concatenate({filename, ".lua"})})

	script_data, file_error := os.read_entire_file(file_path, context.allocator)
	if file_error != nil && !os.exists(file_path) {
		fmt.eprintf("Could not find file \"%s\".\n", file_path)
		fmt.eprintln(
			"Make sure that you typed it's name correctly and that the file is located at:",
		)
		fmt.eprintln(tables_dir)
		os.exit(1)
	}

	old_cwd, _ := os.get_working_directory(context.allocator)
	cwd_err := os.set_working_directory(tables_dir)
	if cwd_err != nil {
		log.fatal(cwd_err)
	}

	L := lua.L_newstate()
	defer lua.close(L)

	lua.L_openlibs(L)

	lua_funcs := [?]struct {
		name: string,
		cf:   lua.CFunction,
	} {
		{name = "d3", cf = lua_d3},
		{name = "d4", cf = lua_d4},
		{name = "d6", cf = lua_d6},
		{name = "d8", cf = lua_d8},
		{name = "d10", cf = lua_d10},
		{name = "d12", cf = lua_d12},
		{name = "d20", cf = lua_d20},
		{name = "d100", cf = lua_d100},
		{name = "dF", cf = lua_dF},
		{name = "dice", cf = dice},
	}

	for f in lua_funcs {
		lua.pushcfunction(L, f.cf)
		lua.setglobal(L, fmt.caprint(f.name))
	}

	// Run code and check if it succeeded
	if lua.L_dostring(L, fmt.caprint(string(script_data))) != 0 {
		// Get the error string from the top of the stack and print it
		error := lua.tostring(L, -1)
		fmt.eprintln(error)
		// Pop the error off of the stack
		lua.pop(L, 1)
	}

	lua.getglobal(L, fmt.caprint(func_name))
	call_status := lua.pcall(L, 0, 0, 0)
	if lua.Status(call_status) != .OK {
		// Get the error string from the top of the stack and print it
		error := lua.tostring(L, -1)
		fmt.println(error)
		// Pop the error off of the stack
		lua.pop(L, 1)
	}

	cwd_err = os.set_working_directory(old_cwd)
	if cwd_err != nil {
		log.fatal(cwd_err)
	}
}


dice :: proc "c" (L: ^lua.State) -> i32 {
	context = runtime.default_context()
	num := lua.tointeger(L, 1)
	sides := lua.tointeger(L, 2)
	res := roll(int(num), int(sides))
	lua.pushnumber(L, lua.Number(res))
	return 1
}

lua_d3 :: proc "c" (L: ^lua.State) -> i32 {
	context = runtime.default_context()
	num := lua.tointeger(L, 1)
	res := d3(int(num))
	lua.pushnumber(L, lua.Number(res))
	return 1
}

lua_d4 :: proc "c" (L: ^lua.State) -> i32 {
	context = runtime.default_context()
	num := lua.tointeger(L, 1)
	res := d4(int(num))
	lua.pushnumber(L, lua.Number(res))
	return 1
}

lua_d6 :: proc "c" (L: ^lua.State) -> i32 {
	context = runtime.default_context()
	num := lua.tointeger(L, 1)
	res := d6(int(num))
	lua.pushnumber(L, lua.Number(res))
	return 1
}

lua_d8 :: proc "c" (L: ^lua.State) -> i32 {
	context = runtime.default_context()
	num := lua.tointeger(L, 1)
	res := d8(int(num))
	lua.pushnumber(L, lua.Number(res))
	return 1
}

lua_d10 :: proc "c" (L: ^lua.State) -> i32 {
	context = runtime.default_context()
	num := lua.tointeger(L, 1)
	res := d10(int(num))
	lua.pushnumber(L, lua.Number(res))
	return 1
}

lua_d12 :: proc "c" (L: ^lua.State) -> i32 {
	context = runtime.default_context()
	num := lua.tointeger(L, 1)
	res := d12(int(num))
	lua.pushnumber(L, lua.Number(res))
	return 1
}

lua_d20 :: proc "c" (L: ^lua.State) -> i32 {
	context = runtime.default_context()
	num := lua.tointeger(L, 1)
	res := d20(int(num))
	lua.pushnumber(L, lua.Number(res))
	return 1
}

lua_d100 :: proc "c" (L: ^lua.State) -> i32 {
	context = runtime.default_context()
	num := lua.tointeger(L, 1)
	res := d100(int(num))
	lua.pushnumber(L, lua.Number(res))
	return 1
}

lua_dF :: proc "c" (L: ^lua.State) -> i32 {
	context = runtime.default_context()
	num := lua.tointeger(L, 1)
	res := dF(int(num))
	lua.pushnumber(L, lua.Number(res))
	return 1
}
