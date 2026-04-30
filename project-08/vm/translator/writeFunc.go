package translator

func (t *Translator) pushZeroKTimes(n int) {
	t.writer.WriteString("@0\n")
	t.writer.WriteString("D=A\n")
	for range n {
		t.writer.WriteString("@SP\n")
		t.writer.WriteString("A=M\n")
		t.writer.WriteString("M=D\n")
		t.writer.WriteString("@SP\n")
		t.writer.WriteString("M=M+1\n")
	}
}

func (t *Translator) WriteFunction(funName string, numLocalVars int) {
	t.currentFunction = funName
	t.writer.WriteString("(" + funName + ")" + "\n")
	t.pushZeroKTimes(numLocalVars)
}
