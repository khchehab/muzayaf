package date

import (
	"time"
)

// Past generates a random date in the past
// It can use a relative date if specified in the options (default is today)
// It can also specify how many years, months, or days to go back
func Past(opts ...OptionFunc) time.Time {
	o := applyOptions(opts)

	// Calculate the maximum past date (1 day before the relative date)
	maxDate := o.relative.AddDate(0, 0, -1)

	// Calculate the minimum past date based on the specified years, months, and days
	minDate := o.relative.AddDate(-o.years, -o.months, -o.days)

	// Ensure minDate is before maxDate (in case negative values were provided)
	if minDate.After(maxDate) {
		minDate = o.relative.AddDate(-1, 0, 0) // Default to 1 year in the past
	}

	// Generate a random date between the minimum and maximum dates
	return Between(minDate, maxDate)
}
