/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package workspaces

import (
	"github.com/spf13/cobra"
)

// workspacesCmd represents the workspaces command
var WorkspacesCmd = &cobra.Command{
	Use:   "workspaces",
	Short: "Manage workspaces (collections of a path and a host)",
}

func init() {
	WorkspacesCmd.AddCommand(listWorkspacesCmd)
	WorkspacesCmd.AddCommand(addCmd)
	WorkspacesCmd.AddCommand(deleteWorkspacesCmd)
}
