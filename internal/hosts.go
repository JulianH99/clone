package internal

import (
	"os"
	"path"
	"strings"
)

type host string

func parseConfigFile(contents []byte) []host {
	lines := strings.Split(string(contents), "\n")
	hosts := make([]host, 0)

	for _, line := range lines {
		if strings.HasPrefix(line, "Host ") {
			parts := strings.Split(line, " ")
			h := strings.TrimSpace(parts[1])
			hosts = append(hosts, host(h))
		}
	}

	return hosts
}

func readSshConfigFile() ([]byte, error) {
	homedir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	sshConfigPath := path.Join(homedir, ".ssh", "config")
	contents, err := os.ReadFile(sshConfigPath)

	if err != nil {
		return nil, err
	}
	return contents, nil
}

func SshHosts() ([]host, error) {
	sshConfigContents, err := readSshConfigFile()

	if err != nil {
		return nil, err
	}
	hosts := parseConfigFile(sshConfigContents)

	return hosts, nil
}
