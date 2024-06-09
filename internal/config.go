package internal

import (
	"github.com/spf13/viper"
)

type config struct {
	Workspaces []workspace `yaml:"workspaces"`
}

func GetConfig() config {
	var c config
	viper.Unmarshal(&c)
	return c
}

func AddNewWorkspace(newW workspace) error {
	workspaces := GetConfig().Workspaces
	workspaces = append(workspaces, newW)
	viper.Set("workspaces", workspaces)
	return viper.WriteConfig()
}
