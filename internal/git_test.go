package internal

import "testing"

func TestReplaceHostEmpyHost(t *testing.T) {
	testUrl := "git@github.com:JulianH99/clone.git"
	host := ""

	if ReplaceHost(testUrl, host) != testUrl {
		t.Errorf("expected %s got %s", testUrl, ReplaceHost(testUrl, host))
	}
}

func TestReplaceHost(t *testing.T) {
	testUrl := "git@github.com:JulianH99/clone.git"
	host := "github.com-personal"

	if ReplaceHost(testUrl, host) != "git@github.com-personal:JulianH99/clone.git" {
		t.Errorf("expected %s got %s", "git@github.com-personal:JulianH99/clone.git", ReplaceHost(testUrl, host))
	}
}

func TestCheckValidSshUrl(t *testing.T) {
	testUrl := "git@github.com:JulianH99/clone.git"

	if !CheckValidSshUrl(testUrl) {
		t.Errorf("expected %t got %t", true, CheckValidSshUrl(testUrl))
	}
}

func TestCheckValidSshUrlEmptyUrl(t *testing.T) {
	testUrl := ""

	if CheckValidSshUrl(testUrl) {
		t.Errorf("expected %t got %t", false, CheckValidSshUrl(testUrl))
	}
}

func TestCheckValidSshUrlInvalidUrl(t *testing.T) {
	testUrl := "github.com:JulianH99/clone.git"

	if CheckValidSshUrl(testUrl) {
		t.Errorf("expected %t got %t", false, CheckValidSshUrl(testUrl))
	}
}
