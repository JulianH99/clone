package ui

import (
	"github.com/charmbracelet/lipgloss"
)

func baseContainer() lipgloss.Style {
	return lipgloss.NewStyle().
		Padding(1).
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(PrimaryColor)
}

func InContainer(s string) string {
	return baseContainer().Render(s)
}

func InColoredContainer(s string) string {
	return baseContainer().
		Foreground(AccentColor).
		Render(s)
}
