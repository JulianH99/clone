package core

type Workspace struct {
	Name     string `mapstructure:"name" json:"name"`
	Path     string `mapstructure:"path" json:"path"`
	Hostname string `mapstructure:"hostName" json:"hostName"`
}

type Workspaces struct {
	Workspaces []Workspace `mapstructure:"workspaces" json:"workspaces"`
}

func FindLongestPathLength(workspaces []Workspace) int {
	longest := 0
	for _, workspace := range workspaces {
		if len(workspace.Path) > longest {
			longest = len(workspace.Path)
		}
	}
	return longest
}
