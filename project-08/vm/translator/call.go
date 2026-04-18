package translator

import "strconv"

var returnLineGenerator int

func (t *Translator) generateLabelForFunc() string {
	returnLineGenerator++
	return t.currentFunction + "$" + "ret" + strconv.Itoa(returnLineGenerator)
}

func (t *Translator) writeCall(funcName string, numArgs int) {
	// push label for current func address/line no
	retLabel := t.generateLabelForFunc()
	t.writer.WriteString("@" + retLabel + "\n")
	t.writer.WriteString("D=A\n")
	t.writer.WriteString("@SP\n")
	t.writer.WriteString("A=M\n")
	t.writer.WriteString("M=D\n")
	t.writer.WriteString("@SP\n")
	t.writer.WriteString("M=M+1\n")

	// push LCL
	t.pushSymbol("LCL")
	t.pushSymbol("ARG")
	t.pushSymbol("THIS")
	t.pushSymbol("THAT")

	// set ARG
	t.writer.WriteString("@SP\n")
	t.writer.WriteString("D=M\n")
	t.writer.WriteString("@" + strconv.Itoa(numArgs) + "\n")
	t.writer.WriteString("D=D-A\n")
	t.writer.WriteString("@5\n")
	t.writer.WriteString("D=D-A\n")
	t.writer.WriteString("@ARG\n")
	t.writer.WriteString("M=D\n")

	// set LCL
	t.writer.WriteString("@SP\n")
	t.writer.WriteString("D=M\n")
	t.writer.WriteString("@LCL\n")
	t.writer.WriteString("M=D\n")

	// goto f
	t.writer.WriteString("@" + funcName + "\n")
	t.writer.WriteString("0;JMP\n")

	// write label
	t.writer.WriteString("(" + retLabel + ")" + "\n")
}
