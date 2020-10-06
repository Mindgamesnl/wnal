package utils

import "regexp"

func ValidatePort(port string) bool {
	re := regexp.MustCompile("^([0-9]{1,4}|[1-5][0-9]{4}|6[0-4][0-9]{3}|65[0-4][0-9]{2}|655[0-2][0-9]|6553[0-5])$")
	if re.Match([]byte(port)) {
		return true
	}
	return false
}
