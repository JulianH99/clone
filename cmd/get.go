package cmd

import (
	"errors"
	"fmt"
	"path"

	"github.com/JulianH99/clone/core"
	"github.com/spf13/cobra"
)

var Subfolder string

var getCmd = &cobra.Command{
	Use: "get [Workspace] [Url]",

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

		workspaceInConfig, err := core.CheckWorkspaceInConfig(workspaceName)

		if err != nil {
			return err
		}

		if workspaceInConfig == nil {
			return errors.New("The requested workspace does not exists, please add it to the config with `clone workspace` command")
		}

		validPath := core.CheckValidUrl(url)

		if !validPath {
			return errors.New("repository url should only contain username/reponame")
		}

		fmt.Println("Clonning", url, "into workspace", workspaceInConfig.Name)
		fmt.Println("Path", path.Join(workspaceInConfig.Path, Subfolder))

		cloneErr := core.CloneRepository(*workspaceInConfig, url, Subfolder)

		if cloneErr != nil {
			return cloneErr
		}

		return nil
	},
}

func init() {
	getCmd.Flags().StringVarP(&Subfolder, "Subfolder", "s", "", "Specify a subfolder to clone the project into")
}
