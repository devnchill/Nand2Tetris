package translator

func (t *Translator) WriteInit() {
	t.writer.WriteString("@256")
	t.writer.WriteString("D=A")
	t.writer.WriteString("@SP")
	t.writer.WriteString("M=D")
	t.WriteCall("Sys.init", 0)
}
