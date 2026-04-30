package translator

func (t *Translator) WriteLabel(label string) {
	if t.currentFunction == "" {
		panic("no current function found")
	}
	t.writer.WriteString("(" + t.currentFunction + "$" + label + ")" + "\n")
}
