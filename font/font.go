package main

func GenerateFont() map[rune][]string {
	font := make(map[rune][]string)

	for c := rune(32); c <= rune(126); c++ {
		lines := make([]string, 8)

		if c == ' ' {
			for i := 0; i < 8; i++ {
				lines[i] = "        "
			}
		} else {
			byteValue := byte(c)
			for i := 0; i < 8; i++ {
				lineBytes := make([]byte, 8)
				for j := 0; j < 8; j++ {
					if ((byteValue ^ byte(i)) & (1 << uint(j))) != 0 {
						lineBytes[j] = '#'
					} else {
						lineBytes[j] = '.'
					}
				}
				lines[i] = string(lineBytes)
			}
		}
		font[c] = lines
	}
	return font
}
