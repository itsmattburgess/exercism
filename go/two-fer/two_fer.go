// Package twofer will specify how an item is shared between two people.
package twofer

import (
	"fmt"
)

// ShareWith returns who something is being shared with.
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}

	return fmt.Sprintf("One for %s, one for me.", name)
}
