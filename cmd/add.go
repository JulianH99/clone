/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"slices"

	"github.com/JulianH99/clone/internal"
	"github.com/JulianH99/clone/internal/ui"
	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a new workspace into the configuration",
	RunE: func(cmd *cobra.Command, args []string) error {
		hosts, err := internal.SshHosts()
		workspaces := internal.WorkspacesToNames(internal.GetConfig().Workspaces)

		if err != nil {
			return err
		}

		var (
			name string
			path string
			host string
		)

		hostAsOptions := make([]huh.Option[string], 0)
		hostAsOptions = append(hostAsOptions, huh.NewOption("None", ""))

		for _, host := range hosts {
			hostAsOptions = append(hostAsOptions, huh.NewOption(string(host), string(host)))
		}

		form := huh.NewForm(
			huh.NewGroup(
				huh.NewInput().
					Title("Name").
					Value(&name).
					Validate(func(s string) error {
						if s == "" {
							return errors.New("Value cannot be empty")
						}

						if slices.Contains(workspaces, s) {
							return errors.New("Workspace name already in use")
						}
						return nil
					}),
				huh.NewInput().
					Title("Path").
					Value(&path).
					Validate(func(s string) error {
						isDir, err := internal.IsDir(s)

						if err != nil {
							return err
						}

						if !isDir {
							return errors.New("Path provided is not a directory")
						}
						return nil
					}),
				huh.NewSelect[string]().
					Title("What host would ou like to use?").
					Options(hostAsOptions...).
					Value(&host),
			),
		)

		err = form.Run()

		if err != nil {
			return err
		}

		w := internal.NewWorkspace(name, path, host)
		err = internal.AddNewWorkspace(w)

		if err != nil {
			return err
		}

		fmt.Print(ui.InColoredContainer(fmt.Sprintf("Workspace %s created", w.Name)))

		return nil
	},
}

func init() {
	workspacesCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
