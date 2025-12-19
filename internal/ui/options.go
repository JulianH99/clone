package ui

import (
	"fmt"

	"github.com/JulianH99/clone/internal/workspaces"
	"github.com/charmbracelet/huh"
)

func HostsToOptions(hosts []string) []huh.Option[string] {
	options := make([]huh.Option[string], len(hosts))
	for i, host := range hosts {
		options[i] = huh.NewOption(host, host)
	}
	return options
}

func WorkspacesToOptions(ws []workspaces.Workspace) []huh.Option[string] {
	options := make([]huh.Option[string], len(ws))
	for i, w := range ws {
		options[i] = huh.NewOption(fmt.Sprintf("%s => %s", w.Name, w.Path), w.Name)
	}

	return options
}
