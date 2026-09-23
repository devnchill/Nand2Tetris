package main

import (
	"fmt"
	"nand2tetris/compiler-I/engine"
	"os"
)

func main() {
	engine := engine.NewCompilationEngine("../test/ArrayTest/Main.jack", "./build/jack.xml")
	err := engine.CompileClass()
	defer engine.Close()
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
}
