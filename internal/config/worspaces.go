package config

import (
	"github.com/JulianH99/clone/internal/workspaces"
	"github.com/spf13/viper"
)

func AddNewWorkspace(newW workspaces.Workspace) error {
	workspaces := GetConfig().Workspaces
	workspaces = append(workspaces, newW)
	viper.Set("workspaces", workspaces)
	return viper.WriteConfig()
}

func RemoveWorkspace(workspaceIndex int, workspaces []workspaces.Workspace) error {
	newWorkspaces := append(workspaces[:workspaceIndex], workspaces[workspaceIndex+1:]...)
	viper.Set("workspaces", newWorkspaces)
	return viper.WriteConfig()
}

func SetWorkspace(index int, workspace workspaces.Workspace) error {
	currentWorkspaces := GetConfig().Workspaces
	currentWorkspaces[index] = workspace

	viper.Set("workspaces", currentWorkspaces)
	return viper.WriteConfig()
}
