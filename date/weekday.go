package date

import (
	"github.com/khchehab/muzayaf/internal"
	"github.com/khchehab/muzayaf/random"
)

// Weekday returns a random weekday name
// It can use a specific locale if specified in the options (default is "en")
func Weekday(opts ...OptionFunc) string {
	o := applyOptions(opts)

	// Validate locale
	if _, exists := fallbackValues[o.locale]; !exists {
		// If locale doesn't exist in fallbackValues, use "en" as fallback
		o.locale = "en"
	}

	// Try to load weekdays from the specified locale
	data, err := internal.LoadJsonFile("date", o.locale, "weekdays.json")
	if err != nil {
		return fallbackValues[o.locale]["weekday"]
	}

	pool := internal.GetStringSlice(data, "weekdays")
	if len(pool) == 0 {
		return fallbackValues[o.locale]["weekday"]
	}

	return pool[random.IntN(len(pool))]
}
