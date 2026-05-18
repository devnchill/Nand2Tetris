package translator

func (t *Translator) WriteGoto(dest string) {

	if t.currentFunction == "" {
		t.writer.WriteString("@" + dest + "\n")
	} else {
		t.writer.WriteString("@" + t.currentFunction + "$" + dest + "\n")
	}
	t.writer.WriteString("0;JMP\n")
}
