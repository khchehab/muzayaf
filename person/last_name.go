package person

import (
	"github.com/khchehab/muzayaf/internal"
	"github.com/khchehab/muzayaf/random"
)

// LastName generates a random last name based on the provided options
// It returns a last name from the available last names in the specified locale
func LastName(opts ...OptionFunc) string {
	o := applyOptions(opts)

	data, err := internal.LoadJsonFile("person", o.locale, "last_names.json")
	if err != nil {
		return fallbackValues[o.locale]["last_name"]
	}

	pool := internal.GetStringSlice(data, "last_names")
	if len(pool) == 0 {
		return fallbackValues[o.locale]["last_name"]
	}

	return pool[random.IntN(len(pool))]
}
