/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package workspaces

import (
	"fmt"

	"github.com/JulianH99/clone/internal/config"
	"github.com/JulianH99/clone/internal/ui"
	"github.com/charmbracelet/bubbles/table"
	"github.com/spf13/cobra"
)

// listWorkspacesCmd represents the listWorkspaces command
var listWorkspacesCmd = &cobra.Command{
	Use:   "list",
	Short: "List configured workspaces",
	RunE: func(cmd *cobra.Command, args []string) error {
		workspaces := config.GetConfig().Workspaces

		if len(workspaces) == 0 {
			fmt.Println(ui.InContainer("No workspaces. Use clone workspaces add to create a new workspace"))
			return nil
		}

		columns := []table.Column{
			{Title: "Name", Width: 15},
			{Title: "Path", Width: 30},
		}

		rows := []table.Row{}
		for _, w := range workspaces {
			rows = append(rows, table.Row{w.Name, w.Path})
		}

		t := table.New(
			table.WithColumns(columns),
			table.WithRows(rows),
			table.WithFocused(false),
			table.WithHeight(len(rows)),
		)
		t.SetStyles(ui.TableStyles())

		fmt.Print(ui.InContainer(t.View()))
		return nil
	},
}
