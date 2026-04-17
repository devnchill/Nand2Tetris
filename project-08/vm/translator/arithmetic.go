package translator

import "fmt"

var labelCounter int

func (t *Translator) WriteArithmeticCommand(command string) {
	t.writer.WriteString("\n//" + command + "\n")
	switch command {
	case "add":
		t.writer.WriteString("@SP\n")
		t.writer.WriteString("AM=M-1\n")
		t.writer.WriteString("D=M\n")
		t.writer.WriteString("A=A-1\n")
		t.writer.WriteString("M=M+D\n")
	case "sub":
		t.writer.WriteString("@SP\n")
		t.writer.WriteString("AM=M-1\n")
		t.writer.WriteString("D=M\n")
		t.writer.WriteString("A=A-1\n")
		t.writer.WriteString("M=M-D\n")
	case "neg":
		t.writer.WriteString("@SP\n")
		t.writer.WriteString("A=M-1\n")
		t.writer.WriteString("M=-M\n")
	case "eq", "gt", "lt":
		eqLabel := fmt.Sprintf("EQ_%d", labelCounter)
		endLabel := fmt.Sprintf("END_%d", labelCounter)
		labelCounter++

		jumpCond := map[string]string{
			"eq": "JEQ",
			"gt": "JGT",
			"lt": "JLT",
		}[command]

		t.writer.WriteString("@SP\n")
		t.writer.WriteString("AM=M-1\n")
		t.writer.WriteString("D=M\n")
		t.writer.WriteString("A=A-1\n")
		t.writer.WriteString("D=M-D\n")
		t.writer.WriteString("@" + eqLabel + "\n")
		t.writer.WriteString("D;" + jumpCond + "\n")
		t.writer.WriteString("@SP\n")
		t.writer.WriteString("A=M-1\n")
		t.writer.WriteString("M=0\n")
		t.writer.WriteString("@" + endLabel + "\n")
		t.writer.WriteString("0;JMP\n")
		t.writer.WriteString(eqLabel + "\n")
		t.writer.WriteString("@SP\n")
		t.writer.WriteString("A=M-1\n")
		t.writer.WriteString("M=-1\n")
		t.writer.WriteString(endLabel + "\n")
	case "and":
		{
			t.writer.WriteString("@SP\n")
			t.writer.WriteString("AM=M-1\n")
			t.writer.WriteString("D=M\n")
			t.writer.WriteString("A=A-1\n")
			t.writer.WriteString("M=M&D\n")
		}
	case "or":
		{
			t.writer.WriteString("@SP\n")
			t.writer.WriteString("AM = M - 1\n")
			t.writer.WriteString("D = M\n")
			t.writer.WriteString("A = A - 1\n")
			t.writer.WriteString("M = M | D\n")
		}
	case "not":
		t.writer.WriteString("@SP\n")
		t.writer.WriteString("A=M-1\n")
		t.writer.WriteString("M=!M\n")
	default:
		panic(fmt.Sprintf("invalid arithmetic command: %s", command))
	}
}
