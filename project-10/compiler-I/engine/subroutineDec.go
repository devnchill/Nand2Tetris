package engine

import (
	"fmt"
	"nand2tetris/compiler-I/lexer"
)

func (e *Engine) compileSubroutine() error {
	e.writeLine("<subroutineDec>")
	e.indent++
	tokenType, lexeme, err := e.next()
	if err != nil {
		return err
	}
	if tokenType != lexer.Keyword || (lexeme != "constructor" && lexeme != "function" && lexeme != "method") {
		return fmt.Errorf("expected tokenType keyword with lexeme (constructor | function | method) but got %s %s", tokenTypeToString[tokenType], lexeme)
	}
	e.writeToken(tokenType, lexeme)

	tokenType, lexeme, err = e.next()
	if err != nil {
		return err
	}
	if tokenType != lexer.Keyword || (lexeme != "int" && lexeme != "char" && lexeme != "boolean" && lexeme != "void") {
		return fmt.Errorf("expected keyword with lexeme (int | char | boolean | void), got %s %s", tokenTypeToString[tokenType], lexeme)
	}
	e.writeToken(tokenType, lexeme)

	tokenType, lexeme, err = e.next()
	if err != nil {
		return err
	}
	e.writeToken(tokenType, lexeme)

	tokenType, lexeme, err = e.next()
	if err != nil {
		return err
	}
	if tokenType != lexer.Symbol || lexeme != "(" {
		return fmt.Errorf("expected keyword with lexeme '(', got %s %s", tokenTypeToString[tokenType], lexeme)
	}
	e.writeToken(tokenType, lexeme)

	tokenType, lexeme, err = e.next()
	if err != nil {
		return err
	}
	if tokenType != lexer.Identifier {
		return fmt.Errorf("expected keyword type identifier,got %s", tokenTypeToString[tokenType])
	}
	e.writeToken(tokenType, lexeme)

	tokenType, lexeme, err = e.next()
	if err != nil {
		return err
	}
	if tokenType != lexer.Symbol || lexeme != ")" {
		return fmt.Errorf("expected keyword with lexeme ')', got %s %s", tokenTypeToString[tokenType], lexeme)
	}
	e.writeToken(tokenType, lexeme)

	e.indent--
	e.writeLine("</subroutineDec>")

	return nil
}
