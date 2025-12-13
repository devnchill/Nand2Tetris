package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func readFile(filePath string) {
	filePointer, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer filePointer.Close()
	scanner := bufio.NewScanner(filePointer)
	var instructions []string
	for scanner.Scan() {
		currLine := scanner.Text()
		l1 := strings.Split(strings.TrimSpace(currLine), "//")[0]
		if strings.TrimSpace(l1) == "" {
			continue
		}
		l2 := strings.Join(strings.Split(l1, " "), "")
		instructions = append(instructions, l2)
	}
	for _, in := range instructions {
		fmt.Printf("%s\n", in)
	}
}

func main() {
	args := os.Args
	asmFilePath := args[1]
	readFile(asmFilePath)
}
