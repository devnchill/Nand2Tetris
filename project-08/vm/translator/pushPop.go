package translator

import (
	"fmt"
	"nand2tetris/vm/parser"
	"strconv"
)

func (t *Translator) WritePushPopCommand(command string, commandType parser.CommandType, segment string, index int, vmFileName string) {
	t.writer.WriteString("\n//" + command + "\n")
	segmentToSymbol := map[string]string{
		"local":    "LCL",
		"argument": "ARG",
		"this":     "THIS",
		"that":     "THAT",
	}
	switch commandType {
	case parser.C_PUSH:
		{
			switch segment {
			case "local", "argument", "this", "that":
				{
					fmt.Fprintf(t.writer, "@%s\n", segmentToSymbol[segment])
					t.writer.WriteString("D=M\n")
					fmt.Fprintf(t.writer, "@%d\n", index)
					t.writer.WriteString("A=D+A\n")
					t.writer.WriteString("D=M\n")
					t.pushD()
				}
			case "pointer":
				{
					var symbol string
					switch index {
					case 0:
						symbol = "THIS"
					case 1:
						symbol = "THAT"
					default:
						panic("invalid index for pointer segment:" + strconv.Itoa(index))
					}
					fmt.Fprintf(t.writer, "@%s\n", symbol)
					t.writer.WriteString("D=M\n")
					t.pushD()
				}
			case "constant":
				{
					fmt.Fprintf(t.writer, "@%d\n", index)
					t.writer.WriteString("D=A\n")
					t.pushD()
				}
			case "temp":
				{
					addr := 5 + index
					fmt.Fprintf(t.writer, "@%d\n", addr)
					t.writer.WriteString("D=M\n")
					t.pushD()
				}
			case "static":
				{
					symbol := vmFileName + "." + strconv.Itoa(index)
					fmt.Fprintf(t.writer, "@%s\n", symbol)
					t.writer.WriteString("D=M\n")
					t.pushD()
				}
			default:
				{
					panic("push not implemented for segment: " + segment)
				}
			}
		}
	case parser.C_POP:
		{
			switch segment {
			case "local", "argument", "this", "that":
				{
					// calc base addr + index
					fmt.Fprintf(t.writer, "@%s\n", segmentToSymbol[segment])
					t.writer.WriteString("D=M\n")
					fmt.Fprintf(t.writer, "@%d\n", index)
					t.writer.WriteString("D=D+A\n")

					// store base addr in R13 register
					t.writer.WriteString("@R13\n")
					t.writer.WriteString("M=D\n")

					// pop from stack and store element in Data register
					t.popD()

					// grab bases addr + index from R13 and set value of that address to value stored in data register
					t.writer.WriteString("@R13\n")
					t.writer.WriteString("A=M\n")
					t.writer.WriteString("M=D\n")
				}
			case "pointer":
				{
					var symbol string
					switch index {
					case 0:
						symbol = "THIS"
					case 1:
						symbol = "THAT"
					default:
						panic("invalid index for pointer segment:" + strconv.Itoa(index))
					}
					t.popD()
					fmt.Fprintf(t.writer, "@%s\n", symbol)
					t.writer.WriteString("M=D\n")
				}
			case "temp":
				{
					t.popD()
					addr := 5 + index
					fmt.Fprintf(t.writer, "@%d\n", addr)
					t.writer.WriteString("M=D\n")
				}
			case "static":
				{
					t.popD()
					symbol := vmFileName + "." + strconv.Itoa(index)
					fmt.Fprintf(t.writer, "@%s\n", symbol)
					t.writer.WriteString("M=D\n")
				}

			default:
				{
					panic("pop not implemented for segment: " + segment)
				}
			}
		}
	}
}
