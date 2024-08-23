package utils

import "regexp"

func MatchString(input string, pattern string) bool {
	matched, err := regexp.MatchString(pattern, input)
	if err != nil {
		return false
	}
	return matched
}
