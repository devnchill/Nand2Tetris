package engine

import (
	"nand2tetris/compiler-I/lexer"
)

/*
parameterList: (type identifier (',' identifier)*)?

	the list is optional, so it is allowed to hold nothing at all
*/
func (e *Engine) compileParameterList() error {
	e.writeLine("<parameterList>")
	e.indent++

	tokenType, lexeme, err := e.peek()
	if err != nil {
		return err
	}
	if tokenType != lexer.Symbol || lexeme != ")" {
		for {
			err = e.expectType()
			if err != nil {
				return err
			}
			err = e.expectIdentifier()
			if err != nil {
				return err
			}

			if !e.peekIsSymbol(",") {
				break
			}
			err = e.expectSymbol(",")
			if err != nil {
				return err
			}
		}
	}

	e.indent--
	e.writeLine("</parameterList>")

	return nil
}
