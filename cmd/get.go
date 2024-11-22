package cmd

import (
	"errors"
	"fmt"
	"path"
	"slices"
	"strings"

	"github.com/JulianH99/clone/internal"
	"github.com/JulianH99/clone/internal/config"
	"github.com/JulianH99/clone/internal/dir"
	"github.com/JulianH99/clone/internal/ui"
	"github.com/JulianH99/clone/internal/workspaces"
	"github.com/charmbracelet/huh/spinner"
	"github.com/spf13/cobra"
)

var (
	customPath    string
	workspaceName string
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Clones a github repository. Specify [domainName] [user]/[repoName]",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 2 {
			return errors.New("Not enough parameters. Please specify [domainName] [user]/[repoName]")
		}

		domainName, repo := args[0], args[1]
		repoParts := strings.Split(repo, "/")

		githubSshUrl := fmt.Sprintf("git@github.com-%s:%s.git", domainName, repo)

		fmt.Println("this are params", domainName, githubSshUrl)

		if workspaceName != "" {
			workspaceList := config.GetConfig().Workspaces
			workspacesNames := workspaces.WorkspacesToNames(workspaceList)

			if !slices.Contains(workspacesNames, workspaceName) {
				return errors.New("No workspace with the provided name was found")
			}

			// workspace will be used over custom path if both flags are
			// provided
			for _, w := range workspaceList {
				if w.Name == workspaceName {
					customPath = w.Path
				}
			}
		}

		if customPath != "" {
			customPath = path.Join(dir.ExpandHome(customPath), repoParts[1])
			fmt.Printf("%s\n", ui.InContainer(fmt.Sprintf("Cloning into path %s", customPath)))
		}

		err := spinner.New().Title("Executing git clone").Action(func() {
			if err := internal.Clone(githubSshUrl, customPath); err != nil {
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
	RootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	getCmd.Flags().StringVarP(&customPath, "path", "p", "", "Custom path to clone to (it will be passed down to the `git clone` command)")
	getCmd.Flags().StringVarP(&workspaceName, "workspace", "w", "", "Workspace name to be use when cloning. The path associated will be passed on to git clone command")
}
