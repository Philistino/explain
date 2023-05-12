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
	Padding(0, 1).
	Margin(0, 1)

var urlStyle = lipgloss.NewStyle().
	Padding(1, 0, 1, 0).
	Foreground(lipgloss.Color("12"))

// ansi colors start at 0 so this should end at 15 and restart back at 1. Maybe it should actually only go to 14
func ansiRepeater() func(string) int {
	refColor := make(map[string]int, 20)

	// setting ref color to white instances when
	// helpRef == "", meaning no help box will be given
	refColor[""] = 15
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

func buildCmds(colorRepeater func(string) int, cmds []CmdPart) (string, bool) {
	bits := make([]string, len(cmds))
	// track if an additional box should be added to the explanantions
	// for cases where an explanation was not found
	addNoHelp := false
	for i, cmd := range cmds {
		if cmd.HelpRef == "" {
			addNoHelp = true
		}
		style := cmdStyle.
			Copy().
			BorderForeground(
				lipgloss.Color(fmt.Sprint(colorRepeater(cmd.HelpRef))),
			)
		bits[i] = style.Render(cmd.CmdPart)
	}
	return lipgloss.JoinHorizontal(lipgloss.Top, bits...), addNoHelp
}

func buildExps(colorRepeater func(string) int, expls []Explanation, addNoHelp bool) string {
	bits := make([]string, len(expls))
	for i, cmd := range expls {
		style := expStyle.
			Copy().
			BorderForeground(
				lipgloss.Color(fmt.Sprint(colorRepeater(cmd.HelpRef))),
			)
		bits[i] = style.Render(cmd.Help)
	}
	if addNoHelp {
		bits = append(
			bits,
			expStyle.
				Copy().
				BorderForeground(lipgloss.Color(fmt.Sprint(15))). // 15 is for the color white
				Render("Unknown"),
		)
	}
	return lipgloss.JoinVertical(lipgloss.Center, bits...)
}

func Output(cmds []CmdPart, expls []Explanation, url string) string {
	colors := ansiRepeater()
	physicalWidth, _, _ := term.GetSize(int(os.Stdout.Fd()))
	if physicalWidth > 0 {
		width := cmdStyle.GetHorizontalFrameSize()
		cmdStyle = cmdStyle.MaxWidth(physicalWidth - width)

		width = cmdStyle.GetHorizontalFrameSize()
		expStyle = expStyle.Width(physicalWidth - width)

	}
	cmdString, addNoHelp := buildCmds(colors, cmds)
	expStrings := buildExps(colors, expls, addNoHelp)

	return lipgloss.JoinVertical(
		lipgloss.Center,
		cmdString,
		expStrings,
		urlStyle.Render(url),
	)
}
