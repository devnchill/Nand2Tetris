package engine

import (
	"fmt"
	"nand2tetris/compiler-I/lexer"
)

// term: integerConstant | stringConstant | keyword | varName |
// varName '[' expression ']' | subroutineCall | '(' expression ')'
// subroutineCall: varName | varName '.' identifier expressionList?
func (e *Engine) compileTerm() error {
	e.writeLine("<term>")
	e.indent++

	tokenType, lexeme, err := e.next()
	if err != nil {
		return err
	}

	switch tokenType {
	case lexer.IntegerConstant:
		err = e.expectIntegerConstant(lexeme)
	case lexer.StringConstant:
		err = e.expectStringConstant(lexeme)
	case lexer.Keyword:
		if !contains([]string{"true", "false", "null", "this"}, lexeme) {
			err = fmt.Errorf("expected keyword (true | false | null | this), got %s", lexeme)
		} else {
			e.writeToken(tokenType, lexeme)
		}
	case lexer.Identifier:
		e.writeToken(tokenType, lexeme)
		err = e.compileVarNameTerm()
	case lexer.Symbol:
		switch lexeme {
		case "(":
			e.writeToken(tokenType, lexeme)
			err = e.compileExpression()
			if err == nil {
				err = e.expectSymbol(")")
			}
		case "-", "~":
			// unary operators, they apply to a single term
			e.writeToken(tokenType, lexeme)
			err = e.compileTerm()
		default:
			err = fmt.Errorf("expected term, got symbol %s", lexeme)
		}
	default:
		err = fmt.Errorf("expected term, got %s %s", tokenTypeToString[tokenType], lexeme)
	}

	if err != nil {
		return err
	}

	e.indent--
	e.writeLine("</term>")

	return nil
}

// array access, subroutine call or a plain varName
// the identifier itself is already consumed when this is called
func (e *Engine) compileVarNameTerm() error {
	if e.peekIsSymbol("[") {
		return e.expectInsideBrackets()
	}

	if e.peekIsSymbol(".") {
		err := e.expectSymbol(".")
		if err != nil {
			return err
		}

		err = e.expectIdentifier()
		if err != nil {
			return err
		}
	}

	return e.expectExpressionListIfParenthesized()
}
