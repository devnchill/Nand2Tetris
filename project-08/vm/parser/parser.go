package parser

import (
	"bufio"
	"nand2tetris/vm/util"
	"os"
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

func (p *Parser) Close() {
	p.filePointer.Close()
}
