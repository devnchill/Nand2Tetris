// Package parser is responsible for parsing the asm commmands
package parser

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/devnchill/Nand2Tetris/project-06/assembler/util"
)

type Parser struct {
	inputScanner      *bufio.Scanner
	currentCommand    string
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
		currentCommand: "",
	}, nil
}

func (p *Parser) advance() bool {
	for p.inputScanner.Scan() {
		instruction := strings.Split(strings.TrimSpace(p.inputScanner.Text()), "//")[0]
		if instruction == "" {
			continue
		}
		p.currentCommand = strings.Join(strings.Split(instruction, " "), "")
		p.lenCurrentCommand = len(p.currentCommand)
		return true
	}
	return false
}

// to be called after advancing to new line
func (p *Parser) getCommandType() (TCommandType, error) {
	if p.currentCommand[0] == '@' {
		return ACommand, nil
	}
	if p.currentCommand[0] == '(' && p.currentCommand[p.lenCurrentCommand-1] == ')' {
		return LCommand, nil
	}
	return CCommand, nil
}

// NOTE: Atm we are not worrying about symbols so we'll convert all xxx to decimal value
func (p *Parser) getSymbol(commandType TCommandType) (string, error) {
	if commandType == ACommand {
		val, err := util.DecimalToBinary(p.currentCommand)
		if err != nil {
			return "", err
		}
		return val, nil
	}
	return strings.Split((strings.Split(p.currentCommand, ")")[0]), "(")[1], nil
}

// should only be called for C instructions
// Format of CInstruction = getDest = comp;jump
func (p *Parser) getDest() string {
	return strings.TrimSpace(strings.Split(p.currentCommand, "=")[0])
}

// should only be called for C instructions
func (p *Parser) getComp() string {
	return strings.Split(strings.Split(p.currentCommand, ";")[0], "=")[1]
}

// should only be called for C instructions
func (p *Parser) getJump() string {
	if strings.Contains(p.currentCommand, ";") {
		return strings.Split(p.currentCommand, ";")[1]
	}
	return ""
}

func (p *Parser) Parse() {
	for p.advance() {
		commandType, err := p.getCommandType()
		if err != nil {
			fmt.Printf("Invalid Insturction detected \n")
			fmt.Printf("%s was the currentCommand\n", p.currentCommand)
			break
		}
		// now we have a valid commadntype
		fmt.Printf("current Insturction -> %s\n", p.currentCommand)
		fmt.Printf("commandType ->%d\n", commandType)
		if commandType == ACommand || commandType == LCommand {
			fmt.Println(p.getSymbol(commandType))
		} else {
			fmt.Printf("dest -> %s\n", p.getDest())
			fmt.Printf("comp -> %s\n", p.getComp())
			fmt.Printf("jump -> %s\n", p.getJump())
		}
	}
	fmt.Println("Done Parsing")
}
