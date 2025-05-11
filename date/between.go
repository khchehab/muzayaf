package date

import (
	"github.com/khchehab/muzayaf/random"
	"time"
)

// Between generates a random date between two dates
func Between(from, to time.Time) time.Time {
	// Ensure from is before to
	if from.After(to) {
		from, to = to, from
	}

	// Calculate the difference in seconds between the two dates
	diff := to.Unix() - from.Unix()

	// If the dates are the same, return the date
	if diff == 0 {
		return from
	}

	// Generate a random number of seconds within the range
	randomSeconds := random.Int64N(diff + 1)

	// Add the random number of seconds to the from date
	return from.Add(time.Duration(randomSeconds) * time.Second)
}
