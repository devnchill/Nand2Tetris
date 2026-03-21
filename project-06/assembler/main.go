package main

import (
	"fmt"
	"os"

	parser "github.com/devnchill/Nand2Tetris/project-06/assembler/parser"
	symboltable "github.com/devnchill/Nand2Tetris/project-06/assembler/symbolTable"
	translator "github.com/devnchill/Nand2Tetris/project-06/assembler/translator"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run . <file.asm>")
		return
	}
	filePath := os.Args[1]
	p, err := parser.NewParser(filePath)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	p2, err := parser.NewParser(filePath)
	t := translator.NewTranslator()
	table := symboltable.NewSymbolTable()
	firstPass(p, table)
	secondPass(p2, table, &t)
}
