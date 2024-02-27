package validations

import (
	"errors"
	"regexp"
	"strings"

	"github.com/JulianH99/clone/core"
	"github.com/spf13/viper"
)

var isLetter = regexp.MustCompile("[a-z]+").MatchString

func CheckWorkspaceInConfig(workspaceName string) (*core.Workspace, error) {

	var config core.Workspaces

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
