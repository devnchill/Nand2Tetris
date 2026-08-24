package main

import (
	"nand2tetris/compiler-I/lexer"
)

func main() {
	lexer := lexer.NewLexer("../test/ArrayTest/Main.jack")
	for lexer.HasMoreTokens() {
		lexer.Advance()
	}
}
