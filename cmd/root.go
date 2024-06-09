/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "clone",
	Short: "Clone github projects to a saved workspace using a registered custom domain from your ssh config file",
	Long:  ``,
	// RunE: func(cmd *cobra.Command, args []string) error {
	// 	return nil
	// },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.clone.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	viper.SetConfigName("clone")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("$HOME/.config/")
	viper.SafeWriteConfig()
	emptyArray := make([]any, 0)
	viper.SetDefault("workspaces", emptyArray)

	err := viper.ReadInConfig()

	if err != nil {
		fmt.Println("error reading config file. please check your .config path", err)
	}

}
