package engine

import (
	"fmt"
	"nand2tetris/compiler-I/lexer"
)

func (e *Engine) CompileClass() error {
	if e.lex.HasMoreTokens() {
		e.lex.Advance()
	}
	if e.lex.GetTokenType() != lexer.Keyword || e.lex.GetLexeme() != "class" {
		return fmt.Errorf(
			"expected keyword 'class',got %s %q", tokenTypeToString[e.lex.GetTokenType()], e.lex.GetLexeme(),
		)
	}

	e.writeLine("<class>")
	e.indent++

	// Class
	tType := tokenTypeToString[e.lex.GetTokenType()]
	token := e.lex.GetLexeme()
	e.writeLine("<" + tType + "> " + token + " </" + tType + ">")

	if e.lex.HasMoreTokens() {
		e.lex.Advance()
	} else {
		return fmt.Errorf("No more tokens")
	}

	if e.lex.GetTokenType() != lexer.Identifier {
		return fmt.Errorf(
			"expected className of token type Identifier,got %s ", tokenTypeToString[e.lex.GetTokenType()],
		)
	}

	// ClassName
	tType = tokenTypeToString[e.lex.GetTokenType()]
	token = e.lex.GetLexeme()
	e.writeLine("<" + tType + "> " + token + " </" + tType + ">")

	if e.lex.HasMoreTokens() {
		e.lex.Advance()
	}

	if e.lex.GetTokenType() != lexer.Symbol {
		return fmt.Errorf(
			"expected '{' of token type Symbol,got %s ", tokenTypeToString[e.lex.GetTokenType()],
		)
	}

	// {
	tType = tokenTypeToString[e.lex.GetTokenType()]
	token = e.lex.GetLexeme()
	e.writeLine("<" + tType + "> " + token + " </" + tType + ">")

	e.compileClassVarDec()

	e.indent--
	e.writeLine("</class>")

	return nil
}
