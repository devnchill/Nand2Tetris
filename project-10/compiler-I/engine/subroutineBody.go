package engine

// subroutineBody: '{' varDec* statements '}'
func (e *Engine) compileSubroutineBody() error {
	e.writeLine("<subroutineBody>")
	e.indent++

	err := e.expectSymbol("{")
	if err != nil {
		return err
	}

	for e.peekIsKeyword("var") {
		err = e.compileVarDec()
		if err != nil {
			return err
		}
	}

	err = e.compileStatements()
	if err != nil {
		return err
	}

	err = e.expectSymbol("}")
	if err != nil {
		return err
	}

	e.indent--
	e.writeLine("</subroutineBody>")

	return nil
}
