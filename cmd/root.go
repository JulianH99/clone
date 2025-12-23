/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/JulianH99/clone/cmd/hosts"
	workspacesCmd "github.com/JulianH99/clone/cmd/workspaces"
	"github.com/JulianH99/clone/internal"
	"github.com/JulianH99/clone/internal/config"
	"github.com/JulianH99/clone/internal/dir"
	"github.com/JulianH99/clone/internal/ui"
	"github.com/JulianH99/clone/internal/workspaces"
	"github.com/adrg/xdg"
	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	host          string
	workspaceName string
	customPath    string
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "clone",
	Short: "Clone github projects to a saved workspace using a registered custom host from your ssh config file",
	Long:  `Use clone [gitUser]/[repoName] to clone to the current path`,
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		var (
			repo          = args[0]
			cfg           = config.GetConfig()
			rawHosts, err = internal.SshHosts()
			fields        []huh.Field
			workspace     workspaces.Workspace
		)

		if err != nil {
			return err
		}
		repoValidation, err := regexp.Compile("^[a-zA-Z0-9-_.]+/[a-zA-Z0-9-_.]+$")
		if err != nil {
			return err
		}

		if !repoValidation.MatchString(repo) {
			return fmt.Errorf("invalid repository name, please use the format user/repoName")
		}

		hosts := ui.HostsToOptions(rawHosts)
		workspaceOptions := ui.WithDefault(ui.WorkspacesToOptions(cfg.Workspaces))

		if host == "" {
			fields = append(fields,
				huh.NewSelect[string]().
					Title("Host").
					Options(hosts...).
					Value(&host),
			)
		} else {
			for _, h := range rawHosts {
				if h == host || strings.Contains(h, host) {
					host = h
				}
			}
		}

		if workspaceName == "" {
			fields = append(fields,
				huh.NewSelect[workspaces.Workspace]().
					Title("Workspace").
					Options(workspaceOptions...).
					Value(&workspace),
			)
		} else {
			for _, w := range cfg.Workspaces {
				if w.Name == workspaceName {
					workspace = w
				}
			}
			if workspace.Name == "" {
				return fmt.Errorf("invalid workspace name. Please see available workspaces with clone workspaces list")
			}
		}

		if len(fields) != 0 {
			form := huh.NewForm(
				huh.NewGroup(
					fields...,
				),
			)

			err = form.Run()
			if err != nil {
				return err
			}
		}

		sshUrl := fmt.Sprintf("git@%s:%s.git", host, repo)
		p := ""
		repoName := strings.Split(repo, "/")[1]
		if workspace.Path != "" {
			p = filepath.Join(
				dir.ExpandHome(workspace.Path),
				repoName,
			)
		}

		if customPath != "" {
			p = filepath.Join(
				dir.ExpandHome(customPath),
				repoName,
			)
		}
		p = filepath.Clean(p)

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
		defer cancel()

		err = internal.Clone(ctx, sshUrl, p)
		if err != nil {
			fmt.Println(ui.InContainer(
				fmt.Sprintf("Error cloning %s to %s: %s", repo, p, err),
			))
			return err
		}

		return nil
	},
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
	viper.AddConfigPath(xdg.ConfigHome)
	emptyArray := make([]any, 0)
	viper.SetDefault("workspaces", emptyArray)
	viper.SetDefault("links", emptyArray)

	err := viper.ReadInConfig()
	if err != nil {
		_ = viper.SafeWriteConfig()
	}

	RootCmd.Flags().StringVarP(&host, "host", "t", "", "Host to be used when cloning, you can specify a shorthand, like work (for github.com-work) or the full host")
	RootCmd.Flags().StringVarP(&workspaceName, "workspace", "w", "", "Workspace to be used when cloning")
	RootCmd.Flags().StringVarP(&customPath, "path", "p", "", "Custom path to be pased to git command. Will be used over workspace if both plags are provided")
	RootCmd.AddCommand(hosts.ListsHostsCmd)
	RootCmd.AddCommand(workspacesCmd.WorkspacesCmd)
}
