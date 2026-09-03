package main

import (
	"bufio"
	"nand2tetris/compiler-I/lexer"
	"os"
)

func main() {
	lexer := lexer.NewLexer("../test/ArrayTest/Main.jack")
	f, err := os.Create("demoT.Xml")
	if err != nil {
		os.Exit(1)
	}
	w := bufio.NewWriter(f)
	ttype := []string{
		"keyword",
		"symbol",
		"identifier",
		"integerConstant",
		"stringConstant",
	}
	w.WriteString("<tokens>\n")
	for lexer.HasMoreTokens() {
		lexer.Advance()
		t := ttype[lexer.GetTokenType()]
		w.WriteString("<" + t + "> " + lexer.GetLexeme() + " </" + t + ">\n")
	}
	w.WriteString("<tokens>\n")
	w.Flush()
	f.Close()
}
