package cmd

import (
	"errors"
	"fmt"

	"github.com/JulianH99/clone/core"
	"github.com/JulianH99/clone/validations"
	"github.com/spf13/cobra"
)

var originCmd = &cobra.Command{
	Use:   "origin [Workspace] [Url]",
	Short: "Set the default git origin",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 2 {
			return errors.New("No params were provided, please see command usage")
		}

		if args[0] == "" {
			return errors.New("The workspace cannot be null")
		}

		if args[1] == "" {
			return errors.New("Path cannot be null")
		}

		workspaceName, url := args[0], args[1]

		workspaceInConfig, err := validations.CheckWorkspaceInConfig(workspaceName)

		if err != nil {
			return err
		}

		if workspaceInConfig == nil {
			return errors.New("The requested workspace does not exists, please add it to the config with `clone workspace` command")
		}

		if !validations.CheckValidUrl(url) {
			return errors.New("Repository url should only contain username/repository")
		}

		fmt.Println("Setting default git origin")

		if err := core.SetOrigin(*workspaceInConfig, url); err != nil {
			return err
		}

		return nil
	},
}
