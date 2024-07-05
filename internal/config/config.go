package config

import (
	"github.com/JulianH99/clone/internal/workspaces"
	"github.com/spf13/viper"
)

type config struct {
	Workspaces []workspaces.Workspace `yaml:"workspaces"`
}

func GetConfig() config {
	var c config
	viper.Unmarshal(&c)
	return c
}
