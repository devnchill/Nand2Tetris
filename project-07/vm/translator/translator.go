package translator

import (
	"bufio"
	"fmt"
	"nand2tetris/vm/parser"
	"nand2tetris/vm/util"
	"os"
	"strconv"
)

type Translator struct {
	filePointer *os.File
	writer      *bufio.Writer
}

func NewTranslator(file string) *Translator {
	fmt.Printf("Creating new file: %s\n", file)
	fp, err := os.Create(file)
	util.Check(err)
	writer := bufio.NewWriter(fp)
	return &Translator{
		writer:      writer,
		filePointer: fp,
	}
}

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

func (t *Translator) Close() {
	t.writer.Flush()
	t.filePointer.Close()
}
