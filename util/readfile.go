package util

import (
	"bufio"
	"fmt"
	"os"
)

func ReadFile(filename string) []string {

	var lines []string

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	_ = file.Close()
	return lines
}
