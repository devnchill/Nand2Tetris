package main

import (
	"fmt"
	"nand2tetris/compiler-I/lexer"
)

func main() {
	lexer := lexer.NewLexer("./Makefile")
	fmt.Printf("hasMoreTokens -> %t", lexer.HasMoreTokens())
}
