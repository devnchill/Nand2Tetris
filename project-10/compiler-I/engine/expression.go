package engine

var binaryOperators = []string{"+", "-", "*", "/", "<", ">", "=", "&", "|"}

// expression: term (('+' | '-') term)*
// the arithmetic (* /), comparison (< > <= >= <>) and logical (= & |)
// operators are method calls on ints and booleans in the target code,
// they are kept as binary operators of the expression here
func (e *Engine) compileExpression() error {
	e.writeLine("<expression>")
	e.indent++

	err := e.compileTerm()
	if err != nil {
		return err
	}

	for e.peekIsSymbol(binaryOperators...) {
		err = e.compileBinaryOperator()
		if err != nil {
			return err
		}

		err = e.compileTerm()
		if err != nil {
			return err
		}
	}

	e.indent--
	e.writeLine("</expression>")

	return nil
}

// '<=', '>=' and '<>' are tokenized as two separate symbols
func (e *Engine) compileBinaryOperator() error {
	_, lexeme, err := e.peek()
	if err != nil {
		return err
	}

	err = e.expectSymbol(binaryOperators...)
	if err != nil {
		return err
	}

	if contains([]string{"<", ">"}, lexeme) && e.peekIsSymbol("=", ">") {
		return e.expectSymbol("=", ">")
	}

	return nil
}

/*
expressionList: (expression (',' expression)*)?

	the list is optional, so it is allowed to hold nothing at all
*/
func (e *Engine) compileExpressionList() error {
	e.writeLine("<expressionList>")
	e.indent++

	if !e.peekIsSymbol(")") {
		for {
			err := e.compileExpression()
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
	e.writeLine("</expressionList>")

	return nil
}
