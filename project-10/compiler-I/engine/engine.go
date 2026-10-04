package engine

import (
	"bufio"
	"fmt"
	"nand2tetris/compiler-I/lexer"
	"os"
)

type Engine struct {
	Writer         *bufio.Writer
	Fp             *os.File
	indent         int
	inputFilePath  string
	lex            lexer.Lexer
	outputFilePath string
}

var tokenTypeToString = [5]string{
	"keyword",
	"symbol",
	"identifier",
	"integerConstant",
	"stringConstant",
}

func NewCompilationEngine(inputFilePath string, outputFilePath string) Engine {

	fp, err := os.Create(outputFilePath)
	if err != nil {
		fmt.Println("Error opening file to write" + outputFilePath)
	}
	return Engine{
		Writer:         bufio.NewWriter(fp),
		Fp:             fp,
		inputFilePath:  inputFilePath,
		lex:            lexer.NewLexer(inputFilePath),
		outputFilePath: outputFilePath,
	}
}

func (e *Engine) Close() {
	e.Writer.Flush()
	e.Fp.Close()
}

func (e *Engine) compileDo() error {
	return fmt.Errorf("compileDo: not implemented")
}

func (e *Engine) compileLet() error {
	return fmt.Errorf("compileLet: not implemented")
}

func (e *Engine) compileWhile() error {
	return fmt.Errorf("compileWhile: not implemented")
}

func (e *Engine) compileReturn() error {
	return fmt.Errorf("compileReturn: not implemented")
}

func (e *Engine) compileIf() error {
	return fmt.Errorf("compileIf: not implemented")
}

func (e *Engine) compileExpression() error {
	return fmt.Errorf("compileExpression: not implemented")
}

func (e *Engine) compileTerm() error {
	return fmt.Errorf("compileTerm: not implemented")
}

func (e *Engine) compileExpressionList() error {
	return fmt.Errorf("compileExpressionList: not implemented")
}
