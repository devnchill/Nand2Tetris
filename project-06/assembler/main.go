package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	asmFilePath := args[1]
	instructions, err := readFile(asmFilePath)
	if err != nil {
		fmt.Printf("%s", err)
		return
	}
	for _, in := range instructions {
		fmt.Printf("%s\n", in)
	}
}
