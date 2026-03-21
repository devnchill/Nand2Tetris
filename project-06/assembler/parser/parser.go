// Package parser is responsible for parsing the asm commmands
package parser

import (
	"bufio"
	"os"
	"strings"
)

type Parser struct {
	inputScanner      *bufio.Scanner
	CurrentCommand    string
	lenCurrentCommand int
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
		CurrentCommand: "",
	}, nil
}

func (p *Parser) Advance() bool {
	for p.inputScanner.Scan() {
		instruction := strings.Split(strings.TrimSpace(p.inputScanner.Text()), "//")[0]
		if instruction == "" {
			continue
		}
		p.CurrentCommand = strings.Join(strings.Split(instruction, " "), "")
		p.lenCurrentCommand = len(p.CurrentCommand)
		return true
	}
	return false
}

// to be called after advancing to new line
func (p *Parser) GetCommandType() (TCommandType, error) {
	if p.CurrentCommand[0] == '@' {
		return ACommand, nil
	}
	if p.CurrentCommand[0] == '(' && p.CurrentCommand[p.lenCurrentCommand-1] == ')' {
		return LCommand, nil
	}
	return CCommand, nil
}

// NOTE: Atm we are not worrying about symbols so we'll convert all xxx to decimal value
func (p *Parser) GetSymbol(commandType TCommandType) (string, error) {
	if commandType == ACommand {
		return p.CurrentCommand[1:], nil
	}
	// for LCommand
	return strings.Split((strings.Split(p.CurrentCommand, ")")[0]), "(")[1], nil
}

// should only be called for C instructions
// Format of CInstruction = getDest = comp;jump
func (p *Parser) HasDest() bool {
	return strings.Contains(p.CurrentCommand, "=")
}

func (p *Parser) HasJump() bool {
	return strings.Contains(p.CurrentCommand, ";")
}

// should only be called for C instructions and only when dest is present
func (p *Parser) GetDest() string {
	return strings.TrimSpace(strings.Split(p.CurrentCommand, "=")[0])
}

// should only be called for C instructions
func (p *Parser) GetComp() string {
	if p.HasDest() && p.HasJump() {
		return strings.Split(strings.Split(p.CurrentCommand, ";")[0], "=")[1]
	} else if p.HasDest() && !p.HasJump() {
		return strings.Split(p.CurrentCommand, "=")[1]
	} else if !p.HasDest() && p.HasJump() {
		return strings.Split(p.CurrentCommand, ";")[0]
	} else {
		return p.CurrentCommand
	}
}

// should only be called for C instructions and only when jump is present
func (p *Parser) GetJump() string {
	if strings.Contains(p.CurrentCommand, ";") {
		return strings.Split(p.CurrentCommand, ";")[1]
	}
	return ""
}
