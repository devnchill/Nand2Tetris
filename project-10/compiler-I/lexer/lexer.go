package lexer

import (
	"log"
	"os"
	"strconv"
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
	return l.skipTrivia(l.pointer) < len(l.source)
}

// skipTrivia advances pointer past whitespace and comments
func (l *Lexer) skipTrivia(pointer int) int {
	for {
		for pointer < len(l.source) && l.isWhiteSpace(l.source[pointer]) {
			pointer++
		}
		if pointer == len(l.source) {
			return pointer
		}
		if l.source[pointer] == '/' && pointer+1 < len(l.source) && (l.source[pointer+1] == '*' || l.source[pointer+1] == '/') {
			pointer = l.skipComments(pointer, l.source[pointer+1] == '*')
		} else {
			break
		}
	}
	return pointer
}

// only to be called if HasMoreTokens is true
func (l *Lexer) Advance() {
	l.pointer = l.skipTrivia(l.pointer)
	var sb strings.Builder
	if l.source[l.pointer] == '"' {
		l.pointer++
		for l.pointer < len(l.source) && l.source[l.pointer] != '"' {
			sb.WriteByte(l.source[l.pointer])
			l.pointer++
		}
		l.pointer++
		l.currentToken.tokenType = StringConstant
	} else if l.isDigit() {
		for l.pointer < len(l.source) && l.isDigit() {
			sb.WriteByte(l.source[l.pointer])
			l.pointer++
		}
		l.currentToken.tokenType = IntegerConstant
	} else if symbols[l.source[l.pointer]] {
		l.currentToken.tokenType = Symbol
		sb.WriteByte(l.source[l.pointer])
		l.pointer++
	} else if l.isUnderScore() || l.isAlphabet() {
		for l.isDigit() || l.isAlphabet() || l.isUnderScore() {
			sb.WriteByte(l.source[l.pointer])
			l.pointer++
		}
		if keywords[sb.String()] {
			l.currentToken.tokenType = Keyword
		} else {
			l.currentToken.tokenType = Identifier
		}
	}
	l.currentToken.lexeme = sb.String()
}

func (l *Lexer) isAlphabet() bool {
	return (l.source[l.pointer] >= 'a' && l.source[l.pointer] <= 'z') || (l.source[l.pointer] >= 'A' && l.source[l.pointer] <= 'Z')
}

func (l *Lexer) isDigit() bool {
	return l.source[l.pointer] >= '0' && l.source[l.pointer] <= '9'
}
func (l *Lexer) isUnderScore() bool {
	return l.source[l.pointer] == '_'
}

func (l *Lexer) skipComments(currPointer int, isMultilineComment bool) int {
	/*
		 only call when we know it is a comment
			pointer would be at first '/' everytime this is called
	*/
	if !isMultilineComment {
		// move cursor to second `/`
		for currPointer < len(l.source) && l.source[currPointer] != '\n' {
			currPointer++
		}
		if len(l.source)-currPointer >= 2 {
			// skip '\n'
			currPointer += 1
		}
	} else {
		// move cursor to second `*`
		for currPointer+1 < len(l.source) && !(l.source[currPointer] == '*' && l.source[currPointer+1] == '/') {
			currPointer++
		}
		currPointer += 2
	}
	return currPointer
}

func (l *Lexer) getLexeme() string {
	return l.currentToken.lexeme
}

func (l *Lexer) getTokenType() TokenType {
	return l.currentToken.tokenType
}

func (l *Lexer) GetTokenTypeAndLexeme() (TokenType, string) {
	return l.getTokenType(), l.getLexeme()
}

/* below methods would read `currentToken` stored in Lexer
   and perform operations on that to extract desired token
*/

func (l *Lexer) GetKeyword() string {
	return l.currentToken.lexeme
}

func (l *Lexer) GetSymbol() byte {
	return l.currentToken.lexeme[0]
}

func (l *Lexer) GetIdentifier() string {
	return l.currentToken.lexeme
}

func (l *Lexer) GetIntVal() int {
	/* technically it shouldn't throw error
	reason being, we only append byte to lexeme if byte is a digit
	*/

	val, err := strconv.Atoi(l.currentToken.lexeme)
	if err != nil {
		log.Panicf("invalid lexeme -> %s for tokenType -> %d", l.currentToken.lexeme, l.currentToken.tokenType)
	}
	return val
}

func (l *Lexer) GetStringVal() string {
	return l.currentToken.lexeme
}
