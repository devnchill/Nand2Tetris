package translator

func (t *Translator) pushSymbol(symbol string) {
	t.writer.WriteString("@" + symbol + "\n")
	t.writer.WriteString("D=M\n")

	t.writer.WriteString("@SP\n")
	t.writer.WriteString("A=M\n")
	t.writer.WriteString("M=D\n")

	t.writer.WriteString("@SP\n")
	t.writer.WriteString("M=M+1\n")
}
