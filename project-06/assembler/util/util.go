// Package util would implement utility methods to be used in other packages
package util

import (
	"strconv"
	"unicode/utf8"
)

func reverse(s string) string {
	size := len(s)
	buf := make([]byte, size)
	for start := 0; start < size; {
		r, n := utf8.DecodeRuneInString(s[start:])
		start += n
		utf8.EncodeRune(buf[size-start:], r)
	}
	return string(buf)
}

func DecimalToBinary(instruction string) (string, error) {
	decimalNumber, err := strconv.Atoi(instruction[1:])
	var binary string
	if err != nil {
		return "", err
	}
	for decimalNumber > 0 {
		if decimalNumber%2 == 0 {
			binary = binary + string('0')
		} else {
			binary = binary + string('1')
		}
		decimalNumber = decimalNumber / 2
	}
	return reverse(binary), nil
}
