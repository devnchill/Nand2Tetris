package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: vm -f <file.vm ...> -d <directory ...>")
		os.Exit(1)
	}

	files, directories := parseArgs(os.Args)
	fmt.Println(files)
	fmt.Println(directories)

	dirFiles, err := collectVMFiles(directories)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	files = append(files, dirFiles...)

	if err := validateVMFiles(files); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, file := range files {
		fmt.Println(file)
		generateASM(file)
	}
}
