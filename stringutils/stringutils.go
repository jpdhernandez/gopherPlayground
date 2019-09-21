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

// WordCount returns the word count of a given string
func WordCount(s string) int {
	pattern := regexp.MustCompile(`[\s.]+`)
	splitStringSlice := pattern.Split(s, -1)

	return len(splitStringSlice) - 1
}
