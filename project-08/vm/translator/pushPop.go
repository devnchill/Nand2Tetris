package translator

import (
	"fmt"
	"nand2tetris/vm/parser"
	"strconv"
)

func (t *Translator) WritePushPopCommand(command string, commandType parser.CommandType, segment string, index int, vmFileName string) {

	t.writer.WriteString("\n//" + command + "\n")
	var segmentToBase = map[string]string{
		"local":    "LCL",
		"argument": "ARG",
		"this":     "THIS",
		"that":     "THAT",
		"temp":     "5",
		"pointer":  "3",
	}
	base := ""
	if b, ok := segmentToBase[segment]; ok {
		base = b
	}
	useM := ""
	switch segment {
	case "local", "this", "that", "argument":
		{
			useM = "A=D+M"
		}
	case "temp", "pointer":
		{
			useM = "A=D+A"
		}
	}
	switch commandType {
	case parser.C_PUSH:
		{
			switch segment {
			case "constant":
				{
					fmt.Fprintf(t.writer, "@%d\n", index)
					t.writer.WriteString("D=A\n")
					t.writer.WriteString("@SP\n")
					t.writer.WriteString("A=M\n")
					t.writer.WriteString("M=D\n")
					t.writer.WriteString("@SP\n")
					t.writer.WriteString("M=M+1\n")
				}
			case "static":
				{
					t.writer.WriteString("@" + vmFileName + "." + strconv.Itoa(index) + "\n")
					t.writer.WriteString("D=M\n")
					t.writer.WriteString("@SP\n")
					t.writer.WriteString("A=M\n")
					t.writer.WriteString("M=D\n")
					t.writer.WriteString("@SP\n")
					t.writer.WriteString("M=M+1\n")
				}
			default:
				{
					fmt.Fprintf(t.writer, "@%d\n", index)
					t.writer.WriteString("D=A\n")
					fmt.Fprintf(t.writer, "@%s\n", base)
					fmt.Fprintf(t.writer, "@%s\n", useM)
					t.writer.WriteString("D=M\n")
					t.writer.WriteString("@SP\n")
					t.writer.WriteString("A=M\n")
					t.writer.WriteString("M=D\n")
					t.writer.WriteString("@SP\n")
					t.writer.WriteString("M=M+1\n")
				}
			}
		}
	case parser.C_POP:
		switch segment {
		case "local", "argument", "this", "that", "temp", "pointer":
			{
				fmt.Fprintf(t.writer, "@%d\n", index)
				t.writer.WriteString("D=A\n")
				t.writer.WriteString("@" + base + "\n")
				t.writer.WriteString(useM + "\n")
				t.writer.WriteString("D=A\n")
				t.writer.WriteString("@R13\n")
				t.writer.WriteString("M=D\n")
				t.writer.WriteString("@SP\n")
				t.writer.WriteString("AM=M-1\n")
				t.writer.WriteString("D=M\n")
				t.writer.WriteString("@R13\n")
				t.writer.WriteString("A=M\n")
				t.writer.WriteString("M=D\n")
			}
		case "pop":
			{
				t.writer.WriteString("@SP\n")
				t.writer.WriteString("AM=M-1\n")
				t.writer.WriteString("D=M\n")
				fmt.Fprintf(t.writer, "@%s.%d", vmFileName, index)
				t.writer.WriteString("M=D\n")
			}
		default:
			panic("pop not implemented for segment: " + segment)
		}
	}
}
