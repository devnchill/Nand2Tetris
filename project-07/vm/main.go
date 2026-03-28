package main

import (
	"fmt"
	parser "nand2tetris/vm/parser"
	"nand2tetris/vm/translator"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: vmtranslator <file.vm>")
		os.Exit(1)
	}
	file := os.Args[1]
	if filepath.Ext(file) != ".vm" {
		fmt.Println("Error: input file must have .vm extension")
		os.Exit(1)
	}
	outFile := filepath.Join(filepath.Dir(file), strings.TrimPrefix(filepath.Base(file), filepath.Ext(file))) + ".asm"
	p := parser.NewParser(file)
	t := translator.NewTranslator(outFile)
}
