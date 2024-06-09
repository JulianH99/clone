package ui

import (
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/lipgloss"
)

var (
	PrimaryColor lipgloss.Color = lipgloss.Color("#7571F9")
	GreenColor   lipgloss.Color = lipgloss.Color("#02BF87")
)

func TableStyles() table.Styles {
	s := table.DefaultStyles()

	s.Header = s.Header.Foreground(GreenColor)
	s.Selected = s.Cell.Padding(0)
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(PrimaryColor).
		BorderBottom(true)

	return s
}
