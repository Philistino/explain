package main

import (
	"github.com/charmbracelet/lipgloss"
)

// tintForestBlue (ForestBlue) is a collection of lipgloss styles.
//
// Modified from: https://github.com/lrstanley/bubbletint/
type tintForestBlue struct{}

func (t *tintForestBlue) BrightBlue() lipgloss.TerminalColor {
	return lipgloss.Color("#39a7a2")
}

func (t *tintForestBlue) BrightCyan() lipgloss.TerminalColor {
	return lipgloss.Color("#6096bf")
}

func (t *tintForestBlue) BrightGreen() lipgloss.TerminalColor {
	return lipgloss.Color("#6bb48d")
}

func (t *tintForestBlue) BrightPurple() lipgloss.TerminalColor {
	return lipgloss.Color("#7e62b3")
}

func (t *tintForestBlue) BrightRed() lipgloss.TerminalColor {
	return lipgloss.Color("#fb3d66")
}

func (t *tintForestBlue) BrightYellow() lipgloss.TerminalColor {
	return lipgloss.Color("#30c85a")
}

func (t *tintForestBlue) Blue() lipgloss.TerminalColor {
	return lipgloss.Color("#8ed0ce")
}

func (t *tintForestBlue) Cyan() lipgloss.TerminalColor {
	return lipgloss.Color("#31658c")
}

func (t *tintForestBlue) Green() lipgloss.TerminalColor {
	return lipgloss.Color("#92d3a2")
}

func (t *tintForestBlue) Purple() lipgloss.TerminalColor {
	return lipgloss.Color("#5e468c")
}

func (t *tintForestBlue) Red() lipgloss.TerminalColor {
	return lipgloss.Color("#f8818e")
}

func (t *tintForestBlue) White() lipgloss.TerminalColor {
	return lipgloss.Color("#e2d8cd")
}

func (t *tintForestBlue) Yellow() lipgloss.TerminalColor {
	return lipgloss.Color("#1a8e63")
}

func (t *tintForestBlue) ByNumber(n int) lipgloss.TerminalColor {
	cMap := map[int]lipgloss.TerminalColor{
		0:  t.Blue(), // was blue
		1:  t.Cyan(),
		2:  t.Green(),
		3:  t.Purple(),
		4:  t.Red(),
		5:  t.BrightBlue(),
		6:  t.BrightCyan(),
		7:  t.BrightGreen(),
		8:  t.BrightPurple(),
		9:  t.BrightRed(),
		10: t.Yellow(),
		11: t.BrightYellow(),
	}
	return cMap[n]
}
