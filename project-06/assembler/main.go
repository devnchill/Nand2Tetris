package main

import (
	"fmt"
	"os"

	parser "github.com/devnchill/Nand2Tetris/project-06/assembler/parser"
)

func main() {
	filePath := os.Args[1]
	p, err := parser.NewParser(filePath)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	p.Parse()
}
