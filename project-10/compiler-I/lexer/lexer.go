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

func (l *Lexer) Advance() {
	for l.pointer < len(l.source) && l.isWhiteSpace(l.source[l.pointer]) {
		l.pointer++
	}
	if l.source[l.pointer] == '/' && l.source[l.pointer+1] == '/' {
		l.pointer = l.skipComments()
	}
}

func (l *Lexer) skipComments() int {
	/*
		 only call when we know it is a comment
			pointer would be at first '/' everytime this is called
	*/
	// move cursor to second `/`
	curr := l.pointer + 1
	for curr < len(l.source) && l.source[curr] != '\n' {
		curr++
	}
	return curr
}

func (l *Lexer) peek() byte {
	if l.pointer+1 >= len(l.source) {
		return 0
	}
	return l.source[l.pointer+1]
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
