package cmd

import (
	"fmt"

	"github.com/JulianH99/clone/internal/config"
	"github.com/JulianH99/clone/internal/ui"
	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
)

var deleteWorkspacesCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a workspace",
	RunE: func(cmd *cobra.Command, args []string) error {
		var (
			workspaces = config.GetConfig().Workspaces
			options    = make([]huh.Option[int], len(workspaces))
			workspace  int
		)

		for i, w := range workspaces {
			options[i] = huh.NewOption(fmt.Sprintf("%s => %s", w.Name, w.Path), i)
		}

		form := huh.NewForm(
			huh.NewGroup(
				huh.NewSelect[int]().
					Title("Workspace to delete").
					Options(options...).
					Value(&workspace),
			),
		)

		err := form.Run()

		if err != nil {
			return fmt.Errorf("Error running form %w", err)
		}

		config.RemoveWorkspace(workspace, workspaces)

		fmt.Print(ui.InColoredContainer(fmt.Sprintf("Workspace %s deleted", workspaces[workspace].Name)))

		return nil
	},
}

func init() {
	workspacesCmd.AddCommand(deleteWorkspacesCmd)
}
