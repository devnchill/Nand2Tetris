package util

import (
	"fmt"
	"os"
)

func ParseArgs(args []string) ([]string, []string) {
	if len(args) < 2 {
		fmt.Println("Usage: vm -f <file.vm ...> -d <directory ...>")
		os.Exit(1)
	}

	var fileFlag, dirFlag bool
	var dirFlagIdx int

	if args[1] == "-f" {
		fileFlag = true
		for i, arg := range args[2:] {
			if arg != "-d" {
				continue
			}
			if i == len(args)-1 {
				fmt.Println("invalid syntax")
				fmt.Println("no directory provided")
				os.Exit(1)
			}
			dirFlagIdx = i
			dirFlag = true
			break
		}
	} else if args[1] == "-d" {
		dirFlag = true
	} else {
		fmt.Println("invalid syntax")
		os.Exit(1)
	}

	var files, directories []string

	if fileFlag && !dirFlag {
		files = args[2:]
	} else if !fileFlag && dirFlag {
		directories = args[2:]
	} else {
		files = args[2 : dirFlagIdx-1]
		directories = args[dirFlagIdx+1:]
	}

	return files, directories
}
