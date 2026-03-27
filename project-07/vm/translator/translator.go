package translator

import (
	"bufio"
	"nand2tetris/vm/util"
	"os"
)

type Translator struct {
	filePointer *os.File
	writer      *bufio.Writer
}

func NewTranslator(file string) *Translator {
	fp, err := os.Create(file)
	util.Check(err)
	defer fp.Close()
	writer := bufio.NewWriter(fp)
	return &Translator{
		writer:      writer,
		filePointer: fp,
	}
}

func (t *Translator) writeArithmeticCommand(command string) {
}

func (t *Translator) writePushPopCommand(command string) {
}
func (t *Translator) close() {
	t.writer.Flush()
	t.filePointer.Close()
}
