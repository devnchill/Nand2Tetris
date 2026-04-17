package translator

func (t *Translator) WriteIf(label string) {
	label = t.currentFunction + "$" + label
	t.writer.WriteString("@SP\n")
	t.writer.WriteString("AM=M-1\n")
	t.writer.WriteString("D=M\n")
	t.writer.WriteString("@" + label + "\n")
	t.writer.WriteString("D;JNE\n")
}
