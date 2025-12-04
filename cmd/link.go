package cmd

import (
	"errors"
	"slices"

	"github.com/JulianH99/clone/internal/config"
	"github.com/JulianH99/clone/internal/workspaces"
	"github.com/spf13/cobra"
)

var linkCmd = &cobra.Command{
	Use:   "link",
	Short: "Link a workspace to a domain. Use clone link [hostName] [workspaceName]",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		hostName := args[0]
		workspaceName := args[1]

		if hostName == "" || workspaceName == "" {
			return errors.New("arguments cannot be empty")
		}

		// TODO: add a wrapper/facade around worskpaces names or checks, it's no
		// good to call config directly
		workspacesList := config.GetConfig().Workspaces
		wNames := workspaces.WorkspacesToNames(workspacesList)

		if !slices.Contains(wNames, workspaceName) {
			return errors.New("no workspace with the provided name was found")
		}

		return nil
	},
}
