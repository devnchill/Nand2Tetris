package parser

import (
	"bufio"
	"log"
	"nand2tetris/vm/util"
	"os"
	"strconv"
	"strings"
)

type CommandType int

const (
	C_ARITHMETIC CommandType = iota
	C_PUSH
	C_POP
	C_LABEL
	C_GOTO
	C_IF
	C_FUNCTION
	C_RETURN
	C_CALL
)

type Parser struct {
	filePointer        *os.File
	inputScanner       *bufio.Scanner
	currentCommand     string
	hasMoreCommand     bool
	currentCommandType CommandType
	arg1               string
	arg2               string
}

func NewParser(file string) *Parser {
	fp, err := os.Open(file)
	util.Check(err)
	inputScanner := bufio.NewScanner(fp)
	return &Parser{
		filePointer:  fp,
		inputScanner: inputScanner,
	}
}

func (p *Parser) advance() {
	for p.inputScanner.Scan() {
		line := strings.TrimSpace(p.inputScanner.Text())
		if line == "" {
			continue
		}
		if idx := strings.Index(line, "\\"); idx != -1 {
			line = line[:idx]
		}
		p.currentCommand = line
		p.hasMoreCommand = true
		cmdType, err := p.getCommandType()
		util.Check(err)
		p.currentCommandType = cmdType
		return
	}
	p.hasMoreCommand = false
}

func (p *Parser) getCommandType() (CommandType, error) {
	tokens := strings.Fields(p.currentCommand)
	cmd := tokens[0]
	switch cmd {
	case "add":
	case "sub":
	case "neg":
	case "eq":
	case "gt":
	case "lt":
	case "and":
	case "or":
	case "not":
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
			return C_IF, nil
		}
	}
	log.Fatal("error determining command type")
	return 0, nil
}

func (p *Parser) getFirstArg() string {
	cmdType := p.currentCommandType
	if cmdType == C_ARITHMETIC {
		return p.currentCommand
	}
	return strings.Fields(p.currentCommand)[1]
}

func (p *Parser) getSecondArg() int {
	cmdType := p.currentCommandType
	if cmdType == C_PUSH || cmdType == C_POP || cmdType == C_FUNCTION || cmdType == C_CALL {
		intVal, err := strconv.Atoi(strings.Fields(p.currentCommand)[2])
		util.Check(err)
		return intVal
	}
	return 0
}

func (p *Parser) Close() {
	p.filePointer.Close()
}
