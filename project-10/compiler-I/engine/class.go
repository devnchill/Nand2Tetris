package engine

// class: 'class' identifier '{' classVarDec* subroutineDec* '}'
func (e *Engine) CompileClass() error {
	e.writeLine("<class>")
	e.indent++

	err := e.expectKeyword("class")
	if err != nil {
		return err
	}

	err = e.expectIdentifier()
	if err != nil {
		return err
	}

	err = e.expectSymbol("{")
	if err != nil {
		return err
	}

	for e.peekIsKeyword("static", "field") {
		err = e.compileClassVarDec()
		if err != nil {
			return err
		}
	}

	for e.peekIsKeyword("constructor", "function", "method") {
		err = e.compileSubroutine()
		if err != nil {
			return err
		}
	}

	err = e.expectSymbol("}")
	if err != nil {
		return err
	}

	e.indent--
	e.writeLine("</class>")

	return nil
}
