package person

import (
	"github.com/khchehab/muzayaf/internal"
	"github.com/khchehab/muzayaf/random"
)

// Suffix generates a random name suffix based on the provided options
// It can filter by suffix type (academic or generational) if specified in the options
func Suffix(opts ...OptionFunc) string {
	o := applyOptions(opts)

	data, err := internal.LoadJsonFile("person", o.locale, "suffixes.json")
	if err != nil {
		return fallbackValues[o.locale]["suffix"]
	}

	var pool []string
	if o.suffixType == SuffixTypeAcademic {
		pool = internal.GetStringSlice(data, SuffixTypeAcademic)
	} else if o.suffixType == SuffixTypeGenerational {
		pool = internal.GetStringSlice(data, SuffixTypeGenerational)
	} else {
		pool = append(internal.GetStringSlice(data, SuffixTypeAcademic), internal.GetStringSlice(data, SuffixTypeGenerational)...)
	}

	if len(pool) == 0 {
		return fallbackValues[o.locale]["suffix"]
	}

	return pool[random.IntN(len(pool))]
}
