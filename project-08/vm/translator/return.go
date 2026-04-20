package translator

func (t *Translator) writeReturn() {
	// FRAME = LCL
	t.writer.WriteString("@LCL\n")
	t.writer.WriteString("D=M\n")
	t.writer.WriteString("@R13\n")
	t.writer.WriteString("M=D\n") // R13 = FRAME

	// RET = *(FRAME - 5)
	t.writer.WriteString("@5\n")
	t.writer.WriteString("A=D-A\n")
	t.writer.WriteString("D=M\n")
	t.writer.WriteString("@R14\n")
	t.writer.WriteString("M=D\n") // R14 = return address

	// *ARG = pop()
	t.writer.WriteString("@SP\n")
	t.writer.WriteString("AM=M-1\n")
	t.writer.WriteString("D=M\n")
	t.writer.WriteString("@ARG\n")
	t.writer.WriteString("A=M\n")
	t.writer.WriteString("M=D\n")

	// SP = ARG + 1
	t.writer.WriteString("@ARG\n")
	t.writer.WriteString("D=M+1\n")
	t.writer.WriteString("@SP\n")
	t.writer.WriteString("M=D\n")

	// THAT = *(FRAME-1)
	t.writer.WriteString("@R13\n")
	t.writer.WriteString("AM=M-1\n")
	t.writer.WriteString("D=M\n")
	t.writer.WriteString("@THAT\n")
	t.writer.WriteString("M=D\n")

	// THIS = *(FRAME-2)
	t.writer.WriteString("@R13\n")
	t.writer.WriteString("AM=M-1\n")
	t.writer.WriteString("D=M\n")
	t.writer.WriteString("@THIS\n")
	t.writer.WriteString("M=D\n")

	// ARG = *(FRAME-3)
	t.writer.WriteString("@R13\n")
	t.writer.WriteString("AM=M-1\n")
	t.writer.WriteString("D=M\n")
	t.writer.WriteString("@ARG\n")
	t.writer.WriteString("M=D\n")

	// LCL = *(FRAME-4)
	t.writer.WriteString("@R13\n")
	t.writer.WriteString("AM=M-1\n")
	t.writer.WriteString("D=M\n")
	t.writer.WriteString("@LCL\n")
	t.writer.WriteString("M=D\n")

	// goto RET
	t.writer.WriteString("@R14\n")
	t.writer.WriteString("A=M\n")
	t.writer.WriteString("0;JMP\n")
}
