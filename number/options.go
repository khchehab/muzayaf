package number

import (
	"math"
)

// Option struct holds configuration for number generation
type Option struct {
	// Int options
	intMin      int
	intMax      int
	intMultiple int

	// Float options
	floatMin            float64
	floatMax            float64
	floatFractionDigits int

	// Binary options
	binaryMin    int
	binaryMax    int
	binaryPrefix bool

	// Octal options
	octalMin    int
	octalMax    int
	octalPrefix bool

	// Hex options
	hexMin    int
	hexMax    int
	hexPrefix bool

	// Roman numeral options
	romanMin int
	romanMax int
}

type OptionFunc func(*Option)

// defaultOption returns the default configuration
func defaultOption() Option {
	return Option{
		// Int defaults
		intMin:      0,
		intMax:      math.MaxInt,
		intMultiple: 1,

		// Float defaults
		floatMin:            0.0,
		floatMax:            1.0,
		floatFractionDigits: 2,

		// Binary defaults
		binaryMin:    0,
		binaryMax:    1,
		binaryPrefix: false,

		// Octal defaults
		octalMin:    0,
		octalMax:    7,
		octalPrefix: false,

		// Hex defaults
		hexMin:    0,
		hexMax:    15,
		hexPrefix: false,

		// Roman numeral defaults
		romanMin: 1,
		romanMax: 3999,
	}
}

// applyOptions applies the provided option functions to the default options
func applyOptions(opts []OptionFunc) Option {
	opt := defaultOption()
	for _, o := range opts {
		o(&opt)
	}
	return opt
}

// WithIntMin sets the minimum value for random integers
func WithIntMin(min int) OptionFunc {
	return func(o *Option) {
		o.intMin = min
	}
}

// WithIntMax sets the maximum value for random integers
func WithIntMax(max int) OptionFunc {
	return func(o *Option) {
		o.intMax = max
	}
}

// WithIntMultiple sets the multiple value for random integers
func WithIntMultiple(multiple int) OptionFunc {
	return func(o *Option) {
		if multiple > 0 {
			o.intMultiple = multiple
		}
	}
}

// WithFloatMin sets the minimum value for random floats
func WithFloatMin(min float64) OptionFunc {
	return func(o *Option) {
		o.floatMin = min
	}
}

// WithFloatMax sets the maximum value for random floats
func WithFloatMax(max float64) OptionFunc {
	return func(o *Option) {
		o.floatMax = max
	}
}

// WithFloatFractionDigits sets the number of fraction digits for random floats
func WithFloatFractionDigits(digits int) OptionFunc {
	return func(o *Option) {
		if digits >= 0 {
			o.floatFractionDigits = digits
		}
	}
}

// WithBinaryMin sets the minimum value for random binary numbers
func WithBinaryMin(min int) OptionFunc {
	return func(o *Option) {
		o.binaryMin = min
	}
}

// WithBinaryMax sets the maximum value for random binary numbers
func WithBinaryMax(max int) OptionFunc {
	return func(o *Option) {
		o.binaryMax = max
	}
}

// WithBinaryPrefix sets whether to include the binary prefix (0b)
func WithBinaryPrefix(include bool) OptionFunc {
	return func(o *Option) {
		o.binaryPrefix = include
	}
}

// WithOctalMin sets the minimum value for random octal numbers
func WithOctalMin(min int) OptionFunc {
	return func(o *Option) {
		o.octalMin = min
	}
}

// WithOctalMax sets the maximum value for random octal numbers
func WithOctalMax(max int) OptionFunc {
	return func(o *Option) {
		o.octalMax = max
	}
}

// WithOctalPrefix sets whether to include the octal prefix (0)
func WithOctalPrefix(include bool) OptionFunc {
	return func(o *Option) {
		o.octalPrefix = include
	}
}

// WithHexMin sets the minimum value for random hexadecimal numbers
func WithHexMin(min int) OptionFunc {
	return func(o *Option) {
		o.hexMin = min
	}
}

// WithHexMax sets the maximum value for random hexadecimal numbers
func WithHexMax(max int) OptionFunc {
	return func(o *Option) {
		o.hexMax = max
	}
}

// WithHexPrefix sets whether to include the hexadecimal prefix (0x)
func WithHexPrefix(include bool) OptionFunc {
	return func(o *Option) {
		o.hexPrefix = include
	}
}

// WithRomanMin sets the minimum value for random roman numerals
func WithRomanMin(min int) OptionFunc {
	return func(o *Option) {
		if min < 1 {
			min = 1
		}
		o.romanMin = min
	}
}

// WithRomanMax sets the maximum value for random roman numerals
func WithRomanMax(max int) OptionFunc {
	return func(o *Option) {
		if max > 3999 {
			max = 3999
		}
		o.romanMax = max
	}
}
