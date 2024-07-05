package dir_test

import (
	"os"
	"testing"

	"github.com/JulianH99/clone/internal/dir"
)

func TestExpandHome(t *testing.T) {
	homedir, _ := os.UserHomeDir()
	testPath := "~/Documents/projects/personal/"
	expected := homedir + "/Documents/projects/personal/"

	newPath := dir.ExpandHome(testPath)

	if newPath != expected {
		t.Errorf("expected %s got %s", expected, newPath)
	}
}

func TestEmptyOrNil(t *testing.T) {
	testPath := "~/Documents"
	testPath2 := "~/Documents/example"

	t1, _ := dir.IsEmptyDir(testPath)
	t2, _ := dir.IsEmptyDir(testPath2)

	if t1 != false {
		t.Errorf("expected %t got %t", false, t1)
	}

	if t2 != true {
		t.Errorf("expected %t got %t", true, t2)
	}
}
