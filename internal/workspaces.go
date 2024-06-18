package internal

type workspace struct {
	Name string `yaml:"name"`
	Path string `yaml:"path"`
	Host string `yaml:"host"`
}

func NewWorkspace(name, path, host string) workspace {
	return workspace{name, path, host}
}

func WorkspacesToNames(workspaces []workspace) []string {
	names := make([]string, len(workspaces))

	for i, w := range workspaces {
		names[i] = w.Name
	}

	return names
}
