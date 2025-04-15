package solo

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	lua "github.com/yuin/gopher-lua"
)

func ReadTable(argument string) {
	cfg_dir, dir_err := os.UserConfigDir()
	if dir_err != nil {
		log.Fatal(dir_err)
	}

	filename, table_fn, called_fn := strings.Cut(argument, ":")

	func_name := "main"
	if called_fn {
		func_name = table_fn
	}

	tables_dir := filepath.Join(cfg_dir, "oracle", "tables")
	file_path := filepath.Join(tables_dir, fmt.Sprint(filename, ".lua"))

	lua.LuaPathDefault += ";"
	lua.LuaPathDefault += fmt.Sprint(tables_dir, lua.LuaDirSep, "?.lua;")
	lua.LuaPathDefault += fmt.Sprint(tables_dir, lua.LuaDirSep, "?", lua.LuaDirSep, "init.lua;")

	_, file_error := os.ReadFile(file_path)
	if file_error != nil && os.IsNotExist(file_error) {
		fmt.Printf("Could not find file \"%s\".\n", file_path)
		fmt.Println("Make sure that you typed it's name correctly and that the file is located at:")
		fmt.Println(tables_dir)
		os.Exit(1)
	}

	L := lua.NewState()
	defer L.Close()

	for name, fn := range lua_funcs {
		L.SetGlobal(name, L.NewFunction(fn))
	}

	lua_error := L.DoFile(file_path)
	if lua_error != nil {
		fmt.Printf("Lua Error: %s\n", lua_error.Error())
		os.Exit(1)
	}

	if fnerr := L.CallByParam(lua.P{
		Fn:      L.GetGlobal(func_name),
		NRet:    0,
		Protect: true,
	}); fnerr != nil {
		panic(fnerr)
	}
}

var lua_funcs = map[string]lua.LGFunction{
	"d3":   lua_d3,
	"d4":   lua_d4,
	"d6":   lua_d6,
	"d8":   lua_d8,
	"d10":  lua_d10,
	"d12":  lua_d12,
	"d20":  lua_d20,
	"d100": lua_d100,
	"dF":   lua_dF,
	"dice": dice,
}

func dice(L *lua.LState) int {
	num := L.ToInt(1)
	sides := L.ToInt(2)
	_, res := roll(num, sides)
	L.Push(lua.LNumber(res))
	return 1
}

func lua_d3(L *lua.LState) int {
	num := L.ToInt(1)
	L.Push(lua.LNumber(d3(num)))
	return 1
}

func lua_d4(L *lua.LState) int {
	num := L.ToInt(1)
	L.Push(lua.LNumber(d4(num)))
	return 1
}

func lua_d6(L *lua.LState) int {
	num := L.ToInt(1)
	L.Push(lua.LNumber(d6(num)))
	return 1
}

func lua_d8(L *lua.LState) int {
	num := L.ToInt(1)
	L.Push(lua.LNumber(d8(num)))
	return 1
}

func lua_d10(L *lua.LState) int {
	num := L.ToInt(1)
	L.Push(lua.LNumber(d10(num)))
	return 1
}

func lua_d12(L *lua.LState) int {
	num := L.ToInt(1)
	L.Push(lua.LNumber(d12(num)))
	return 1
}

func lua_d20(L *lua.LState) int {
	num := L.ToInt(1)
	L.Push(lua.LNumber(d20(num)))
	return 1
}

func lua_d100(L *lua.LState) int {
	num := L.ToInt(1)
	L.Push(lua.LNumber(d100(num)))
	return 1
}

func lua_dF(L *lua.LState) int {
	num := L.ToInt(1)
	L.Push(lua.LNumber(dF(num)))
	return 1
}
