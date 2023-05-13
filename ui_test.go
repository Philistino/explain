package main

import (
	"fmt"
	"testing"

	"github.com/charmbracelet/lipgloss"
)

// testing that colors are returned in in the correct
// repeating pattern
func TestColorRepeater(t *testing.T) {
	t.Parallel()
	colors := themeRepeater()

	gotColor := make(map[lipgloss.TerminalColor]int)

	// 12 colors in the map
	for i := 0; i < 22; i++ {
		color := colors(fmt.Sprint(i))
		_, ok := gotColor[color]
		if i < 11 {
			if ok {
				t.Error("Color should not be in map: ", i)
			}
			gotColor[color] = i
		} else {
			if !ok {
				t.Error("Color should already be in map: ", i)
			}
			if gotColor[color] != i-11 {
				t.Error("Color should match previous round: ", i)
			}
		}
	}
}

func TestColorCounterSameVal(t *testing.T) {
	t.Parallel()
	colors := themeRepeater()

	gotColor := make(map[string]lipgloss.TerminalColor)

	for i := 0; i < 5; i++ {
		helpRef := fmt.Sprint(i)
		color := colors(helpRef)
		gotColor[helpRef] = color
	}
	for i := 0; i < 5; i++ {
		helpRef := fmt.Sprint(i)
		color := colors(helpRef)
		c := gotColor[helpRef]
		if c != color {
			t.Error("Colors should match")
		}
	}
}
