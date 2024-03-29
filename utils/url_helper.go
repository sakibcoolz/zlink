package utils

import (
	"fmt"
	"os"
	"strings"
	"zlink/literals"
)

// filter for url format return boolean result for url.
func UrlNotAllowed(url string) bool {

	if strings.Contains(url, literals.LOCALHOST) {
		return true
	}

	if strings.Contains(url, literals.DOUBLEQT) {
		return true
	}

	if strings.Contains(url, literals.SINGLEQT) {
		return true
	}

	if strings.Contains(url, literals.TICK) {
		return true
	}

	return false
}

// modify url if it missed http://
func ModifiyUrl(url string) string {

	if !strings.HasPrefix(url, literals.HTTP) {
		url = fmt.Sprintf("http://%s", url)
	}

	return url
}

func UrlMaker(path string) string {
	return fmt.Sprintf("http://%s:%s/%s", os.Getenv("SERVICEHOST"),
		os.Getenv("SERVICEPORT"),
		path)
}
