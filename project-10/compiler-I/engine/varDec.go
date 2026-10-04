package engine

// varDec: 'var' type identifier (',' identifier)* ';'
func (e *Engine) compileVarDec() error {
	e.writeLine("<varDec>")
	e.indent++

	err := e.expectKeyword("var")
	if err != nil {
		return err
	}
	err = e.expectType()
	if err != nil {
		return err
	}
	for {
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
	err = e.expectSymbol(";")
	if err != nil {
		return err
	}

	e.indent--
	e.writeLine("</varDec>")

	return nil
}
