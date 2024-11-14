package workspaces

import (
	"errors"
	"fmt"
	"slices"

	"github.com/JulianH99/clone/internal/config"
	"github.com/JulianH99/clone/internal/dir"
	"github.com/JulianH99/clone/internal/ui"
	"github.com/JulianH99/clone/internal/workspaces"
	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a new workspace into the configuration",
	RunE: func(cmd *cobra.Command, args []string) error {
		workspaceList := workspaces.WorkspacesToNames(config.GetConfig().Workspaces)

		var (
			name string
			path string
		)

		hostAsOptions := make([]huh.Option[string], 0)
		hostAsOptions = append(hostAsOptions, huh.NewOption("None", ""))

		form := huh.NewForm(
			huh.NewGroup(
				huh.NewInput().
					Title("Name").
					Value(&name).
					Validate(func(s string) error {
						if s == "" {
							return errors.New("Value cannot be empty")
						}

						if slices.Contains(workspaceList, s) {
							return errors.New("Workspace name already in use")
						}
						return nil
					}),
				huh.NewInput().
					Title("Path").
					Value(&path).
					Validate(func(s string) error {
						isDir, err := dir.IsDir(s)

						if err != nil {
							return err
						}

						if !isDir {
							return errors.New("Path provided is not a directory")
						}
						return nil
					}),
			),
		)

		err := form.Run()

		if err != nil {
			return err
		}

		w := workspaces.NewWorkspace(name, path)
		err = config.AddNewWorkspace(w)

		if err != nil {
			return err
		}

		fmt.Print(ui.InColoredContainer(fmt.Sprintf("Workspace %s created", w.Name)))

		return nil
	},
}
