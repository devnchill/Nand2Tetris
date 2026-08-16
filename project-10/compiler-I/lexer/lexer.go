package lexer

import (
	"log"
	"os"
)

type Lexer struct {
	source       []byte
	pointer      int
	currentToken Token
}

func NewLexer(filepath string) Lexer {
	file, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}

	return Lexer{
		source:  file,
		pointer: 0,
	}
}

func (l *Lexer) isWhiteSpace() bool {
	char := l.source[l.pointer]
	return char == ' ' || char == '\n' || char == '\t' || char == '\r'
}

func (l *Lexer) HasMoreTokens() bool {
	return l.pointer < len(l.source)
}

func (l *Lexer) Advance() {
}

func (l *Lexer) GetTokenType() {
}

func (l *Lexer) GetKeyword() {
}

func (l *Lexer) GetSymbol() {
}

func (l *Lexer) GetIdentifier() {
}

func (l *Lexer) GetIntVal() {

}

func (l *Lexer) GetStringVal() {
}
