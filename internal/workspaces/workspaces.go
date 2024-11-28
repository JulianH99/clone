package workspaces

type Workspace struct {
	Name string `yaml:"name"`
	Path string `yaml:"path"`
}

func NewWorkspace(name, path string) Workspace {
	return Workspace{name, path}
}

func WorkspacesToNames(workspaces []Workspace) []string {
	names := make([]string, len(workspaces))

	for i, w := range workspaces {
		names[i] = w.Name
	}

	return names
}
