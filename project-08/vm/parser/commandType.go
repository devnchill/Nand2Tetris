package parser

import (
	"fmt"
	"log"
	"strings"
)

type CommandType int

const (
	C_ARITHMETIC CommandType = iota
	C_PUSH
	C_POP
	C_LABEL
	C_GOTO
	C_IF_GOTO
	C_FUNCTION
	C_RETURN
	C_CALL
)

func (p *Parser) getCommandType() (CommandType, error) {
	fmt.Printf("current command --> %s\n", p.CurrentCommand)
	tokens := strings.Fields(p.CurrentCommand)
	cmd := tokens[0]
	switch cmd {
	case "add", "sub", "neg", "eq", "gt", "lt", "and", "or", "not":
		{
			return C_ARITHMETIC, nil
		}
	case "push":
		{
			return C_PUSH, nil
		}
	case "pop":
		{
			return C_POP, nil
		}
	case "function":
		{
			return C_FUNCTION, nil
		}
	case "call":
		{
			return C_CALL, nil
		}
	case "return":
		{
			return C_RETURN, nil
		}
	case "label":
		{
			return C_LABEL, nil
		}
	case "goto":
		{
			return C_GOTO, nil
		}
	case "if-goto":
		{
			return C_IF_GOTO, nil
		}
	}
	log.Fatal("error determining command type")
	return 0, nil
}
