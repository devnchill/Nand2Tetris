package parser

import (
	"nand2tetris/vm/util"
	"strings"
)

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
