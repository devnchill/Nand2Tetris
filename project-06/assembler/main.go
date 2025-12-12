package main

import (
	"bufio"
	"fmt"
	"os"
)

func readFile(filePath string) {
	filePointer, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer filePointer.Close()
	scanner := bufio.NewScanner(filePointer)
	for scanner.Scan() {
		line := scanner.Text()
		for i, ch := range line {
			if (ch == '/' && rune(line[i+1]) == '/') || ch == ' ' {
				// ignoring all comments and empty line and empty space
				fmt.Println("found a comment exiting")
				break
			} else {
				fmt.Printf("%c", ch)
			}
		}
		fmt.Println()
	}
}

func main() {
	args := os.Args
	asmFilePath := args[1]
	readFile(asmFilePath)
}
