package translator

import "log"

func (t *Translator) WriteLabel(label string) {
	if t.currentFunction == "" {
		log.Fatal("no current function found")
		return
	}
	t.writer.WriteString("(" + t.currentFunction + "$" + label + ")" + "\n")
}
