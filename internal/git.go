package internal

import (
	"fmt"
	"os/exec"
	"strings"
)

// clone a repository to the given path
func Clone(repository string, path string) error {

	commandStr := fmt.Sprintf("git clone %s %s", repository, path)
	fmt.Println(commandStr)

	command := exec.Command("git", "clone", repository, path)

	output, err := command.Output()

	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println(output)

	return nil

}

// replace the host in the git ssh url for the given host in the ${host}
// parameter
func ReplaceHost(repository string, host string) string {
	var (
		atSign    = strings.Index(repository, "@")
		colonSign = strings.Index(repository, ":")
		oldHost   = repository[atSign+1 : colonSign]
	)

	return strings.Replace(repository, oldHost, host, 1)
}
