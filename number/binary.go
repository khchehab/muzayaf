package number

import (
	"github.com/khchehab/muzayaf/random"
	"strconv"
)

// Binary generates a random binary number as a string based on the provided options
func Binary(opts ...OptionFunc) string {
	o := applyOptions(opts)

	// Validate options
	if o.binaryMin < 0 {
		o.binaryMin = 0
	}
	if o.binaryMax < o.binaryMin {
		o.binaryMin, o.binaryMax = o.binaryMax, o.binaryMin
	}

	// Generate a random number in the range
	var value int
	if o.binaryMin == o.binaryMax {
		value = o.binaryMin
	} else {
		value = o.binaryMin + random.IntN(o.binaryMax-o.binaryMin+1)
	}

	// Convert to binary string
	binaryStr := strconv.FormatInt(int64(value), 2)

	// Add prefix if requested
	if o.binaryPrefix {
		return binaryPrefix + binaryStr
	}

	return binaryStr
}
