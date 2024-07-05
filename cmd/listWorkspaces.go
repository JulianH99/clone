/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

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

		columns := []table.Column{
			{Title: "Name", Width: 15},
			{Title: "Path", Width: 30},
			{Title: "Host", Width: 20},
		}

		rows := []table.Row{}
		for _, w := range workspaces {
			rows = append(rows, table.Row{w.Name, w.Path, w.Host})
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

func init() {
	workspacesCmd.AddCommand(listWorkspacesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listWorkspacesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listWorkspacesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
