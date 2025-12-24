// Package translator is responsible for translating parsed symbolic hack to machine/binary code
package translator

import "errors"

type Translator struct {
	destToBinary map[string]string
	compToBinary map[string]string
	jumpToBinary map[string]string
}

func NewTranslator() Translator {
	return Translator{
		destToBinary: map[string]string{
			"":    "000",
			"M":   "001",
			"D":   "010",
			"MD":  "011",
			"A":   "100",
			"AM":  "101",
			"AD":  "110",
			"AMD": "111",
		},
		jumpToBinary: map[string]string{
			"":    "000",
			"JGT": "001",
			"JEQ": "010",
			"JGE": "011",
			"JLT": "100",
			"JNE": "101",
			"JLE": "110",
			"JMP": "111",
		},
	}
}

func (t Translator) TranslateDest(dest string) (string, error) {
	val, ok := t.destToBinary[dest]
	if ok {
		return val, nil
	}
	return "", errors.New("invalid dest symbolic code: dest was" + dest)
}

func (t Translator) TranslateComp(comp string) (string, error) {
	return "", nil
}

func (t Translator) TranslateJump(jump string) (string, error) {
	val, ok := t.jumpToBinary[jump]
	if ok {
		return val, nil
	}
	return "", errors.New("invalid jump symbolic code: jump was" + jump)
}
