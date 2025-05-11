package number

import (
	"github.com/khchehab/muzayaf/random"
	"strconv"
	"strings"
)

// Hex generates a random hexadecimal number as a string based on the provided options
func Hex(opts ...OptionFunc) string {
	o := applyOptions(opts)

	// Validate options
	if o.hexMin < 0 {
		o.hexMin = 0
	}
	if o.hexMax < o.hexMin {
		o.hexMin, o.hexMax = o.hexMax, o.hexMin
	}

	// Generate a random number in the range
	var value int
	if o.hexMin == o.hexMax {
		value = o.hexMin
	} else {
		value = o.hexMin + random.IntN(o.hexMax-o.hexMin+1)
	}

	// Convert to hexadecimal string
	hexStr := strconv.FormatInt(int64(value), 16)

	// Convert to lowercase for consistency
	hexStr = strings.ToLower(hexStr)

	// Add prefix if requested
	if o.hexPrefix {
		return hexPrefix + hexStr
	}

	return hexStr
}
