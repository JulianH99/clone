package workspace

import (
	"fmt"

	"github.com/JulianH99/clone/core"
	"github.com/JulianH99/clone/core/ui"
	"github.com/spf13/cobra"
)

var listWorkspacesCmd = &cobra.Command{
	Use:   "list",
	Short: "List all workspaces available",
	RunE: func(cmd *cobra.Command, args []string) error {

		workspaces := core.GetWorkspaces()

		table := ui.DisplayTable(workspaces)
		fmt.Println(table)

		return nil
	},
}
