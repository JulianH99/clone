package cmd

import (
	"errors"
	"fmt"

	"github.com/JulianH99/clone/internal"
	"github.com/charmbracelet/huh/spinner"
	"github.com/spf13/cobra"
)

var customPath string

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Clones a github repository. Specify [domainName] [user]/[repoName]",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 2 {
			return errors.New("Not enough parameters. Please specify [domainName] [user]/[repoName]")
		}

		domainName, repo := args[0], args[1]

		githubSshUrl := fmt.Sprintf("git@github.com-%s:%s.git", domainName, repo)

		fmt.Println("this are params", domainName, githubSshUrl)

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
}
