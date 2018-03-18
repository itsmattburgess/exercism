// Package isogram determines is a word or phrase is an isogram
package isogram

import (
	"strings"
)

// IsIsogram returns true if the word is an isogram
func IsIsogram(word string) bool {
	charCount := map[string]int{}
	word = strings.ToLower(word)

	for _, letter := range strings.Split(word, "") {
		if _, exists := charCount[letter]; exists {
			charCount[letter]++
		} else {
			charCount[letter] = 1
		}

		if charCount[letter] > 1 && letter != "-" && letter != " " {
			return false
		}
	}

	return true
}
