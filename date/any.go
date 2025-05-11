package date

import (
	"time"
)

// Any generates a random date
// It can use a relative date if specified in the options
// It can also specify how many years, months, or days to go back or forward
func Any(opts ...OptionFunc) time.Time {
	o := applyOptions(opts)

	// Calculate the start and end dates based on the specified years, months, and days
	startDate := o.relative.AddDate(-o.years, -o.months, -o.days)
	endDate := o.relative.AddDate(o.years, o.months, o.days)

	// Generate a random date between the start and end dates
	return Between(startDate, endDate)
}
