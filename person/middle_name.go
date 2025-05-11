package person

import (
	"github.com/khchehab/muzayaf/internal"
	"github.com/khchehab/muzayaf/random"
)

// MiddleName generates a random middle name based on the provided options
// It can filter by gender if specified in the options
func MiddleName(opts ...OptionFunc) string {
	o := applyOptions(opts)

	data, err := internal.LoadJsonFile("person", o.locale, "middle_names.json")
	if err != nil {
		return fallbackValues[o.locale]["middle_name"]
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
		return fallbackValues[o.locale]["middle_name"]
	}

	return pool[random.IntN(len(pool))]
}
