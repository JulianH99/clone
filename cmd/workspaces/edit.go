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

var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Allows for editing of a workspace, changing its name or path",
	RunE: func(cmd *cobra.Command, args []string) error {
		var (
			workspacesList  = config.GetConfig().Workspaces
			workspacesNames = workspaces.WorkspacesToNames(workspacesList)
			options         = make([]huh.Option[int], len(workspacesList))
			workspace       int
		)

		if len(workspacesList) == 0 {
			fmt.Println(ui.InContainer("No workspaces created yet"))
		}

		for i, w := range workspacesList {
			options[i] = huh.NewOption(fmt.Sprintf("%s => %s", w.Name, w.Path), i)
		}

		form := huh.NewForm(
			huh.NewGroup(
				huh.NewSelect[int]().
					Title("Workspace to edit").
					Options(options...).
					Value(&workspace),
			),
		)

		err := form.Run()
		if err != nil {
			return fmt.Errorf("Error running form %w", err)
		}
		newName := workspacesList[workspace].Name
		newPath := workspacesList[workspace].Path

		form = huh.NewForm(
			huh.NewGroup(
				huh.NewInput().
					Title("Name").
					Value(&newName).
					Validate(func(s string) error {
						if s == "" {
							return errors.New("Value cannot be empty")
						}

						if slices.Contains(workspacesNames, s) && s != workspacesList[workspace].Name {
							return errors.New("Workspace name already in use")
						}
						return nil
					}),
				huh.NewInput().
					Title("Path").
					Value(&newPath).
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

		err = form.Run()
		if err != nil {
			return err
		}

		err = config.SetWorkspace(workspace, workspaces.NewWorkspace(newName, newPath))
		if err != nil {
			return err
		}

		fmt.Print(ui.InContainer("Workspace edited!"))

		return nil
	},
}
