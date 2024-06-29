package ui

import (
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/lipgloss"
)

var (
	PrimaryColor lipgloss.Color = lipgloss.Color("#2596be")
	AccentColor  lipgloss.Color = lipgloss.Color("#f5f5fd")
)

func TableStyles() table.Styles {
	s := table.DefaultStyles()

	s.Header = s.Header.Foreground(AccentColor)
	s.Selected = s.Cell.Padding(0)
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(PrimaryColor).
		BorderBottom(true)

	return s
}
