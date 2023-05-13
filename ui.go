package main

import (
	"os"

	"github.com/charmbracelet/lipgloss"
	"golang.org/x/term"
)

// themes: CrystalViolet, CyberCube, DimmedMonokai,
// Prefer: DoomPeacock, Mirage
// DraculaPlus, Espresso, ForestBlue, Mirage

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
	Foreground(lipgloss.Color("12")).
	Margin(0, 1)

func themeRepeater() func(string) lipgloss.TerminalColor {
	theme := TintForestBlue{}
	refColor := make(map[string]lipgloss.TerminalColor, 20)

	// setting ref color to white instances when
	// helpRef == "", meaning no help box will be given
	refColor[""] = theme.White()
	counter := -1
	return func(helpRef string) lipgloss.TerminalColor {
		color, ok := refColor[helpRef]
		if ok {
			return color
		}
		counter++
		colorNumb := counter % 11 // 12 colors in the map, 11 is the largest integer mapped
		color = theme.ByNumber(colorNumb)
		refColor[helpRef] = color
		return color
	}
}

func buildCmds(colorRepeater func(string) lipgloss.TerminalColor, cmds []CmdPart) (string, bool) {
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
				colorRepeater(cmd.HelpRef),
			)
		bits[i] = style.Render(cmd.CmdPart)
	}
	return lipgloss.JoinHorizontal(lipgloss.Top, bits...), addNoHelp
}

func buildExps(colorRepeater func(string) lipgloss.TerminalColor, expls []Explanation, addNoHelp bool) string {
	bits := make([]string, len(expls))
	for i, cmd := range expls {
		style := expStyle.
			Copy().
			BorderForeground(
				colorRepeater(cmd.HelpRef),
			)
		bits[i] = style.Render(cmd.Help)
	}
	if addNoHelp {
		bits = append(
			bits,
			expStyle.
				Copy().
				BorderForeground(lipgloss.Color("15")). // 15 is for the color white
				Render("Unknown"),
		)
	}
	return lipgloss.JoinVertical(lipgloss.Center, bits...)
}

func Output(cmds []CmdPart, expls []Explanation, url string) string {
	colors := themeRepeater()
	physicalWidth, _, _ := term.GetSize(int(os.Stdout.Fd()))
	if physicalWidth > 0 {
		width := cmdStyle.GetHorizontalFrameSize()
		cmdStyle = cmdStyle.MaxWidth(physicalWidth - width)

		width = cmdStyle.GetHorizontalFrameSize()
		expStyle = expStyle.Width(physicalWidth - width)

		width = urlStyle.GetHorizontalFrameSize()
		urlStyle = urlStyle.Width(physicalWidth - width)
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
