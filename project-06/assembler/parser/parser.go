// Package parser is responsible for parsing the asm commmands
package parser

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/devnchill/Nand2Tetris/project-06/assembler/translator"
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
		decimalNumber, err := strconv.Atoi(p.currentCommand[1:])
		val := fmt.Sprintf("%b", decimalNumber)
		if err != nil {
			return "", err
		}
		return val, nil
	}
	// for LCommand
	return strings.Split((strings.Split(p.currentCommand, ")")[0]), "(")[1], nil
}

// should only be called for C instructions
// Format of CInstruction = getDest = comp;jump
func (p *Parser) hasDest() bool {
	return strings.Contains(p.currentCommand, "=")
}

func (p *Parser) hasJump() bool {
	return strings.Contains(p.currentCommand, ";")
}

// should only be called for C instructions and only when dest is present
func (p *Parser) getDest() string {
	return strings.TrimSpace(strings.Split(p.currentCommand, "=")[0])
}

// should only be called for C instructions
func (p *Parser) getComp() string {
	if p.hasDest() && p.hasJump() {
		return strings.Split(strings.Split(p.currentCommand, ";")[0], "=")[1]
	} else if p.hasDest() && !p.hasJump() {
		return strings.Split(p.currentCommand, "=")[1]
	} else if !p.hasDest() && p.hasJump() {
		return strings.Split(p.currentCommand, ";")[0]
	} else {
		return p.currentCommand
	}
}

// should only be called for C instructions and only when jump is present
func (p *Parser) getJump() string {
	if strings.Contains(p.currentCommand, ";") {
		return strings.Split(p.currentCommand, ";")[1]
	}
	return ""
}

func (p *Parser) Parse() {
	t := translator.NewTranslator()
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
			if p.hasDest() {
				fmt.Printf("dest -> %s\n", p.getDest())
				destInBinary, err := t.TranslateDest(p.getDest())
				if err != nil {
					fmt.Println(err)
					break
				}
				fmt.Printf("dest in binary -> %s\n", destInBinary)
			}
			fmt.Printf("comp -> %s\n", p.getComp())
			compInBinary, err := t.TranslateComp(p.getComp())
			if err != nil {
				fmt.Println(err)
				break
			}
			fmt.Printf("comp in binary -> %s\n", compInBinary)

			if p.hasJump() {
				fmt.Printf("jump -> %s\n", p.getJump())
				jumpInBinary, err := t.TranslateJump(p.getJump())
				if err != nil {
					fmt.Println(err)
					break
				}
				fmt.Printf("jump in binary -> %s\n", jumpInBinary)

			}
		}
	}
	fmt.Println("Done Parsing")
}
