package internal

import (
	"fmt"
	"os"
	"testing"
)

func TestExpandHome(t *testing.T) {

	homedir, _ := os.UserHomeDir()
	testPath := "~/Documents/projects/personal/"
	expected := homedir + "/Documents/projects/personal/"

	newPath := expandHome(testPath)

	if newPath != expected {
		t.Error(fmt.Sprintf("expected %s got %s", expected, newPath))
	}
}
