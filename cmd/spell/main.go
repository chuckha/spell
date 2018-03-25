package main

import (
	"fmt"
	"io"
	"os"

	"github.com/chuckha/spell"
	"github.com/chuckha/spell/parsers"
)

const (
	punctuation = ""
)

type Tokenizer interface {
	Tokenize(io.Reader) []string
}

// main looks in some default configuration spots to load a config file.
// .spellconfig
func main() {
	if len(os.Args) < 2 {
		fmt.Println("Requires one argument.")
	}

	path, err := spell.FindConfig("")
	fmt.Println("loading config file:", path)
	if err != nil {
		panic(err)
	}
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	cfg := spell.LoadConfig(f)
	ff := spell.NewFileFinder(cfg)
	files, err := ff.Find()
	if err != nil {
		panic(err)
	}

	for _, f := range files {
		file, err := os.Open(f)
		if err != nil {
			fmt.Printf("error opening file: %v\n", err)
			continue
		}
		defer file.Close()
		code, err := parsers.Parse(file)
		if err != nil {
			panic(err)
		}

		misspellings, err := spell.FindMisspellings(code)
		if err != nil {
			fmt.Println("super error:", err)
		}
		for _, m := range misspellings {
			fmt.Printf("%v: %v\n", f, m.Word)
		}
	}
	fmt.Println("checked these files", files)

	// walk func
	// find file
	// determine file type // can assume it's always go
	// file name => []byte => ast.File => get valid tokens (requires some conifg) => send misspellings back
}

// ksjdflksjdlfks
