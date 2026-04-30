package translator

import (
	"bufio"
	"fmt"
	"nand2tetris/vm/util"
	"os"
)

type Translator struct {
	filePointer *os.File
	writer      *bufio.Writer
}

func NewTranslator(file string) *Translator {
	fmt.Printf("Creating new file: %s\n", file)
	fp, err := os.Create(file)
	util.Check(err)
	writer := bufio.NewWriter(fp)
	return &Translator{
		writer:      writer,
		filePointer: fp,
	}
}

func (t *Translator) Close() {
	t.writer.Flush()
	t.filePointer.Close()
}
