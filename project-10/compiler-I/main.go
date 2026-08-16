package main

import (
	"nand2tetris/compiler-I/lexer"
)

func main() {
	lexer := lexer.NewLexer("./Makefile")
	lexer.HasMoreTokens()
}
