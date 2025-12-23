package dir_test

import (
	"os"
	"path"
	"testing"

	"github.com/JulianH99/clone/internal/dir"
)

func TestExpandHome(t *testing.T) {
	homedir, _ := os.UserHomeDir()
	testPath := "~/Documents/projects/personal/"
	expected := path.Join(homedir, "/Documents/projects/personal/")

	newPath := dir.ExpandHome(testPath)

	if newPath != expected {
		t.Errorf("expected %s got %s", expected, newPath)
	}
}
