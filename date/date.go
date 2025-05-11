// Package date provides functionality for generating random date-related data
// such as dates, months, weekdays, and timezones.
package date

var (
	fallbackValues = map[string]map[string]string{
		"en": {
			"timezone": "America/New_York",
			"weekday":  "Monday",
			"month":    "January",
		},
	}
)
