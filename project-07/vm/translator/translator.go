package translator

import (
	"bufio"
	"fmt"
	"nand2tetris/vm/util"
	"os"
)

type Translator struct {
	filePointer *os.File
	writer      *bufio.Writer
}

func NewTranslator(file string) *Translator {
	fp, err := os.Create(file)
	util.Check(err)
	defer fp.Close()
	writer := bufio.NewWriter(fp)
	return &Translator{
		writer:      writer,
		filePointer: fp,
	}
}

var labelCounter int

func (t *Translator) writeArithmeticCommand(command string) {
	t.writer.WriteString("//" + command + "\n")
	switch command {
	case "add":
		t.writer.WriteString(`
		@SP
		AM=M-1
		D=M
		A=A-1
		M=M+D
		`)
	case "sub":
		t.writer.WriteString(`
		@SP
		AM=M-1
		D=M
		A=A-1
		M=M-D
		`)
	case "neg":
		t.writer.WriteString(`
		@SP
		A=M-1
		M=-M
		`)
	case "eq", "gt", "lt":
		eqLabel := fmt.Sprintf("EQ_%d", labelCounter)
		endLabel := fmt.Sprintf("END_%d", labelCounter)
		labelCounter++

		jumpCond := map[string]string{
			"eq": "JEQ",
			"gt": "JGT",
			"lt": "JLT",
		}[command]

		t.writer.WriteString(fmt.Sprintf(`
		@SP
		AM=M-1
		D=M
		A=A-1
		D=M-D
		@%s
		D;%s
		@SP
		A=M-1
		M=0
		@%s
		0;JMP
		(%s)
		@SP
		A=M-1
		M=-1
		(%s)
		`, eqLabel, jumpCond, endLabel, eqLabel, endLabel))
	case "and":
		t.writer.WriteString(`
		@SP
		AM=M-1
		D=M
		A=A-1
		M=M&D
		`)
	case "or":
		t.writer.WriteString(`
		@SP
		AM=M-1
		D=M
		A=A-1
		M=M|D
		`)
	case "not":
		t.writer.WriteString(`
		@SP
		A=M-1
		M=!M
		`)
	default:
		panic(fmt.Sprintf("invalid arithmetic command: %s", command))
	}
}

func (t *Translator) writePushPopCommand(command string) {
}

func (t *Translator) Close() {
	t.writer.Flush()
	t.filePointer.Close()
}
