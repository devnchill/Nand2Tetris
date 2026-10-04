package engine

// subroutineDec:
// ('constructor' | 'function' | 'method') type identifier '(' parameterList ')' subroutineBody
func (e *Engine) compileSubroutine() error {
	e.writeLine("<subroutineDec>")
	e.indent++

	err := e.expectKeyword("constructor", "function", "method")
	if err != nil {
		return err
	}

	err = e.expectTypeOrVoid()
	if err != nil {
		return err
	}

	err = e.expectIdentifier()
	if err != nil {
		return err
	}

	err = e.expectSymbol("(")
	if err != nil {
		return err
	}

	err = e.compileParameterList()
	if err != nil {
		return err
	}

	err = e.expectSymbol(")")
	if err != nil {
		return err
	}

	err = e.compileSubroutineBody()
	if err != nil {
		return err
	}

	e.indent--
	e.writeLine("</subroutineDec>")

	return nil
}
