package engine

import "fmt"

// statements: statement*
func (e *Engine) compileStatements() error {
	e.writeLine("<statements>")
	e.indent++

	for e.peekIsStatement() {
		err := e.compileStatement()
		if err != nil {
			return err
		}
	}

	e.indent--
	e.writeLine("</statements>")

	return nil
}

// '{' statements '}'
func (e *Engine) compileBlock() error {
	err := e.expectSymbol("{")
	if err != nil {
		return err
	}

	err = e.compileStatements()
	if err != nil {
		return err
	}

	return e.expectSymbol("}")
}

// statement: letStatement | ifStatement | whileStatement | doStatement | returnStatement | block
func (e *Engine) compileStatement() error {
	_, lexeme, err := e.peek()
	if err != nil {
		return err
	}

	switch lexeme {
	case "{":
		return e.compileBlock()
	case "let":
		return e.compileLet()
	case "if":
		return e.compileIf()
	case "while":
		return e.compileWhile()
	case "do":
		return e.compileDo()
	case "return":
		return e.compileReturn()
	}

	return fmt.Errorf("expected statement keyword (let | if | while | do | return), got %s", lexeme)
}

// letStatement: 'let' varName '=' expression ';'
func (e *Engine) compileLet() error {
	e.writeLine("<letStatement>")
	e.indent++

	err := e.expectKeyword("let")
	if err != nil {
		return err
	}

	err = e.expectVarName()
	if err != nil {
		return err
	}

	err = e.expectSymbol("=")
	if err != nil {
		return err
	}

	err = e.compileExpression()
	if err != nil {
		return err
	}

	err = e.expectSymbol(";")
	if err != nil {
		return err
	}

	e.indent--
	e.writeLine("</letStatement>")

	return nil
}

// doStatement: 'do' (varName | varName '.' identifier) expressionList? ';'
func (e *Engine) compileDo() error {
	e.writeLine("<doStatement>")
	e.indent++

	err := e.expectKeyword("do")
	if err != nil {
		return err
	}

	err = e.expectVarName()
	if err != nil {
		return err
	}

	if e.peekIsSymbol(".") {
		err = e.expectSymbol(".")
		if err != nil {
			return err
		}

		err = e.expectIdentifier()
		if err != nil {
			return err
		}
	}

	err = e.expectExpressionListIfParenthesized()
	if err != nil {
		return err
	}

	err = e.expectSymbol(";")
	if err != nil {
		return err
	}

	e.indent--
	e.writeLine("</doStatement>")

	return nil
}

// ifStatement: 'if' '(' expression ')' statement ('else' statement)?
func (e *Engine) compileIf() error {
	e.writeLine("<ifStatement>")
	e.indent++

	err := e.expectKeyword("if")
	if err != nil {
		return err
	}

	err = e.expectSymbol("(")
	if err != nil {
		return err
	}

	err = e.compileExpression()
	if err != nil {
		return err
	}

	err = e.expectSymbol(")")
	if err != nil {
		return err
	}

	err = e.compileStatement()
	if err != nil {
		return err
	}

	if e.peekIsKeyword("else") {
		err = e.expectKeyword("else")
		if err != nil {
			return err
		}

		err = e.compileStatement()
		if err != nil {
			return err
		}
	}

	e.indent--
	e.writeLine("</ifStatement>")

	return nil
}

// whileStatement: 'while' '(' expression ')' statement
func (e *Engine) compileWhile() error {
	e.writeLine("<whileStatement>")
	e.indent++

	err := e.expectKeyword("while")
	if err != nil {
		return err
	}

	err = e.expectSymbol("(")
	if err != nil {
		return err
	}

	err = e.compileExpression()
	if err != nil {
		return err
	}

	err = e.expectSymbol(")")
	if err != nil {
		return err
	}

	err = e.compileStatement()
	if err != nil {
		return err
	}

	e.indent--
	e.writeLine("</whileStatement>")

	return nil
}

// returnStatement: 'return' expression? ';'
func (e *Engine) compileReturn() error {
	e.writeLine("<returnStatement>")
	e.indent++

	err := e.expectKeyword("return")
	if err != nil {
		return err
	}

	if !e.peekIsSymbol(";") {
		err = e.compileExpression()
		if err != nil {
			return err
		}
	}

	err = e.expectSymbol(";")
	if err != nil {
		return err
	}

	e.indent--
	e.writeLine("</returnStatement>")

	return nil
}
