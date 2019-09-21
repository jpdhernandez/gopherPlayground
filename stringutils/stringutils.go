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
	return len(strings.Fields(s))
}

// GetCharCountMap returns the char count of a given string
func GetCharCountMap(s string) map[string]int {
	splitStringSlice := strings.Split(s, "")
	stringMap := make(map[string]int)

	for _, char := range splitStringSlice {
		// get count of iteration
		stringMap[char] = strings.Count(s, char)
	}

	return stringMap
}
