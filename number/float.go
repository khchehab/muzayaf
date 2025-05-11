package number

import (
	"github.com/khchehab/muzayaf/random"
	"math"
)

// Float generates a random float64 based on the provided options
func Float(opts ...OptionFunc) float64 {
	o := applyOptions(opts)

	// Validate options
	if o.floatMin > o.floatMax {
		o.floatMin, o.floatMax = o.floatMax, o.floatMin
	}

	// If min equals max, return min
	if o.floatMin == o.floatMax {
		return o.floatMin
	}

	// Generate a random float in the range [0, 1)
	randomFloat := random.Float64()

	// Scale to the desired range [min, max)
	result := o.floatMin + randomFloat*(o.floatMax-o.floatMin)

	// Round to the specified number of fraction digits if needed
	if o.floatFractionDigits >= 0 {
		multiplier := math.Pow10(o.floatFractionDigits)
		result = math.Round(result*multiplier) / multiplier
	}

	return result
}
