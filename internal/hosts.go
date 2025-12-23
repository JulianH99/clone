package internal

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func parseConfigFile(contents []byte) []string {
	lines := strings.Split(string(contents), "\n")
	hosts := make([]string, 0)
	hostsRegex := regexp.MustCompile(`(?i)host\s`)

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if len(line) > 0 && line[0] == '#' {
			continue
		}
		if hostsRegex.Match([]byte(line)) {
			parts := strings.Split(line, " ")
			h := strings.TrimSpace(parts[1])
			hosts = append(hosts, h)
		}
	}

	return hosts
}

func readSshConfigFile() ([]byte, error) {
	homedir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	sshConfigPath := filepath.Join(homedir, ".ssh", "config")
	contents, err := os.ReadFile(sshConfigPath)
	if err != nil {
		return nil, err
	}
	return contents, nil
}

func SshHosts() ([]string, error) {
	sshConfigContents, err := readSshConfigFile()
	if err != nil {
		return nil, err
	}
	hosts := parseConfigFile(sshConfigContents)

	return hosts, nil
}
