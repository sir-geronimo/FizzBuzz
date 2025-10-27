package anagram

import (
	"regexp"
	"strings"
)

var alphaRegex = regexp.MustCompile(`(?i)[^a-z0-9A-Z]+`)

// IsAnagram returns true if a and b are anagrams of each other.
func IsAnagram(a, b string) bool {
	// normalise strings to lowercase and remove all non-alpha characters
	a = alphaRegex.ReplaceAllString(strings.ToLower(a), "")
	b = alphaRegex.ReplaceAllString(strings.ToLower(b), "")

	// if strings are different length, they are not anagrams
	if len(a) != len(b) {
		return false
	}

	// create a map of runes and their counts
	letters := make(map[rune]int, len(a))
	for _, r := range a {
		letters[r]++
	}

	// decrement counts for each rune in b
	// if any count is negative, then a and b are not anagrams
	for _, r := range b {
		letters[r]--
		if letters[r] < 0 {
			return false
		}
	}

	// double check that all counts are zero
	for _, c := range letters {
		if c != 0 {
			return false
		}
	}

	return true
}
