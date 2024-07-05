package workspaces

type Workspace struct {
	Name string `yaml:"name"`
	Path string `yaml:"path"`
	Host string `yaml:"host"`
}

func NewWorkspace(name, path, host string) Workspace {
	return Workspace{name, path, host}
}

func WorkspacesToNames(workspaces []Workspace) []string {
	names := make([]string, len(workspaces))

	for i, w := range workspaces {
		names[i] = w.Name
	}

	return names
}
