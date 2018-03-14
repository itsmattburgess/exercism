// Package acronym converts phrases to acronyms
package acronym

import (
	"strings"
)

// Abbreviate a string to the first letter from each word
func Abbreviate(s string) string {
	var acronym string

	for _, word := range strings.Split(s, " ") {
		for _, segment := range strings.Split(word, "-") {
			acronym += string(segment[0])
		}
	}

	return strings.ToUpper(acronym)
}
