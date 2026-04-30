package translator

func (t *Translator) popD() {
	t.writer.WriteString("@SP\n")
	t.writer.WriteString("AM=M-1\n")
	t.writer.WriteString("D=M\n")
}
