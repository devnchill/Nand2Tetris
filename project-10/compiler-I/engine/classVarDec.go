package engine

import (
	"fmt"
	"nand2tetris/compiler-I/lexer"
)

func (e *Engine) compileClassVarDec() error {
	// (static | field) type Varname (',' varname)* ';'
	tokenType, lexeme, err := e.next()
	if err != nil {
		return err
	}
	if tokenType != lexer.Keyword || (lexeme != "static" && lexeme != "field") {
		return fmt.Errorf("expected static or field keyword, got %s", tokenTypeToString[tokenType])
	}
	e.writeToken(tokenType, lexeme)

	tokenType, lexeme, err = e.next()
	if err != nil {
		return err
	}
	if tokenType != lexer.Keyword || (lexeme != "int" && lexeme != "char" && lexeme != "boolean") {
		return fmt.Errorf("expected keyword wity lexeme int | char | boolean , got %s %s", tokenTypeToString[tokenType], lexeme)
	}
	e.writeToken(tokenType, lexeme)

	tokenType, lexeme, err = e.next()
	if err != nil {
		return err
	}
	if tokenType != lexer.Identifier {
		return fmt.Errorf("exptected variable name")
	}
	e.writeToken(tokenType, lexeme)

	for {
		tokenType, lexeme, err = e.next()
		if err != nil {
			return err
		}
		if lexeme == "," && tokenType == lexer.Symbol {
			e.writeToken(tokenType, lexeme)
		} else {
			break
		}

		tokenType, lexeme, err = e.next()
		if err != nil {
			return err
		}
		if tokenType == lexer.Identifier {
			e.writeToken(tokenType, lexeme)
		}
	}

	tokenType, lexeme, err = e.next()
	if err != nil {
		return err
	}
	if tokenType == lexer.Symbol && lexeme == ";" {
		e.writeToken(tokenType, lexeme)
	}

	return nil
}
