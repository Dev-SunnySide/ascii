package main

import (
	"strings"
)

func generateAscii(text string, banner map[rune][]string) string {
	if text == "" {
		return ""
	}
	var builder strings.Builder
	split := splitLines(text)
	for _, line := range split {
		if line == "" {
			builder.WriteString("\n")
		} else {
			render := renderLines(line, banner)
			for _, lines := range render {
				builder.WriteString(lines)
				builder.WriteString("\n")
			}
		}
	}
	return builder.String()
}
