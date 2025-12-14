// Package util would implement utility methods to be used in other packages
package util

import "strconv"

func BinaryToDecimal(instruction string) (int64, error) {
	if i, err := strconv.ParseInt(instruction[1:], 2, 1); err != nil {
		return 0, err
	} else {
		return i, nil
	}
}
