package date

import (
	"github.com/khchehab/muzayaf/internal"
	"github.com/khchehab/muzayaf/random"
)

// Month returns a random month name
// It can use a specific locale if specified in the options (default is "en")
func Month(opts ...OptionFunc) string {
	o := applyOptions(opts)

	// Validate locale
	if _, exists := fallbackValues[o.locale]; !exists {
		// If locale doesn't exist in fallbackValues, use "en" as fallback
		o.locale = "en"
	}

	// Try to load months from the specified locale
	data, err := internal.LoadJsonFile("date", o.locale, "months.json")
	if err != nil {
		return fallbackValues[o.locale]["month"]
	}

	pool := internal.GetStringSlice(data, "months")
	if len(pool) == 0 {
		return fallbackValues[o.locale]["month"]
	}

	return pool[random.IntN(len(pool))]
}
