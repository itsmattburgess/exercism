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

type Remark struct {
	remark string
}

// Speak with the teenager and return their response
func Hey(remark string) string {
	r := Remark{strings.TrimSpace(remark)}

	switch {
	case r.isEmpty():
		return fine
	case r.isShouting() && r.isQuestion():
		return calmdown
	case r.isShouting():
		return chillout
	case r.isQuestion():
		return sure
	default:
		return whatever
	}
}

// Is the string blank?
func (r Remark) isEmpty() bool {
	return r.remark == ""
}

// Does the string contain letters
func (r Remark) containsLetters() bool {
	for _, letter := range r.remark {
		if unicode.IsLetter(letter) {
			return true
		}
	}

	return false
}

// Is the string all in caps?
func (r Remark) isShouting() bool {
	return r.containsLetters() && strings.ToUpper(r.remark) == r.remark
}

// Does the string end with a question mark?
func (r Remark) isQuestion() bool {
	return strings.HasSuffix(r.remark, "?")
}
