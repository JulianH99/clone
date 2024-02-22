package core

import (
	"errors"
	"regexp"
	"slices"
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

func CheckWorkspaceInConfig(workspace string) (bool, error) {

	if err := viper.ReadInConfig(); err != nil {
		return false, err
	}

	workspaces := viper.GetStringSlice("workspaces")

	if slices.Contains(workspaces, workspace) {
		return true, nil
	}

	return false, nil
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
