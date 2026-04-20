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
				t.WritePushPopCommand(p.CurrentCommand, p.CurrentCommandType, arg1, arg2)
			}
		case parser.C_LABEL:
			{
				label := p.GetFirstArg()
				t.WriteLabel(label)
			}
		case parser.C_GOTO:
			{
				dest := p.GetFirstArg()
				t.WriteGoto(dest)
			}
		case parser.C_IF_GOTO:
			{
				label := p.GetFirstArg()
				t.WriteIfGoto(label)
			}
		case parser.C_FUNCTION:
			{
				funcName := p.GetFirstArg()
				numLocalVars := p.GetSecondArg()
				t.WriteFunction(funcName, numLocalVars)
			}
		case parser.C_RETURN:
			{
				t.WriteReturn()
			}
		case parser.C_CALL:
			{
				funcName := p.GetFirstArg()
				numArgs := p.GetSecondArg()
				t.WriteCall(funcName, numArgs)
			}
		}
	}
}
