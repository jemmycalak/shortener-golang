package usecase

import "strings"

func IsContainHttps(url string) string {
	var https = "https://"
	if strings.Contains(url, https) {
		return url
	}
	return https + url
}
