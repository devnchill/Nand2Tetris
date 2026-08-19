package lexer

import (
	"log"
	"os"
	"strings"
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

func (l *Lexer) isWhiteSpace(char byte) bool {
	return char == ' ' || char == '\n' || char == '\t' || char == '\r'
}

func (l *Lexer) HasMoreTokens() bool {
	currPointer := l.pointer
	for currPointer < len(l.source) && l.isWhiteSpace(l.source[currPointer]) {
		currPointer++
	}
	if currPointer == len(l.source) {
		return false
	}
	if l.source[currPointer] == '/' && currPointer+1 < len(l.source) && l.source[currPointer+1] == '/' {
		currPointer = l.skipComments()
	}
	return currPointer < len(l.source)
}

// only to be called if HasMoreTokens is true
func (l *Lexer) Advance() {
	for {
		for l.pointer < len(l.source) && l.isWhiteSpace(l.source[l.pointer]) {
			l.pointer++
		}
		if l.source[l.pointer] == '/' && l.source[l.pointer+1] == '/' {
			l.pointer = l.skipComments()
		} else {
			break
		}
	}
	var sb strings.Builder
	if l.source[l.pointer] == '"' {
		for l.pointer < len(l.source) && l.source[l.pointer] != '"' {
			sb.WriteByte(l.source[l.pointer])
			l.pointer++
		}
		l.currentToken.Lexeme = sb.String()
		l.currentToken.Type = StringConstant
	} else if l.source[l.pointer] >= '0' && l.source[l.pointer] <= '9' {
		for l.pointer < len(l.source) && l.source[l.pointer] >= '0' && l.source[l.pointer] <= '9' {
			sb.WriteByte(l.source[l.pointer])
			l.pointer++
		}
		l.currentToken.Lexeme = sb.String()
		l.currentToken.Type = IntegerConstant
	}
}

func (l *Lexer) skipComments() int {
	/*
		 only call when we know it is a comment
			pointer would be at first '/' everytime this is called
	*/
	// move cursor to second `/`
	currPointer := l.pointer + 1
	for currPointer < len(l.source) && l.source[currPointer] != '\n' {
		currPointer++
	}
	return currPointer
}

func (l *Lexer) GetTokenType() TokenType {
	return l.currentToken.Type
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
