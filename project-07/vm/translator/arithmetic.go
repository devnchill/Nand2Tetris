package translator

import "fmt"

var nextLabel = func() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}()

func (t *Translator) WriteArithmeticCommand(command string) {
	t.writer.WriteString("\n//" + command + "\n")
	switch command {
	case "add":
		{
			t.popD()
			t.writer.WriteString("@SP\n")
			t.writer.WriteString("A=M-1\n")
			t.writer.WriteString("M=D+M\n")
		}
	case "sub":
		{
			t.popD()
			t.writer.WriteString("@SP\n")
			t.writer.WriteString("A=M-1\n")
			t.writer.WriteString("M=M-D\n")
		}
	case "neg":
		{
			t.writer.WriteString("@SP\n")
			t.writer.WriteString("A=M-1\n")
			t.writer.WriteString("M=-M\n")
		}
	case "eq", "lt", "gt":
		{
			vmToAsmSymol := map[string]string{
				"eq": "JEQ",
				"lt": "JLT",
				"gt": "JGT",
			}
			/*
				@SP -> A = 0 , M = RAM[0]
				AM = M-1 -> A = RAM[0] - 1 , M = RAM[0]-1
				hence A pointing to topmost elemnt of stack
				RAM[SP] pointing to topmost elemnt of stack
			*/
			symbolicLabel := fmt.Sprintf("%s_%d", vmToAsmSymol[command], nextLabel())
			endLabel := fmt.Sprintf("END_%d", nextLabel())
			t.writer.WriteString("@SP\n")
			t.writer.WriteString("AM=M-1\n")
			t.writer.WriteString("D=M\n")
			t.writer.WriteString("A=A-1\n")
			t.writer.WriteString("D=M-D\n")

			t.writer.WriteString("@" + symbolicLabel + "\n")
			t.writer.WriteString("D;" + vmToAsmSymol[command] + "\n")

			// if they are not equal
			t.writer.WriteString("@SP\n")
			t.writer.WriteString("A=M-1\n")
			t.writer.WriteString("M=0\n")
			t.writer.WriteString("@" + endLabel + "\n")
			t.writer.WriteString("0;JMP\n")

			// define equal label
			t.writer.WriteString("(" + symbolicLabel + ")\n")
			t.writer.WriteString("@SP\n")
			t.writer.WriteString("A=M-1\n")
			t.writer.WriteString("M=-1\n")
			t.writer.WriteString("@" + endLabel + "\n")
			t.writer.WriteString("0;JMP\n")

			// define end label
			t.writer.WriteString("(" + endLabel + ")\n")
		}
	case "and":
		{
			t.writer.WriteString("@SP\n")
			t.writer.WriteString("AM=M-1\n")
			t.writer.WriteString("D=M\n")
			t.writer.WriteString("A=A-1\n")
			t.writer.WriteString("M=D&M\n")
		}
	case "or":
		{
			t.writer.WriteString("@SP\n")
			t.writer.WriteString("AM=M-1\n")
			t.writer.WriteString("D=M\n")
			t.writer.WriteString("A=A-1\n")
			t.writer.WriteString("M=D|M\n")
		}
	case "not":
		t.writer.WriteString("@SP\n")
		t.writer.WriteString("A=M-1\n")
		t.writer.WriteString("M=!M\n")
	default:
		panic(fmt.Sprintf("invalid arithmetic command: %s", command))
	}
}
