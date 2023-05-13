package main

import (
	"github.com/charmbracelet/lipgloss"
)

// TintForestBlue (ForestBlue) is a collection of lipgloss styles.
//
// Reference: https://github.com/lrstanley/bubbletint/blob/master/DEFAULT_TINTS.md#ForestBlue
type TintForestBlue struct{}

func (t *TintForestBlue) BrightBlue() lipgloss.TerminalColor {
	return lipgloss.Color("#39a7a2")
}

func (t *TintForestBlue) BrightCyan() lipgloss.TerminalColor {
	return lipgloss.Color("#6096bf")
}

func (t *TintForestBlue) BrightGreen() lipgloss.TerminalColor {
	return lipgloss.Color("#6bb48d")
}

func (t *TintForestBlue) BrightPurple() lipgloss.TerminalColor {
	return lipgloss.Color("#7e62b3")
}

func (t *TintForestBlue) BrightRed() lipgloss.TerminalColor {
	return lipgloss.Color("#fb3d66")
}

func (t *TintForestBlue) BrightYellow() lipgloss.TerminalColor {
	return lipgloss.Color("#30c85a")
}

func (t *TintForestBlue) Blue() lipgloss.TerminalColor {
	return lipgloss.Color("#8ed0ce")
}

func (t *TintForestBlue) Cyan() lipgloss.TerminalColor {
	return lipgloss.Color("#31658c")
}

func (t *TintForestBlue) Green() lipgloss.TerminalColor {
	return lipgloss.Color("#92d3a2")
}

func (t *TintForestBlue) Purple() lipgloss.TerminalColor {
	return lipgloss.Color("#5e468c")
}

func (t *TintForestBlue) Red() lipgloss.TerminalColor {
	return lipgloss.Color("#f8818e")
}

func (t *TintForestBlue) White() lipgloss.TerminalColor {
	return lipgloss.Color("#e2d8cd")
}

func (t *TintForestBlue) Yellow() lipgloss.TerminalColor {
	return lipgloss.Color("#1a8e63")
}

func (t *TintForestBlue) ByNumber(n int) lipgloss.TerminalColor {
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
