package cmd

import (
	"errors"
	"fmt"
	"path"

	"github.com/JulianH99/clone/internal"
	"github.com/JulianH99/clone/internal/config"
	"github.com/JulianH99/clone/internal/dir"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/huh/spinner"
	"github.com/spf13/cobra"
)

var subDirectory string

type cloneOptions struct {
	path string
	host string
}

// Generates a custom configuration form
// and returns the selected host and path to clone the repository into
func customConfigForm() (*cloneOptions, error) {
	hosts, err := internal.SshHosts()

	if err != nil {
		return nil, err
	}

	var (
		localPath      string
		host           string
		hostsAsOptions = make([]huh.Option[string], len(hosts)+1)
	)

	hostsAsOptions[0] = huh.NewOption("None", "")

	for i, host := range hosts {
		hostsAsOptions[i+1] = huh.NewOption(string(host), string(host))
	}

	customConfigForm := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Path").
				Value(&localPath).
				Validate(func(s string) error {
					isDir, err := dir.IsEmptyDir(s)

					if err != nil {
						return err
					}

					if !isDir {
						return errors.New("Path provided does not point to an empty directory")
					}
					return nil
				}),
			huh.NewSelect[string]().
				Title("Ssh host").
				Value(&host).
				Options(
					hostsAsOptions...,
				),
		),
	)

	err = customConfigForm.Run()

	if err != nil {
		return nil, err
	}

	return &cloneOptions{path: localPath, host: host}, nil
}

func savedConfigForm() (*cloneOptions, error) {
	var (
		workspaces          = config.GetConfig().Workspaces
		workspacesAsOptions = make([]huh.Option[int], len(workspaces))
		workspace           int
	)

	for i, workspace := range workspaces {
		workspacesAsOptions[i] = huh.NewOption(
			fmt.Sprintf("%s => %s", workspace.Name, workspace.Path),
			i,
		)
	}

	savedConfigurtionForm := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[int]().
				Title("Workspace").
				Options(workspacesAsOptions...).
				Value(&workspace),
			huh.NewInput().
				Title("Subdirectory").
				Description("You can specify a subdirectory in which the project will be cloned inside the chosen workspace").
				Value(&subDirectory),
		),
	)

	err := savedConfigurtionForm.Run()
	if err != nil {
		return nil, err
	}

	chosenWorkspace := workspaces[workspace]

	return &cloneOptions{path: chosenWorkspace.Path, host: chosenWorkspace.Host}, nil
}

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Clones a github repository",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.New("You need to provide a valid git ssh url")
		}

		if !internal.CheckValidSshUrl(args[0]) {
			return fmt.Errorf("%s is not a valid ssh url. You must comply with the format: git@[host]:[user]/[repo].git", args[0])
		}

		var (
			url = args[0]

			// initial form value
			configCustom  = "custom"
			configSaved   = "saved"
			configuration string

			// form result
			cloneOptions *cloneOptions
		)

		fmt.Println("Clonning ", url)

		initialForm := huh.NewForm(
			huh.NewGroup(
				huh.NewSelect[string]().
					Title("Configuration choice").
					Options(
						huh.NewOption("New configuration (custom path and host)", configCustom),
						huh.NewOption("Saved workspace configuration", configSaved),
					).
					Value(&configuration),
			),
		)

		err := initialForm.Run()

		if err != nil {
			return err
		}

		if configuration == configCustom {
			cloneOptions, err = customConfigForm()
		} else {
			cloneOptions, err = savedConfigForm()
		}

		if err != nil {
			return err
		}

		if subDirectory != "" {
			cloneOptions.path = path.Join(cloneOptions.path, subDirectory)
		}

		cloneOptions.path = dir.ExpandHome(cloneOptions.path)
		url = internal.ReplaceHost(url, cloneOptions.host)

		err = spinner.New().Title(fmt.Sprintf("Cloning repository to path %s", cloneOptions.path)).Action(func() {
			if err := internal.Clone(url, cloneOptions.path); err != nil {
				fmt.Println("Error running git clone", err)
			}
			fmt.Println("Done")
		}).Run()

		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	getCmd.Flags().StringVarP(&subDirectory, "subdirectory", "s", "", "Subdirectory inside the chosen path or workspace")
}
