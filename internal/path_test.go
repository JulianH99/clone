package internal

import (
	"os"
	"testing"
)

func TestExpandHome(t *testing.T) {
	homedir, _ := os.UserHomeDir()
	testPath := "~/Documents/projects/personal/"
	expected := homedir + "/Documents/projects/personal/"

	newPath := ExpandHome(testPath)

	if newPath != expected {
		t.Errorf("expected %s got %s", expected, newPath)
	}
}

func TestEmptyOrNil(t *testing.T) {
	testPath := "~/Documents"
	testPath2 := "~/Documents/example"

	t1, _ := IsEmptyDir(testPath)
	t2, _ := IsEmptyDir(testPath2)

	if t1 != false {
		t.Errorf("expected %t got %t", false, t1)
	}

	if t2 != true {
		t.Errorf("expected %t got %t", true, t2)
	}
}
