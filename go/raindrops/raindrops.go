// Package raindrops converts a number to a string based on its factors
package raindrops

import (
	"strconv"
)

// Convert returns a string based on the factors of the number
func Convert(num int) string {
	output := ""

	if num%3 == 0 {
		output += "Pling"
	}

	if num%5 == 0 {
		output += "Plang"
	}

	if num%7 == 0 {
		output += "Plong"
	}

	if output != "" {
		return output
	}

	return strconv.Itoa(num)
}
