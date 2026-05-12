package main

import "fmt"

func validate(text string) (rune, error) {
	for i, char := range text {
		if char < 32 || char > 126 {
			return char, fmt.Errorf("invalid char [%c] at position %d  ", char, i)
		}
	}
	return 0, nil
}
