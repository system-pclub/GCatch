package util

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
)

func ReadFileLine(filename string, n int) (string, error) {

	if n < 1 {
		return "", fmt.Errorf("invalid request: line %d", n)
	}

	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}

	defer f.Close()
	bf := bufio.NewReader(f)
	var line string
	for numLine := 0; numLine < n; numLine ++ {
		line, err = bf.ReadString('\n')
		if err == io.EOF {
			switch numLine {
			case 0:
				return "", errors.New("no lines in file")
			case 1:
				return "", errors.New("only 1 line")
			default:
				return "", fmt.Errorf("only %d lines", numLine)
			}
		}
		if err != nil {
			return "", err
		}
	}
	if line == "" {
		return "", fmt.Errorf("line %d empty", n)
	}
	return line, nil
}