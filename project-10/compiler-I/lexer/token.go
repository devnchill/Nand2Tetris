package lexer

type Token struct {
	tokenType TokenType
	lexeme    string
}

type TokenType int

const (
	Keyword TokenType = iota
	Symbol
	Identifier
	IntegerConstant
	StringConstant
)

var keywords = map[string]bool{
	"class":       true,
	"constructor": true,
	"function":    true,
	"method":      true,
	"field":       true,
	"static":      true,
	"var":         true,
	"int":         true,
	"char":        true,
	"boolean":     true,
	"void":        true,
	"true":        true,
	"false":       true,
	"null":        true,
	"this":        true,
	"let":         true,
	"do":          true,
	"if":          true,
	"else":        true,
	"while":       true,
	"return":      true,
}

var symbols = map[byte]bool{
	'(': true,
	')': true,
	'{': true,
	'}': true,
	'[': true,
	']': true,
	'.': true,
	',': true,
	'+': true,
	'-': true,
	'*': true,
	'/': true,
	'&': true,
	'|': true,
	'<': true,
	'>': true,
	'=': true,
	'~': true,
}
