package main

import (
	parser "nand2tetris/vm/parser"
	"nand2tetris/vm/translator"
)

func generateASM(p *parser.Parser, t *translator.Translator, file string) {

	for {
		p.Advance()

		if !p.HasMoreCommand {
			break
		}

		switch p.CurrentCommandType {
		case parser.C_ARITHMETIC:
			{
				t.WriteArithmeticCommand(p.CurrentCommand)
			}
		case parser.C_PUSH, parser.C_POP:
			{
				arg1 := p.GetFirstArg()
				arg2 := p.GetSecondArg()
				t.WritePushPopCommand(p.CurrentCommand, p.CurrentCommandType, arg1, arg2, file)
			}
		}
	}
}
