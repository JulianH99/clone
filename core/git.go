package core

import (
	"fmt"
	"os/exec"
	"path"
	"strings"
)

var format string = "git@github.com-%s:%s.git"

func CloneRepository(workspace Workspace, url string, subfolder string) error {
	var destinationPath string
	projectName := strings.Split(url, "/")[1]

	if subfolder == "" {
		destinationPath = path.Join(workspace.Path, projectName)
	} else {
		destinationPath = path.Join(workspace.Path, subfolder, projectName)
	}

	githubUrl := fmt.Sprintf(format, workspace.Hostname, url)

	command := exec.Command("git", "clone", githubUrl, destinationPath)
	stdOut, err := command.StdoutPipe()

	command.Stderr = command.Stdout

	if err != nil {
		return err
	}

	if err := command.Start(); err != nil {
		return err
	}

	for {
		temp := make([]byte, 1024)
		_, err := stdOut.Read(temp)

		fmt.Println(string(temp))
		if err != nil {
			break
		}
	}

	return nil

}
