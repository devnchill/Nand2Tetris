package engine

import (
	"fmt"
	"nand2tetris/compiler-I/lexer"
	"slices"
	"strings"
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
	e.writeLine("<" + tType + "> " + lexeme + " </" + tType + ">")
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

func (e *Engine) peekIsSymbol(symbol string) bool {
	tokenType, lexeme, err := e.peek()
	if err != nil {
		return false
	}
	return tokenType == lexer.Symbol && lexeme == symbol
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

func (e *Engine) expectSymbol(symbol string) error {
	tokenType, lexeme, err := e.next()
	if err != nil {
		return err
	}
	if tokenType != lexer.Symbol || lexeme != symbol {
		return fmt.Errorf("expected symbol '%s', got %s %s", symbol, tokenTypeToString[tokenType], lexeme)
	}
	e.writeToken(tokenType, lexeme)
	return nil
}
