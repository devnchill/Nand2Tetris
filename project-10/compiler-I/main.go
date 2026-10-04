package main

import (
	"fmt"
	"nand2tetris/compiler-I/engine"
	"os"
)

func main() {
	inputFilePath := "../test/ArrayTest/Main.jack"
	outputFilePath := "./build/jack.xml"
	if len(os.Args) == 3 {
		inputFilePath, outputFilePath = os.Args[1], os.Args[2]
	}

	engine := engine.NewCompilationEngine(inputFilePath, outputFilePath)
	err := engine.CompileClass()
	engine.Close()
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
}
