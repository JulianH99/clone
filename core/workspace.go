package core

import (
	"errors"
	"regexp"
	"strings"

	"github.com/spf13/viper"
)

type Workspace struct {
	Name     string `mapstructure:"name" json:"name"`
	Path     string `mapstructure:"path" json:"path"`
	Hostname string `mapstructure:"hostName" json:"hostName"`
}

type Workspaces struct {
	Workspaces []Workspace `mapstructure:"workspaces" json:"workspaces"`
}

var isLetter = regexp.MustCompile("[a-z]+").MatchString

func CheckWorkspaceInConfig(workspaceName string) (*Workspace, error) {

	var config Workspaces

	viper.Unmarshal(&config)

	workspaces := config.Workspaces

	if len(workspaces) == 0 {
		return nil, errors.New("No workspaces registered")
	}

	for _, workspace := range workspaces {
		if workspace.Name == workspaceName {
			return &workspace, nil
		}
	}

	return nil, nil
}

func ValidateWorkspaceFormat(workspace string) error {

	if workspace == "" {
		return errors.New("Workspace cannot be empty")
	}

	if !isLetter(strings.TrimSpace(workspace)) {
		return errors.New("Workspace name can only contain alphabetical characters")
	}

	return nil
}
