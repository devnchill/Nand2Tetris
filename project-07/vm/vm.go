package main

import (
	parser "nand2tetris/vm/parser"
	"nand2tetris/vm/translator"
	"path/filepath"
	"strings"
)

func generateASM(file string) {
	outFile := filepath.Join(filepath.Dir(file), strings.TrimPrefix(filepath.Base(file), filepath.Ext(file))) + ".asm"

	p := parser.NewParser(file)
	t := translator.NewTranslator(outFile)

	defer p.Close()
	defer t.Close()

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
