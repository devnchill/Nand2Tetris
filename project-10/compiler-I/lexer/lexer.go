package lexer

import (
	"bufio"
	"os"
)

type Lexer struct {
	inputReader *bufio.Reader
}

func NewLexer(filepath string) Lexer {
	file, err := os.Open(filepath)
	if err != nil {
		panic("Unable to open the file")
	}

	defer file.Close()

	inputReader := bufio.NewReader(file)

	return Lexer{
		inputReader: inputReader,
	}
}
