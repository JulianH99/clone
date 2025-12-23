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

func WorkspacesToOptions(ws []workspaces.Workspace) []huh.Option[workspaces.Workspace] {
	options := make([]huh.Option[workspaces.Workspace], len(ws))
	for i, w := range ws {
		options[i] = huh.NewOption(fmt.Sprintf("%s => %s", w.Name, w.Path), w)
	}

	return options
}

func WithDefault[T comparable](options []huh.Option[T]) []huh.Option[T] {
	return append(options, huh.NewOption("None", *new(T)))
}
