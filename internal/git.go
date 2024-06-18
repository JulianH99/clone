package internal

import (
	"os/exec"
	"strings"
)

// clone a repository to the given path
func Clone(repository string, path string) error {

	command := exec.Command("git", "clone", repository, path)

	_, err := command.Output()

	if err != nil {
		return err
	}

	return nil

}

// replace the host in the git ssh url for the given host in the ${host}
// parameter
func ReplaceHost(repository string, host string) string {
	var (
		atSign    = strings.Index(repository, "@")
		colonSign = strings.Index(repository, ":")
		oldHost   = repository[atSign:colonSign]
	)

	return strings.Replace(repository, oldHost, host, 1)
}
