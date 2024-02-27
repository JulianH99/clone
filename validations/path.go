package validations

import (
	"errors"
	"os"
)

func ValidatePath(p string) error {

	stat, err := os.Stat(p)

	if err != nil {
		return err
	}

	if !stat.IsDir() {
		return errors.New("Path should point to a directory")
	}

	return nil
}
