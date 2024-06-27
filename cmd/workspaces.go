/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// workspacesCmd represents the workspaces command
var workspacesCmd = &cobra.Command{
	Use:   "workspaces",
	Short: "Manage workspaces (collections of a path and a host)",
}

func init() {
	rootCmd.AddCommand(workspacesCmd)
}
