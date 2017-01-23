package parser

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strings"
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

func Exists(fileName string) bool {
	_, err := os.Stat(fileName)
	return err == nil || os.IsExist(err)
}

func Mkdir(path string) bool {
	err := os.MkdirAll(path, os.ModePerm)
	return err == nil
}

func CreateFile(path, name string) (*os.File, error) {
	if !Exists(path) {
		Mkdir(path)
	}
	return os.Create(path + "/" + name)
}

func FileName(name string) string {
	return path.Base(name)
}

func IsDir(fileName string) bool {
	stats, err := os.Stat(fileName)
	return err == nil && stats.IsDir()
}

func IsProtoFile(fileName string) bool {
	return strings.HasSuffix(fileName, ".proto")
}
