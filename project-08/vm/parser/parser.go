package parser

import (
	"bufio"
	"fmt"
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
	CurrentCommand     string
	HasMoreCommand     bool
	CurrentCommandType CommandType
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

func (p *Parser) Advance() {
	for p.inputScanner.Scan() {
		line := strings.TrimSpace(p.inputScanner.Text())
		if idx := strings.Index(line, "//"); idx != -1 {
			line = line[:idx]
		}
		if line == "" {
			continue
		}
		p.CurrentCommand = line
		p.HasMoreCommand = true
		cmdType, err := p.getCommandType()
		util.Check(err)
		p.CurrentCommandType = cmdType
		return
	}
	p.HasMoreCommand = false
}

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
			return C_IF, nil
		}
	}
	log.Fatal("error determining command type")
	return 0, nil
}

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

func (p *Parser) Close() {
	p.filePointer.Close()
}
