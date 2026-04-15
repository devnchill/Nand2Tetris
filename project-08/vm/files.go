package main

import (
	"fmt"
	"io/fs"
	"path/filepath"
)

func collectVMFiles(directories []string) ([]string, error) {
	var files []string

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
			return nil, fmt.Errorf("error reading directory %s: %w", dir, err)
		}
	}

	return files, nil
}

func validateVMFiles(files []string) error {
	for _, file := range files {
		if filepath.Ext(file) != ".vm" {
			return fmt.Errorf("invalid file: %s (must be .vm)", file)
		}
	}
	return nil
}
