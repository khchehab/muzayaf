package person

import (
	"github.com/khchehab/muzayaf/internal"
	"github.com/khchehab/muzayaf/random"
)

// Gender generates a random gender based on the provided options
// It returns a gender from the available genders in the specified locale
func Gender(opts ...OptionFunc) string {
	o := applyOptions(opts)

	data, err := internal.LoadJsonFile("person", o.locale, "genders.json")
	if err != nil {
		return fallbackValues[o.locale]["gender"]
	}

	pool := internal.GetStringSlice(data, "genders")
	if len(pool) == 0 {
		return fallbackValues[o.locale]["gender"]
	}

	return pool[random.IntN(len(pool))]
}
