package parser

import (
	"nand2tetris/vm/util"
	"strconv"
	"strings"
)

// should not be called for return command
func (p *Parser) GetFirstArg() string {
	cmdType := p.CurrentCommandType
	if cmdType == C_ARITHMETIC {
		return p.CurrentCommand
	}
	return strings.Fields(p.CurrentCommand)[1]
}

// should not be called for push,pop,function and call
func (p *Parser) GetSecondArg() int {
	cmdType := p.CurrentCommandType
	if cmdType == C_PUSH || cmdType == C_POP || cmdType == C_FUNCTION || cmdType == C_CALL {
		intVal, err := strconv.Atoi(strings.Fields(p.CurrentCommand)[2])
		util.Check(err)
		return intVal
	}
	return 0
}
