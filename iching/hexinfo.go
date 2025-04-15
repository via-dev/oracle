package iching

import (
	"fmt"
	lua "github.com/yuin/gopher-lua"
	luar "layeh.com/gopher-luar"
	"log"
	"os"
	"path/filepath"
)

type HexInfo struct {
	Name      string
	Image     string
	Judgement string
	Lines     []string
}

type TriInfo struct {
	Heaven   string
	Lake     string
	Fire     string
	Thunder  string
	Wind     string
	Water    string
	Mountain string
	Earth    string
}

type Info struct {
	Author    string
	Hexagrams []HexInfo
	Trigrams  TriInfo
}

func ReadLua(filename string) Info {
	cfg_dir, dir_err := os.UserConfigDir()
	if dir_err != nil {
		log.Fatal(dir_err)
	}

	trans_dir := filepath.Join(cfg_dir, "oracle", "iching")
	file_path := filepath.Join(trans_dir, fmt.Sprint(filename, ".lua"))

	_, file_error := os.ReadFile(file_path)
	if file_error != nil && os.IsNotExist(file_error) {
		fmt.Printf("Could not find file \"%s\".\n", file_path)
		fmt.Println("Make sure that you typed it's name correctly and that the file is located at:")
		fmt.Println(trans_dir)
		os.Exit(1)
	}

	L := lua.NewState()
	defer L.Close()

	var info Info
	L.SetGlobal("info", luar.New(L, &info))

	lua_error := L.DoFile(file_path)
	if lua_error != nil {
		fmt.Printf("Lua Error: %s\n", lua_error)
		os.Exit(1)
	}

	if len(info.Hexagrams) != 64 {
		fmt.Printf("Error parsing \"%s\": counted %d hexagram entries instead of required 64.", file_path, len(info.Hexagrams))
		os.Exit(1)
	}

	return info
}
