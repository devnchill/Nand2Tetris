package main

import (
	"bufio"
	"os"
	"strings"
)

func needsBootstrap(files []string) bool {
	for _, file := range files {
		fp, err := os.Open(file)
		if err != nil {
			continue
		}
		scanner := bufio.NewScanner(fp)
		for scanner.Scan() {
			line := scanner.Text()
			if idx := strings.Index(line, "//"); idx != -1 {
				line = strings.TrimSpace(line[:idx])
			}
			fields := strings.Fields(line)
			if len(fields) >= 2 && fields[0] == "function" && fields[1] == "Sys.init" {
				fp.Close()
				return true
			}
		}
		fp.Close()
	}
	return false
}
