package engine

import (
	"fmt"
	"nand2tetris/compiler-I/lexer"
)

func (e *Engine) writeIndent() {
	for i := 0; i < e.indent; i++ {
		e.Writer.WriteString(" ")
	}
}

func (e *Engine) writeLine(s string) {
	e.writeIndent()
	e.Writer.WriteString(s)
	e.Writer.WriteString("\n")
}

func (e *Engine) writeToken(tokenType lexer.TokenType, lexeme string) {
	tType := tokenTypeToString[tokenType]
	e.writeLine("<" + tType + "> " + lexeme + " </" + tType + ">")
}

// next advances to the next token, erroring if none remain
func (e *Engine) next() (lexer.TokenType, string, error) {
	if !e.lex.HasMoreTokens() {
		return 0, "", fmt.Errorf("No more tokens")
	}
	e.lex.Advance()
	tokenType, lexeme := e.lex.GetTokenTypeAndLexeme()
	return tokenType, lexeme, nil
}
