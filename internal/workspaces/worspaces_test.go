package workspaces_test

import (
	"slices"
	"testing"

	"github.com/JulianH99/clone/internal/workspaces"
)

func TestWorkspacesToNames(t *testing.T) {
	ws := []workspaces.Workspace{
		workspaces.NewWorkspace("a", "/a"),
		workspaces.NewWorkspace("b", "/b"),
		workspaces.NewWorkspace("c", "/c"),
	}
	expexted := []string{"a", "b", "c"}

	names := workspaces.WorkspacesToNames(ws)

	if !slices.Equal(names, expexted) {
		t.Errorf("expexted %v, got %v", expexted, names)
	}
}
