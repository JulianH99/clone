package cmd

import (
	"fmt"
	"os"

	"github.com/JulianH99/clone/cmd/workspace"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{

	Use:   "clone",
	Short: "Clone a github project",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func Execute() error {

	// if err := rootCmd.Execute(); err != nil {
	//
	// 	cobra.CheckErr(err)
	//
	// 	return err
	//
	// }
	//
	// return nil
	return rootCmd.Execute()
}

func init() {

	fmt.Println("Initializing commands")

	initConfiguration()

	rootCmd.AddCommand(getCmd)
	rootCmd.AddCommand(workspace.WorkspaceCmd)

}

func initConfiguration() {

	home, err := os.UserHomeDir()

	cobra.CheckErr(err)

	viper.AddConfigPath(home)
	viper.SetConfigType("json")
	viper.SetConfigName(".clone")
	viper.SetDefault("workspaces", []any{})
	viper.SafeWriteConfig()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	} else {
		cobra.CheckErr(err)
	}

}
