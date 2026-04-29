package main

import (
	"fmt"
	parser "nand2tetris/vm/parser"
	"nand2tetris/vm/translator"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: vm -f <file.vm ...> -d <directory ...>")
		os.Exit(1)
	}

	files, directories := parseArgs(os.Args)
	fmt.Println("files ->", files)
	fmt.Println("directories ->", directories)

	dirFiles, err := collectVMFiles(directories)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	files = append(files, dirFiles...)

	if err := validateVMFiles(files); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	outFile := "Prog.asm"
	t := translator.NewTranslator(outFile)
	defer t.Close()
	for _, file := range files {

		p := parser.NewParser(file)
		defer p.Close()

		fmt.Println(file)

		generateASM(p, t, file)
	}
}
