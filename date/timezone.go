package date

import (
	"github.com/khchehab/muzayaf/internal"
	"github.com/khchehab/muzayaf/random"
)

// Timezone generates a random timezone
// The list of timezones is compliant with IANA
func Timezone(opts ...OptionFunc) string {
	o := applyOptions(opts)

	// Validate locale
	if _, exists := fallbackValues[o.locale]; !exists {
		// If locale doesn't exist in fallbackValues, use "en" as fallback
		o.locale = "en"
	}

	// Try to load timezones from the specified locale
	data, err := internal.LoadJsonFile("date", o.locale, "timezones.json")
	if err != nil {
		// If not found in the specified locale, try to load from the base locale
		data, err = internal.LoadJsonFile("date", "base", "timezones.json")
		if err != nil {
			return fallbackValues[o.locale]["timezone"]
		}
	}

	pool := internal.GetStringSlice(data, "timezones")
	if len(pool) == 0 {
		return fallbackValues[o.locale]["timezone"]
	}

	return pool[random.IntN(len(pool))]
}
