package engine

import (
	"fmt"
	"nand2tetris/compiler-I/lexer"
	"slices"
	"strconv"
	"strings"
)

var xmlEscaper = strings.NewReplacer(
	"&", "&amp;",
	"<", "&lt;",
	">", "&gt;",
	`"`, "&quot;",
	"'", "&apos;",
)

func (e *Engine) writeIndent() {
	for i := 0; i < e.indent; i++ {
		e.Writer.WriteString("  ")
	}
}

func (e *Engine) writeLine(s string) {
	e.writeIndent()
	e.Writer.WriteString(s)
	e.Writer.WriteString("\n")
}

func (e *Engine) writeToken(tokenType lexer.TokenType, lexeme string) {
	tType := tokenTypeToString[tokenType]
	e.writeLine("<" + tType + "> " + xmlEscaper.Replace(lexeme) + " </" + tType + ">")
}

func (e *Engine) next() (lexer.TokenType, string, error) {
	if !e.lex.HasMoreTokens() {
		return 0, "", fmt.Errorf("No more tokens")
	}
	e.lex.Advance()
	tokenType, lexeme := e.lex.GetTokenTypeAndLexeme()
	return tokenType, lexeme, nil
}

func (e *Engine) peek() (lexer.TokenType, string, error) {
	if !e.lex.HasMoreTokens() {
		return 0, "", fmt.Errorf("No more tokens")
	}
	tokenType, lexeme := e.lex.Peek()
	return tokenType, lexeme, nil
}

func (e *Engine) peekIsKeyword(lexemes ...string) bool {
	tokenType, lexeme, err := e.peek()
	if err != nil {
		return false
	}
	return tokenType == lexer.Keyword && contains(lexemes, lexeme)
}

func (e *Engine) peekIsSymbol(symbols ...string) bool {
	tokenType, lexeme, err := e.peek()
	if err != nil {
		return false
	}
	return tokenType == lexer.Symbol && contains(symbols, lexeme)
}

// peekIsStatement reports whether the next token can start a statement
func (e *Engine) peekIsStatement() bool {
	tokenType, lexeme, err := e.peek()
	if err != nil {
		return false
	}
	if tokenType == lexer.Symbol && lexeme == "{" {
		return true
	}
	return tokenType == lexer.Keyword && contains([]string{"let", "if", "while", "do", "return"}, lexeme)
}

func contains(lexemes []string, lexeme string) bool {
	return slices.Contains(lexemes, lexeme)
}

func (e *Engine) expectKeyword(lexemes ...string) error {
	tokenType, lexeme, err := e.next()
	if err != nil {
		return err
	}
	if tokenType != lexer.Keyword || !contains(lexemes, lexeme) {
		return fmt.Errorf("expected keyword (%s), got %s %s", strings.Join(lexemes, " | "), tokenTypeToString[tokenType], lexeme)
	}
	e.writeToken(tokenType, lexeme)
	return nil
}

func (e *Engine) expectIdentifier() error {
	tokenType, lexeme, err := e.next()
	if err != nil {
		return err
	}
	if tokenType != lexer.Identifier {
		return fmt.Errorf("expected tokenType identifier, got %s %s", tokenTypeToString[tokenType], lexeme)
	}
	e.writeToken(tokenType, lexeme)
	return nil
}

func (e *Engine) expectType() error {
	tokenType, lexeme, err := e.next()
	if err != nil {
		return err
	}
	isPrimitive := tokenType == lexer.Keyword && contains([]string{"int", "char", "boolean"}, lexeme)
	if !isPrimitive && tokenType != lexer.Identifier {
		return fmt.Errorf("expected type (int | char | boolean | className), got %s %s", tokenTypeToString[tokenType], lexeme)
	}
	e.writeToken(tokenType, lexeme)
	return nil
}

func (e *Engine) expectTypeOrVoid() error {
	tokenType, lexeme, err := e.next()
	if err != nil {
		return err
	}
	isType := tokenType == lexer.Identifier ||
		(tokenType == lexer.Keyword && contains([]string{"int", "char", "boolean", "void"}, lexeme))
	if !isType {
		return fmt.Errorf("expected type (int | char | boolean | void | className), got %s %s", tokenTypeToString[tokenType], lexeme)
	}
	e.writeToken(tokenType, lexeme)
	return nil
}

func (e *Engine) expectSymbol(symbols ...string) error {
	tokenType, lexeme, err := e.next()
	if err != nil {
		return err
	}
	if tokenType != lexer.Symbol || !contains(symbols, lexeme) {
		return fmt.Errorf("expected symbol (%s), got %s %s", strings.Join(symbols, " | "), tokenTypeToString[tokenType], lexeme)
	}
	e.writeToken(tokenType, lexeme)
	return nil
}

// expectExpressionListIfParenthesized compiles the '( expressionList )' of a
// subroutine call, a varName that is called without arguments has no parens
func (e *Engine) expectExpressionListIfParenthesized() error {
	if !e.peekIsSymbol("(") {
		return nil
	}

	err := e.expectSymbol("(")
	if err != nil {
		return err
	}

	err = e.compileExpressionList()
	if err != nil {
		return err
	}

	return e.expectSymbol(")")
}

/*
'[' expression ']'

	the token is already consumed when this is called,
	so it only validates the contents of the brackets
*/
func (e *Engine) expectInsideBrackets() error {
	err := e.expectSymbol("[")
	if err != nil {
		return err
	}

	err = e.compileExpression()
	if err != nil {
		return err
	}

	return e.expectSymbol("]")
}

// varName: 'this' | identifier '[' expression ']'?
func (e *Engine) expectVarName() error {
	tokenType, lexeme, err := e.next()
	if err != nil {
		return err
	}

	isVarName := tokenType == lexer.Identifier || (tokenType == lexer.Keyword && lexeme == "this")
	if !isVarName {
		return fmt.Errorf("expected varName (identifier | this), got %s %s", tokenTypeToString[tokenType], lexeme)
	}
	e.writeToken(tokenType, lexeme)

	if e.peekIsSymbol("[") {
		return e.expectInsideBrackets()
	}

	return nil
}

// expectIntegerConstant validates and writes the lexeme of an already
// consumed integerConstant token, only 15-bit unsigned values fit in a word
func (e *Engine) expectIntegerConstant(lexeme string) error {
	value, err := strconv.Atoi(lexeme)
	if err != nil {
		return fmt.Errorf("invalid integer constant %s", lexeme)
	}
	if value < 0 || value > 32767 {
		return fmt.Errorf("integer constant %d out of range (0..32767)", value)
	}
	e.writeToken(lexer.IntegerConstant, lexeme)
	return nil
}

// expectStringConstant validates and writes the lexeme of an already consumed
// stringConstant token, only printable characters fit in the character set
func (e *Engine) expectStringConstant(lexeme string) error {
	for _, char := range lexeme {
		if char < ' ' || char > '~' {
			return fmt.Errorf("illegal character '%c' (0x%02x) in string constant", char, char)
		}
	}
	e.writeToken(lexer.StringConstant, lexeme)
	return nil
}
