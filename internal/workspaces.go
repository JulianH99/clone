package internal

import "fmt"

type workspace struct {
	Name string `yaml:"name"`
	Path string `yaml:"path"`
	Host string `yaml:"host"`
}

func NewWorkspace(name, path, host string) workspace {
	return workspace{name, path, host}
}

func (w workspace) SaveWorkspace() error {
	ws := GetConfig().Workspaces
	fmt.Println("available workspaces", ws)
	ws = append(ws, w)

	err := writeNewWorkspaces(ws)
	if err != nil {
		return err
	}
	return nil
}

func WorkspacesToNames(workspaces []workspace) []string {
	names := make([]string, len(workspaces))

	for i, w := range workspaces {
		names[i] = w.Name
	}

	return names
}
