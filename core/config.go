package core

import (
	"errors"

	"github.com/spf13/viper"
)

func SaveWorkspace(workspace Workspace) error {

	var config Workspaces

	viper.Unmarshal(&config)

	workspaces := config.Workspaces

	for _, w := range workspaces {
		if w.Name == workspace.Name {
			return errors.New("This worspace already exists")
		}
	}

	workspaces = append(workspaces, workspace)

	viper.Set("workspaces", workspaces)
	viper.WriteConfig()

	return nil
}

func GetWorkspaces() []Workspace {
	var config Workspaces
	viper.Unmarshal(&config)
	return config.Workspaces
}
