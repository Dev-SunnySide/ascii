package main

import (
	"errors"
	"os"
	"strings"
)

func LoadBanner(filename string) (map[rune][]string, error) {

	var (
		asciiCount = 32
		storageMap = make(map[rune][]string)
	)
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	splitChar := strings.Split(string(content), "\n")
	for i, line := range splitChar {

		splitChar[i] = strings.TrimRight(line, "\\r")

	}
	if len(splitChar) > 0 && splitChar[0] == "" {
		splitChar = splitChar[1:]
	}
	if len(splitChar) == 0 || len(splitChar) == 1 && splitChar[0] == "" {
		return nil, errors.New("empty file")
	}
	if len(splitChar) < 9 {
		return nil, errors.New("invalid character line")
	}
	for i := 0; i < len(splitChar); i += 9 {
		if i+8 > len(splitChar) {
			break
		}

		asciiChar := rune(asciiCount)
		asciiCount++
		asciiVal := splitChar[i : i+8]
		storageMap[asciiChar] = asciiVal
	}
	return storageMap, nil
}
