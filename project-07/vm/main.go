package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: vm -f <file.vm ...> -d <directory ...>")
		os.Exit(1)
	}
	var fileFlag, dirFlag bool
	var dirFlagIdx int

	// NOTE: extract it into a helper function
	if os.Args[1] == "-f" {
		fileFlag = true
		for i := range os.Args[2:] {
			if os.Args[i] != "-d" {
				continue
			}
			if i == len(os.Args)-1 {
				fmt.Println("invalid syntax")
				fmt.Println("no directory provided")
				fmt.Println("Usage: vm -f <file.vm> -d <directory>")
				os.Exit(1)
			}
			dirFlagIdx = i
			dirFlag = true
			break
		}
	} else if os.Args[1] == "-d" {
		fileFlag = false
		dirFlag = true
	} else {
		fmt.Println("invalid syntax")
		fmt.Println("Usage: vm -f <file.vm> -d <directory>")
		os.Exit(1)
	}
	var files, directories []string
	if fileFlag && !dirFlag {
		files = os.Args[2:]
	} else if !fileFlag && dirFlag {
		directories = os.Args[2:]
	} else {
		// both files and directories are provided
		files = os.Args[2 : dirFlagIdx-1]
		directories = os.Args[dirFlagIdx+1:]
	}

	fmt.Println(files)
	fmt.Println(directories)

	for _, dir := range directories {
		err := filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() && filepath.Ext(path) == ".vm" {
				files = append(files, path)
			}
			return nil
		})
		if err != nil {
			fmt.Printf("Error reading directory: %s", dir)
		}
	}

	for _, file := range files {
		fmt.Println(file)
		if filepath.Ext(file) != ".vm" {
			fmt.Println("Error: input file must have .vm extension")
			os.Exit(1)
		}
		generateASM(file)
	}
}
