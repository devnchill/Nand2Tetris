package translator

import "log"

func (t *Translator) WriteGoto(dest string) {
	if t.currentFunction == "" {
		log.Fatal("no current function found")
		return
	}
	t.writer.WriteString("@" + t.currentFunction + "$" + dest + "\n")
	t.writer.WriteString("0;JMP\n")
}
