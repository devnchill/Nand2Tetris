package engine

import (
	"fmt"
	"nand2tetris/compiler-I/lexer"
)

func (e *Engine) CompileClass() error {
	tokenType, lexeme, err := e.next()
	if err != nil {
		return err
	}
	if tokenType != lexer.Keyword || lexeme != "class" {
		return fmt.Errorf(
			"expected keyword 'class',got %s %q", tokenTypeToString[tokenType], lexeme,
		)
	}

	e.writeLine("<class>")
	e.indent++
	e.writeToken(tokenType, lexeme)

	tokenType, lexeme, err = e.next()
	if err != nil {
		return err
	}
	if tokenType != lexer.Identifier {
		return fmt.Errorf(
			"expected className of token type Identifier,got %s ", tokenTypeToString[tokenType],
		)
	}
	e.writeToken(tokenType, lexeme)

	tokenType, lexeme, err = e.next()
	if err != nil {
		return err
	}
	if tokenType != lexer.Symbol {
		return fmt.Errorf(
			"expected '{' of token type Symbol,got %s ", tokenTypeToString[tokenType],
		)
	}
	e.writeToken(tokenType, lexeme)

	e.compileClassVarDec()

	e.indent--
	e.writeLine("</class>")

	return nil
}
