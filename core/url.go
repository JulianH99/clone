package core

import (
	"regexp"
)

var urlRegex = regexp.MustCompile("./.")

func CheckValidUrl(url string) bool {
	return urlRegex.MatchString(url)
}
