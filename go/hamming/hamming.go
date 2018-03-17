// Package hamming calculates the hamming distance between two strings
package hamming

import (
	"errors"
	"strings"
)

// Distance returns the hamming distance between two strings
func Distance(a, b string) (int, error) {
	distance := 0

	strandA := strings.Split(a, "")
	strandB := strings.Split(b, "")

	if len(strandA) != len(strandB) {
		return -1, errors.New("Each strand must be an equal length")
	}

	for i, c := range strandA {
		if strandB[i] != c {
			distance++
		}
	}

	return distance, nil
}
