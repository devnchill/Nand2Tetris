package engine

import (
	"fmt"
	"nand2tetris/compiler-I/lexer"
)

// statements: statement*
func (e *Engine) compileStatements() error {
	e.writeLine("<statements>")
	e.indent++

	for {
		tokenType, lexeme, err := e.peek()
		if err != nil {
			return err
		}
		if tokenType != lexer.Keyword {
			break
		}

		switch lexeme {
		case "let":
			err = e.compileLet()
		case "if":
			err = e.compileIf()
		case "while":
			err = e.compileWhile()
		case "do":
			err = e.compileDo()
		case "return":
			err = e.compileReturn()
		default:
			return fmt.Errorf("expected statement keyword (let | if | while | do | return), got %s", lexeme)
		}
		if err != nil {
			return err
		}
	}

	e.indent--
	e.writeLine("</statements>")

	return nil
}
