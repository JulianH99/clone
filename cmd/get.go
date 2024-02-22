package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Subfolder string

var getCmd = &cobra.Command{
	Use: "get [Workspace] [Url]",

	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("get", args)

		if args[0] == "" {
			panic("The workspace cannot be null")
		}

		if args[1] == "" {
			panic("Path cannot be null")
		}

		workspc, url := args[0], args[1]

		fmt.Println("Clonning", url, "into workspace", workspc, Subfolder)
	},
}

func init() {
	getCmd.Flags().StringVarP(&Subfolder, "Subfolder", "s", "", "Specify a subfolder to clone the project into")
}
