package lexer

type Token struct {
	Type   TokenType
	Lexeme string
}

type TokenType int

const (
	Keyword TokenType = iota
	Symbol
	Identifier
	IntegerConstant
	StringConstant
)
