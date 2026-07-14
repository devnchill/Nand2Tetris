package translator

func (t *Translator) WriteInit() {
	t.writer.WriteString("//bootstrapping code\n")
	t.writer.WriteString("@256\n")
	t.writer.WriteString("D=A\n")
	t.writer.WriteString("@SP\n")
	t.writer.WriteString("M=D\n")
	t.WriteCall("Sys.init", 0)
}
