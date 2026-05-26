package main

import "strings"

func splitLines(input string) []string {
	split := strings.Split(input, "\\n")
	return split
}


package main

import (
	"strconv"
	"strings"
)

type Animation struct {
	text   string
	frames int
	data   []string
}

func NewAnimation(text string, frames int) *Animation {
	return &Animation{text: text, frames: frames}
}

func blankFrame(content string) []string {
	lines := make([]string, 10)
	for i := 0; i < 10; i++ {
		lines[i] = strings.Repeat(" ", len(content))
	}
	return lines
}

func build(lines []string) string {
	return strings.Join(lines, "\n")
}

func (a *Animation) GenerateSpinFrames() {
	a.data = nil

	spin := []string{"|", "/", "-", "\\"}

	for i := 0; i < a.frames; i++ {
		tag := spin[i%4] + " " + a.text + " " + spin[i%4]
		lines := blankFrame(tag)
		lines[4] = tag
		a.data = append(a.data, build(lines))
	}
}

func (a *Animation) GenerateWaveFrames() {
	a.data = nil

	for i := 0; i < a.frames; i++ {
		offset := i % 5
		tag := strings.Repeat(" ", offset) + a.text

		lines := blankFrame(tag)
		lines[4] = tag

		lines[(4+offset)%10] = tag

		a.data = append(a.data, build(lines))
	}
}

func (a *Animation) GenerateZoomFrames() {
	a.data = nil

	for i := 0; i < a.frames; i++ {
		repeat := (i % 3) + 1
		tag := strings.Repeat(a.text, repeat)

		lines := blankFrame(tag)
		lines[4] = tag

		a.data = append(a.data, build(lines))
	}
}

func (a *Animation) GetFrame(index int) string {
	if len(a.data) == 0 {
		return ""
	}
	return a.data[index%len(a.data)]
}

func (a *Animation) Play() string {
	var b strings.Builder

	for i, f := range a.data {
		b.WriteString("Frame ")
		b.WriteString(strconv.Itoa(i))
		b.WriteString("\n")
		b.WriteString(f)
		b.WriteString("\n---\n")
	}

	return b.String()
}