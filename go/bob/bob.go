// Package bob allows you to talk to a stroppy teenager
package bob

import (
	"strings"
	"unicode"
)

const (
	sure     = "Sure."
	chillout = "Whoa, chill out!"
	calmdown = "Calm down, I know what I'm doing!"
	fine     = "Fine. Be that way!"
	whatever = "Whatever."
)

// Speak with the teenager and return their response
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)

	switch {
	case isEmpty(remark):
		return fine
	case isShouting(remark) && isQuestion(remark):
		return calmdown
	case isShouting(remark):
		return chillout
	case isQuestion(remark):
		return sure
	default:
		return whatever
	}
}

// Is the string blank?
func isEmpty(remark string) bool {
	return remark == ""
}

// Does the string contain letters
func containsLetters(remark string) bool {
	for _, letter := range remark {
		if unicode.IsLetter(letter) {
			return true
		}
	}

	return false
}

// Is the string all in caps?
func isShouting(remark string) bool {
	return containsLetters(remark) && strings.ToUpper(remark) == remark
}

// Does the string end with a question mark?
func isQuestion(remark string) bool {
	return strings.HasSuffix(remark, "?")
}
