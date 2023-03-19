package src

import (
	"bufio"
	"io"
	"os"
	"strings"
)

func ReadFile(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	read := bufio.NewReader(file)
	var lines []string
	for {
		line, err := read.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			return lines, err
		}
		lines = append(lines, strings.TrimSpace(line))
	}

	return lines, nil
}
