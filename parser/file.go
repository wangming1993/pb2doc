package parser

import (
	"bufio"
	"fmt"
	"os"
)

func ReadFile(filePath string) []string {
	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if "" != scanner.Text() {
			lines = append(lines, scanner.Text())
		}

	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	return lines
}
