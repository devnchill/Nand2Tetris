// Package parser is responsible for parsing the asm commmands
package parser

import (
	"bufio"
	"errors"
	"os"
	"regexp"
	"strings"
)

type Parser struct {
	inputScanner   *bufio.Scanner
	currentCommand string
}

type TCommandType int

const (
	ACommand TCommandType = iota
	CCommand
	LCommand
)

func NewParser(filePath string) (*Parser, error) {
	filePointer, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(filePointer)
	return &Parser{
		inputScanner:   scanner,
		currentCommand: "",
	}, nil
}

func (p *Parser) advance() bool {
	for p.inputScanner.Scan() {
		instruction := strings.Split(strings.TrimSpace(p.inputScanner.Text()), "//")[1]
		if instruction == "" {
			continue
		}
		p.currentCommand = strings.Join(strings.Split(instruction, " "), "")
		return true
	}
	return false
}

// to be called after advancing to new line
func (p *Parser) commandType() (TCommandType, error) {
	if p.currentCommand[0] == '@' {
		return ACommand, nil
	}
	matched1, err := regexp.MatchString("^(.*)$", p.currentCommand)
	if err != nil {
		return 0, err
	}
	if matched1 {
		return LCommand, nil
	}
	// TODO: handle it properly with proper regex
	// guidelines however mentions that they assume each instruction is valid
	matched2, err := regexp.MatchString(".*=.*;.*", p.currentCommand)
	if err != nil {
		return 0, err
	}
	if matched2 {
		return CCommand, nil
	}
	return 0, errors.New("invalid Instruction")
}

// NOTE: Atm we are not worrying about symbols so we'll convert all xxx to decimal value
// TODO: handle labels as well
func (p *Parser) symbol(commandType TCommandType) {
	if commandType == ACommand {
		BinaryToDecimal(p.currentCommand)
	}
}
