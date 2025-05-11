package number

import (
	"github.com/khchehab/muzayaf/random"
	"math"
)

// Int generates a random integer based on the provided options
func Int(opts ...OptionFunc) int {
	o := applyOptions(opts)

	// Validate options
	if o.intMin > o.intMax {
		o.intMin, o.intMax = o.intMax, o.intMin
	}

	// Adjust range to be a multiple of intMultiple
	if o.intMultiple > 1 {
		// Adjust minimum to the next multiple >= intMin
		if o.intMin%o.intMultiple != 0 {
			o.intMin = ((o.intMin / o.intMultiple) + 1) * o.intMultiple
		}

		// Adjust maximum to the previous multiple <= intMax
		o.intMax = (o.intMax / o.intMultiple) * o.intMultiple

		// If adjustments made min > max, set them equal
		if o.intMin > o.intMax {
			o.intMin = o.intMax
		}
	}

	// If min equals max, return min
	if o.intMin == o.intMax {
		return o.intMin
	}

	// Generate random number
	if o.intMultiple > 1 {
		// Calculate how many multiples are in the range
		count := (o.intMax-o.intMin)/o.intMultiple + 1
		// Generate a random index and convert to the actual value
		return o.intMin + (random.IntN(count) * o.intMultiple)
	}

	// Handle the case when intMax is math.MaxInt to prevent overflow
	if o.intMax == math.MaxInt {
		// Use math.MaxInt-1 to prevent overflow when adding 1
		return o.intMin + random.IntN(math.MaxInt-o.intMin)
	}

	// Standard case: generate a random number in the range
	return o.intMin + random.IntN(o.intMax-o.intMin+1)
}
