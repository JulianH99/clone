package workspace

import (
	"errors"
	"fmt"

	"github.com/JulianH99/clone/core"
	"github.com/JulianH99/clone/validations"
	"github.com/spf13/cobra"
)

type WorkspaceRequiredError struct{}

func (w *WorkspaceRequiredError) Error() string {
	return "Workspace name is required"
}

func validWorkspace(workspace string) error {

	return nil
}

func validPath(path string) error {

	return nil

}

var WorkspaceCmd = &cobra.Command{
	Use: "workspace [NAME] [PATH] [DOMAIN_POSTFIX]",
	RunE: func(cmd *cobra.Command, args []string) error {

		if len(args) == 0 {
			return errors.New("No args where passed to the command")
		}

		if len(args) != 3 {
			return errors.New("Not enough parameters where provided")
		}

		// validate workspace
		if err := validations.ValidateWorkspaceFormat(args[0]); err != nil {
			return err
		}

		// validate path
		if err := validations.ValidatePath(args[1]); err != nil {
			return err
		}

		fmt.Println("Saving workspace")

		workspace := core.Workspace{
			Name:     args[0],
			Path:     args[1],
			Hostname: args[2],
		}

		cobra.CheckErr(core.SaveWorkspace(workspace))

		return nil
	},
}
