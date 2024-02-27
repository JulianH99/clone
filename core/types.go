package core

type Workspace struct {
	Name     string `mapstructure:"name" json:"name"`
	Path     string `mapstructure:"path" json:"path"`
	Hostname string `mapstructure:"hostName" json:"hostName"`
}

type Workspaces struct {
	Workspaces []Workspace `mapstructure:"workspaces" json:"workspaces"`
}
