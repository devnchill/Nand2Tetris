package translator

func (t *Translator) WriteLabel(label string) {
	if t.currentFunction == "" {
		t.writer.WriteString("(" + label + ")" + "\n")
	} else {
		t.writer.WriteString("(" + t.currentFunction + "$" + label + ")" + "\n")
	}
}
