package date

import (
	"time"
)

// Future generates a random date in the future
// It can use a relative date if specified in the options (default is today)
// It can also specify how many years, months, or days to go forward
func Future(opts ...OptionFunc) time.Time {
	o := applyOptions(opts)

	// Calculate the minimum future date (1 day after the relative date)
	minDate := o.relative.AddDate(0, 0, 1)

	// Calculate the maximum future date based on the specified years, months, and days
	maxDate := o.relative.AddDate(o.years, o.months, o.days)

	// Ensure maxDate is after minDate (in case negative values were provided)
	if maxDate.Before(minDate) {
		maxDate = o.relative.AddDate(1, 0, 0) // Default to 1 year in the future
	}

	// Generate a random date between the minimum and maximum dates
	return Between(minDate, maxDate)
}
