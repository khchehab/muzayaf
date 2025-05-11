package number

import (
	"github.com/khchehab/muzayaf/random"
	"strconv"
)

// Octal generates a random octal number as a string based on the provided options
func Octal(opts ...OptionFunc) string {
	o := applyOptions(opts)

	// Validate options
	if o.octalMin < 0 {
		o.octalMin = 0
	}
	if o.octalMax < o.octalMin {
		o.octalMin, o.octalMax = o.octalMax, o.octalMin
	}

	// Generate a random number in the range
	var value int
	if o.octalMin == o.octalMax {
		value = o.octalMin
	} else {
		value = o.octalMin + random.IntN(o.octalMax-o.octalMin+1)
	}

	// Convert to octal string
	octalStr := strconv.FormatInt(int64(value), 8)

	// Add prefix if requested
	if o.octalPrefix {
		return octalPrefix + octalStr
	}

	return octalStr
}
