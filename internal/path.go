package internal

import (
	"errors"
	"os"
	"strings"
)

// ExpandHome replaces the ~ in the path for the user homedir
// directory path from os.UserHomeDir()
func ExpandHome(p string) string {
	homedir, err := os.UserHomeDir()
	if err != nil {
		return p
	}
	return strings.Replace(p, "~", homedir, 1)
}

// IsEmptyDir checks if the given path is an empty directory
func IsEmptyDir(p string) (bool, error) {
	if strings.Contains(p, "~") {
		p = ExpandHome(p)
	}

	fileInfo, err := os.Stat(p)

	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return true, nil
		}
		return false, err
	}

	if fileInfo.IsDir() {
		dirEntry, _ := os.ReadDir(p)

		if len(dirEntry) == 0 {
			return true, nil
		}
	}

	return false, nil
}

func IsDir(p string) (bool, error) {
	if strings.Contains(p, "~") {
		p = ExpandHome(p)
	}
	fileInfo, err := os.Stat(p)

	if err != nil {
		return false, err
	}

	if !fileInfo.IsDir() {
		return false, nil
	}
	return true, nil
}
