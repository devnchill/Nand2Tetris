package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func readFile(filePath string) ([]string, error) {
	filePointer, err := os.Open(filePath)
	if err != nil {
		return nil, err
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
	return instructions, nil
}

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
