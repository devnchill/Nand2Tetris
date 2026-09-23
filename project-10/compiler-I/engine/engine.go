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

func (e *Engine) compileSubroutine() {
}
func (e *Engine) compileParameterList() {
}
func (e *Engine) compileVarDec() {
}
func (e *Engine) compileStatements() {
}
func (e *Engine) compileDo() {
}
func (e *Engine) compileLet() {
}
func (e *Engine) compileWhile() {
}
func (e *Engine) compileReturn() {
}
func (e *Engine) compileIf() {
}
func (e *Engine) compileExpression() {
}
func (e *Engine) compileTerm() {
}
func (e *Engine) compileExpressionList() {
}
