package stringutils

import (
	"regexp"
	"strings"
)

// Upper returns the uppercase of the given string argument.
func Upper(s string) string {
	return strings.ToUpper(s)
}

// Lower returns the lowercase of the given string argument.
func Lower(s string) string {
	return strings.ToLower(s)
}

// WordCountByRegex returns the word count of a given string
func WordCountByRegex(s string) int {
	pattern := regexp.MustCompile(`[\s.]+`)
	splitStringSlice := pattern.Split(s, -1)

	return len(splitStringSlice) - 1
}

// WordCountByString returns the word count of a given string
func WordCountByString(s string) int {
	splitStringSlice := strings.Split(s, " ")

	return len(splitStringSlice)
}
