package engine

func (e *Engine) writeIndent() {
	for i := 0; i < e.indent; i++ {
		e.Writer.WriteString(" ")
	}
}

func (e *Engine) writeLine(s string) {
	e.writeIndent()
	e.Writer.WriteString(s)
	e.Writer.WriteString("\n")
}
