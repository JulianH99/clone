package internal

import "os/exec"

func CheckGit() error {
	_, err := exec.LookPath("git")
	return err
}
