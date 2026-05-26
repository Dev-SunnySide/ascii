package main

import (
	"strings"
)

var digitArt = [10][5]string{
	{ // 0
		" ___ ",
		"|   |",
		"| | |",
		"|   |",
		" ___ ",
	},
	{ // 1
		"   | ",
		"   | ",
		"   | ",
		"   | ",
		"   | ",
	},
	{ // 2
		" ___ ",
		"    |",
		" ___ ",
		"|    ",
		" ___ ",
	},
	{ // 3
		" ___ ",
		"    |",
		" ___ ",
		"    |",
		" ___ ",
	},
	{ // 4
		"|   |",
		"|___|",
		"    |",
		"    |",
		"    |",
	},
	{ // 5
		" ___ ",
		"|    ",
		" ___ ",
		"    |",
		" ___ ",
	},
	{ // 6
		" ___ ",
		"|    ",
		" ___ ",
		"|   |",
		" ___ ",
	},
	{ // 7
		" ___ ",
		"    |",
		"    |",
		"    |",
		"    |",
	},
	{ // 8
		" ___ ",
		"|   |",
		" ___ ",
		"|   |",
		" ___ ",
	},
	{ // 9
		" ___ ",
		"|   |",
		" ___ ",
		"    |",
		" ___ ",
	},
}

func StringToArt(input string) string {
	for _, char := range input {
		if (char < '0' || char > '9') && char != '\n' {
			return ""
		}
	}

	lines := strings.Split(input, "\n")
	var resultLines []string

	for _, line := range lines {
		if line == "" {
			continue
		}

		artLines := make([]string, 5)

		for i := 0; i < 5; i++ {
			var rowParts []string
			for _, char := range line {
				digitIdx := char - '0'
				rowParts = append(rowParts, digitArt[digitIdx][i])
			}
			artLines[i] = strings.Join(rowParts, "")
		}

		resultLines = append(resultLines, strings.Join(artLines, "\n"))
	}

	return strings.Join(resultLines, "\n")
}
