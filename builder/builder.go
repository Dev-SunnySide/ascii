package main

import (
	"strings"
)

type textItem struct {
	text  string
	style string
}

type ArtBuilder struct {
	items        []textItem
	currentStyle string
}

func NewArtBuilder() *ArtBuilder {
	return &ArtBuilder{
		items:        []textItem{},
		currentStyle: "normal",
	}
}

func (ab *ArtBuilder) AddText(text string) *ArtBuilder {
	ab.items = append(ab.items, textItem{
		text:  text,
		style: ab.currentStyle,
	})
	return ab
}

func (ab *ArtBuilder) SetStyle(style string) *ArtBuilder {
	switch style {
	case "normal", "bold", "italic", "outline":
		ab.currentStyle = style
		if len(ab.items) > 0 {
			ab.items[len(ab.items)-1].style = style
		}
	default:
		panic("invalid style")
	}
	return ab
}

func (ab *ArtBuilder) Build() string {
	if len(ab.items) == 0 {
		return ""
	}

	var rows [8]string

	for _, item := range ab.items {
		for _, char := range item.text {
			var charRows [8]string
			baseChar := byte(char)

			for i := 0; i < 8; i++ {
				if ((baseChar ^ byte(i)) & 1) == 0 {
					charRows[i] = "##"
				} else {
					charRows[i] = ".."
				}
			}

			switch item.style {
			case "bold":
				for i := 0; i < 8; i++ {
					charRows[i] = charRows[i] + charRows[i]
				}
			case "italic":
				for i := 0; i < 8; i++ {
					spaces := strings.Repeat(" ", i)
					charRows[i] = spaces + charRows[i]
				}
			case "outline":
				for i := 0; i < 8; i++ {
					charRows[i] = "|" + charRows[i] + "|"
				}
			case "normal":
			}

			for i := 0; i < 8; i++ {
				rows[i] += charRows[i]
			}
		}
	}

	var finalBuilder strings.Builder
	for i := 0; i < 8; i++ {
		finalBuilder.WriteString(rows[i])
		finalBuilder.WriteString("\n")
	}

	return finalBuilder.String()
}
