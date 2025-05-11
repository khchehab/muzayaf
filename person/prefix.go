package person

import (
	"github.com/khchehab/muzayaf/internal"
	"github.com/khchehab/muzayaf/random"
)

// Prefix generates a random name prefix (e.g., Mr., Mrs., Dr.) based on the provided options
// It can filter by gender if specified in the options
func Prefix(opts ...OptionFunc) string {
	o := applyOptions(opts)

	data, err := internal.LoadJsonFile("person", o.locale, "prefixes.json")
	if err != nil {
		return fallbackValues[o.locale]["prefix"]
	}

	var pool []string
	if o.gender == GenderFemale {
		pool = internal.GetStringSlice(data, GenderFemale)
	} else if o.gender == GenderMale {
		pool = internal.GetStringSlice(data, GenderMale)
	} else {
		pool = append(internal.GetStringSlice(data, GenderFemale), internal.GetStringSlice(data, GenderMale)...)
	}

	if len(pool) == 0 {
		return fallbackValues[o.locale]["prefix"]
	}

	return pool[random.IntN(len(pool))]
}
