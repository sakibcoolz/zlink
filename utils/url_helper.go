package utils

import (
	"fmt"
	"strings"
	"zlink/literals"
)

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

func ModifiyUrl(url string) string {

	if !strings.HasPrefix(url, literals.HTTP) {
		url = fmt.Sprintf("http://%s", url)
	}

	return url
}
