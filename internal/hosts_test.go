package internal

import (
	"testing"
)

func TestParseConfigFileWithTabbedContent(t *testing.T) {
	content := `
	Host *
		Hostname localhost
		User root
		IdentityFile ~/.ssh/id_rsa

	Host github.com
		Hostname github.com
		User git
		IdentityFile ~/.ssh/id_rsa

	Host github.com
	Hostname github.com
	User git
	IdentityFile ~/.ssh/id_rsa

	`

	hosts := parseConfigFile([]byte(content))

	if len(hosts) != 3 {
		t.Errorf("expected 3 hosts got %d", len(hosts))
	}
}

func TestParseConfigFileWithEmptyHosts(t *testing.T) {
	content := `
	Host 
		Hostname localhost
		User root
		IdentityFile ~/.ssh/id_rsa
	`

	hosts := parseConfigFile([]byte(content))

	if len(hosts) != 0 {
		t.Errorf("expected 0 hosts got %d", len(hosts))
	}
}

func TestParseConfigWithEmpyString(t *testing.T) {
	hosts := parseConfigFile([]byte(""))

	if len(hosts) != 0 {
		t.Errorf("expected 0 hosts got %d", len(hosts))
	}
}

func TestParseConfigWithLowerCaseSettings(t *testing.T) {
	content := `
	host github.com
	hostname github.com
	user git
	identityfile ~/.ssh/id_rsa
	`

	hosts := parseConfigFile([]byte(content))

	if len(hosts) != 1 {
		t.Errorf("expected 1 hosts got %d", len(hosts))
	}
}
