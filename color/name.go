package color

import (
	"github.com/khchehab/muzayaf/internal"
	"github.com/khchehab/muzayaf/random"
)

// ColorName generates a random color name
// It can use a specific locale if specified in the options (default is "en")
func ColorName(opts ...OptionFunc) string {
	o := applyOptions(opts)

	// Validate locale
	if _, exists := fallbackValues[o.locale]; !exists {
		// If locale doesn't exist in fallbackValues, use "en" as fallback
		o.locale = "en"
	}

	data, err := internal.LoadJsonFile("color", o.locale, "names.json")
	if err != nil {
		return fallbackValues[o.locale]["name"]
	}

	pool := internal.GetStringSlice(data, "colors")
	if len(pool) == 0 {
		return fallbackValues[o.locale]["name"]
	}

	return pool[random.IntN(len(pool))]
}
