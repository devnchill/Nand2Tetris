package engine

import (
	"fmt"
	"nand2tetris/compiler-I/lexer"
)

func (e *Engine) compileClassVarDec() error {
	// (static | field) type Varname (',' varname)* ';'
	if e.lex.HasMoreTokens() {
		e.lex.Advance()
	}
	if e.lex.GetTokenType() != lexer.Keyword || (e.lex.GetLexeme() != "static" || e.lex.GetLexeme() != "field") {
		return fmt.Errorf("expected static or field keyword, got %s", tokenTypeToString[e.lex.GetTokenType()])
	}
	tType := tokenTypeToString[e.lex.GetTokenType()]
	token := e.lex.GetLexeme()
	e.writeLine("<" + tType + "> " + token + " </" + tType + ">")

	if e.lex.HasMoreTokens() {
		e.lex.Advance()
	}
	if e.lex.GetTokenType() != lexer.Keyword || (e.lex.GetLexeme() != "int" || e.lex.GetLexeme() != "char" || e.lex.GetLexeme() != "boolean") {
		return fmt.Errorf("expected keyword wity lexeme int | char | boolean , got %s %s", tokenTypeToString[e.lex.GetTokenType()], e.lex.GetLexeme())
	}

	tType = tokenTypeToString[e.lex.GetTokenType()]
	token = e.lex.GetLexeme()
	e.writeLine("<" + tType + "> " + token + " </" + tType + ">")

	if e.lex.HasMoreTokens() {
		e.lex.Advance()
	}
	if e.lex.GetTokenType() != lexer.Identifier {
		return fmt.Errorf("exptected variable name")
	}
	tType = tokenTypeToString[e.lex.GetTokenType()]
	token = e.lex.GetLexeme()
	e.writeLine("<" + tType + "> " + token + " </" + tType + ">")

	for {
		if e.lex.HasMoreTokens() {
			e.lex.Advance()
		}
		if e.lex.GetLexeme() == "," && e.lex.GetTokenType() == lexer.Symbol {
			e.writeLine("<" + tType + "> " + token + " </" + tType + ">")
		} else {
			break
		}

		if e.lex.HasMoreTokens() {
			e.lex.Advance()
		}
		if e.lex.GetTokenType() == lexer.Identifier {
			e.writeLine("<" + tType + "> " + token + " </" + tType + ">")
		}
	}

	if e.lex.HasMoreTokens() {
		e.lex.Advance()
	}
	if e.lex.GetTokenType() == lexer.Symbol && e.lex.GetLexeme() == ";" {
		e.writeLine("<" + tType + "> " + token + " </" + tType + ">")
	}
	return nil
}
