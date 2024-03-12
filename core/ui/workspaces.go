package ui

import (
	"github.com/JulianH99/clone/core"
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/lipgloss"
)

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color("240"))

func DisplayTable(workspaces []core.Workspace) string {

	largestPathLength := core.FindLongestPathLength(workspaces)
	columns := []table.Column{
		{Title: "Name", Width: 20},
		{Title: "Path", Width: largestPathLength},
		{Title: "Hostname", Width: 20},
	}

	rows := make([]table.Row, len(workspaces))

	for i, workspace := range workspaces {
		rows[i] = table.Row{
			workspace.Name,
			workspace.Path,
			workspace.Hostname,
		}
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithHeight(len(rows)),
	)

	s := table.DefaultStyles()

	s.Header = s.Header.
		BorderStyle(lipgloss.ThickBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(true)

	t.SetStyles(s)
	return baseStyle.Render(t.View())

}
