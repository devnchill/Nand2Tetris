package engine

import (
	"bufio"
	"fmt"
	"nand2tetris/compiler-I/lexer"
	"os"
)

type Engine struct {
	indent         int
	inputFilePath  string
	lex            lexer.Lexer
	outputFilePath string
	writer         *bufio.Writer
}

var tokenTypeToString = [5]string{
	"keyword",
	"symbol",
	"identifier",
	"integerConstant",
	"stringConstant",
}

func NewCompilationEngine(inputFilePath string, outputFilePath string) Engine {

	fp, err := os.Open(outputFilePath)
	if err != nil {
		fmt.Println("Error opening file to write" + outputFilePath)
	}
	return Engine{
		inputFilePath:  inputFilePath,
		outputFilePath: outputFilePath,
		lex:            lexer.NewLexer(inputFilePath),
		writer:         bufio.NewWriter(fp),
	}
}

func (e *Engine) writeIndent() {
	for i := 0; i < e.indent; i++ {
		e.writer.WriteString(" ")
	}
}

func (e *Engine) writeLine(s string) {
	e.writeIndent()
	e.writer.WriteString(s)
	e.writer.WriteString("\n")
}

func (e *Engine) compileClass() error {
	if e.lex.HasMoreTokens() {
		e.lex.Advance()
	}
	if e.lex.GetTokenType() != lexer.Keyword || e.lex.GetLexeme() != "class" {
		return fmt.Errorf(
			"expected keyword 'class',got %s %q", tokenTypeToString[e.lex.GetTokenType()], e.lex.GetLexeme(),
		)
	}

	e.writeLine("<class>")
	e.indent++

	tType := tokenTypeToString[e.lex.GetTokenType()]
	token := e.lex.GetLexeme()
	e.writeLine("<" + tType + ">" + token + "</" + tType + ">")

	if e.lex.HasMoreTokens() {
		e.lex.Advance()
	}
	tType = tokenTypeToString[e.lex.GetTokenType()]
	token = e.lex.GetLexeme()
	e.writeLine("<" + tType + ">" + token + "</" + tType + ">")

	if e.lex.HasMoreTokens() {
		e.lex.Advance()
	}
	tType = tokenTypeToString[e.lex.GetTokenType()]
	token = e.lex.GetLexeme()
	e.writeLine("<" + tType + ">" + token + "</" + tType + ">")

	e.compileSubroutine()

	e.indent--
	e.writeLine("</class>")

	return nil
}

func (e *Engine) compileClassVarDec() {
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
