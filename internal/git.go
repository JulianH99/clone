package internal

import (
	"context"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

// clone a repository to the given path
func Clone(ctx context.Context, repository string, path string) error {
	command := exec.CommandContext(ctx, "git", "clone", repository)

	if path != "" {
		command = exec.CommandContext(ctx, "git", "clone", repository, path)
	}
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	err := command.Run()
	return err
}

// replace the host in the git ssh url for the given host in the ${host}
// parameter
// example: git@github.com:JulianH99/clone.git -> git@${host}:JulianH99/clone.git
func ReplaceHost(repository string, host string) string {
	if host == "" {
		return repository
	}

	var (
		atSign    = strings.Index(repository, "@")
		colonSign = strings.Index(repository, ":")
		oldHost   = repository[atSign+1 : colonSign]
	)

	return strings.Replace(repository, oldHost, host, 1)
}

// checks if the given url is a valid ssh url
// with the following format: git@${host}:${user}/${repo}.git
func CheckValidSshUrl(url string) bool {
	re := regexp.MustCompile(`(?i)git@[\w+.]+:\w+\/\w+\.git`)

	return re.Match([]byte(url))
}
