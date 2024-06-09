package internal

import (
	"os"
	"strings"
)

// expandHome replaces the ~ in the path for the user homedir
// directory path from os.UserHomeDir()
func expandHome(p string) string {
	homedir, err := os.UserHomeDir()
	if err != nil {
		return p
	}
	return strings.Replace(p, "~", homedir, 1)
}

func IsDir(p string) (bool, error) {
	if strings.Index(p, "~") != -1 {
		p = expandHome(p)
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
