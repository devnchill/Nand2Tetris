package lexer

type TokenType int

const (
	Keyword TokenType = iota
	Symbol
	Identifier
	IntegerConstant
	StringConstant
)
