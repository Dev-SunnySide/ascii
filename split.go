package main

import "strings"

func splitLines(input string) []string {
	split := strings.Split(input, "\\n")
	return split
}
