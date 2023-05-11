package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/lipgloss"
	"golang.org/x/term"
)

var cmdStyle = lipgloss.NewStyle().
	Border(lipgloss.NormalBorder(), false, false, true, false).
	Padding(1, 0, 0, 0).
	Margin(0, 1)

var expStyle = lipgloss.NewStyle().
	Border(lipgloss.NormalBorder(), true).
	Padding(1).
	Margin(0, 1)

// ansi colors start at 0 so this should end at 15 and restart back at 1. Maybe it should actually only go to 14
func ansiRepeater() func(string) int {
	refColor := make(map[string]int)
	counter := 0
	return func(helpRef string) int {
		color, ok := refColor[helpRef]
		if ok {
			return color
		}
		if counter == 14 {
			counter = 1
			refColor[helpRef] = 1
			return 1
		} else {
			counter++
			refColor[helpRef] = counter
			return counter
		}
	}
}

func BuildCmds(colorRepeater func(string) int, cmds []CmdPart) string {
	bits := make([]string, len(cmds))
	for i, cmd := range cmds {
		if cmd.HelpRef == "" {
			bits[i] = cmdStyle.Render(cmd.CmdPart)
			continue
		}
		style := cmdStyle.
			Copy().
			BorderForeground(
				lipgloss.Color(fmt.Sprint(colorRepeater(cmd.HelpRef))),
			)
		bits[i] = style.Render(cmd.CmdPart)
	}
	return lipgloss.JoinHorizontal(lipgloss.Top, bits...)
}

func BuildExps(colorRepeater func(string) int, expls []Explanation) string {
	bits := make([]string, len(expls))
	for i, cmd := range expls {
		style := expStyle.
			Copy().
			BorderForeground(
				lipgloss.Color(fmt.Sprint(colorRepeater(cmd.HelpRef))),
			)
		bits[i] = style.Render(cmd.Help)
	}
	return lipgloss.JoinVertical(lipgloss.Center, bits...)
}

func Output(cmds []CmdPart, expls []Explanation) string {
	colors := ansiRepeater()
	physicalWidth, _, _ := term.GetSize(int(os.Stdout.Fd()))
	if physicalWidth > 0 {
		width := cmdStyle.GetHorizontalFrameSize()
		cmdStyle = cmdStyle.MaxWidth(physicalWidth - width)

		width = cmdStyle.GetHorizontalFrameSize()
		expStyle = expStyle.Width(physicalWidth - width)

	}
	return lipgloss.JoinVertical(
		lipgloss.Center,
		BuildCmds(colors, cmds),
		BuildExps(colors, expls),
	)
}
