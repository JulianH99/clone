/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/JulianH99/clone/cmd/hosts"
	"github.com/JulianH99/clone/cmd/workspaces"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "clone",
	Short: "Clone github projects to a saved workspace using a registered custom domain from your ssh config file",
	Long:  `Use clone [domainName] [gitUser]/[repoName] to clone to the current path`,
}

func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	viper.SetConfigName("clone")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("$HOME/.config/")
	emptyArray := make([]any, 0)
	viper.SetDefault("workspaces", emptyArray)
	viper.SetDefault("links", emptyArray)

	err := viper.ReadInConfig()
	if err != nil {
		_ = viper.SafeWriteConfig()
	}

	RootCmd.AddCommand(hosts.HostsCmd)
	RootCmd.AddCommand(workspaces.WorkspacesCmd)
}
