package main

import (
	"fmt"
	"strings"
)

func renderLines(text string, banner map[rune][]string) []string {
	var result []string
	for i := 0; i <= 7; i++ {
		var builder strings.Builder
		for _, char := range text {
			charChange, ok := banner[rune(char)]
			if !ok {
				fmt.Println("map key no found")
				return nil

			}
			builder.WriteString(charChange[i])
		}
		result = append(result, builder.String())
	}
	return result
}
