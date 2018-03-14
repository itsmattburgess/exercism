// Calculates the moment someone is a gigasecond old
package gigasecond

import (
	"time"
)

// Return the time after one gigasecond (10^9) has been added
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * 1e9)
}
