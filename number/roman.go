package number

import (
	"github.com/khchehab/muzayaf/random"
	"strings"
)

// Roman numeral mapping
var romanNumerals = []struct {
	Value  int
	Symbol string
}{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

// Roman generates a random Roman numeral as a string based on the provided options
func Roman(opts ...OptionFunc) string {
	o := applyOptions(opts)

	// Validate options - Roman numerals are only valid from 1 to 3999
	if o.romanMin < 1 {
		o.romanMin = 1
	}
	if o.romanMax > 3999 {
		o.romanMax = 3999
	}
	if o.romanMax < o.romanMin {
		o.romanMin, o.romanMax = o.romanMax, o.romanMin
	}

	// Generate a random number in the range
	var value int
	if o.romanMin == o.romanMax {
		value = o.romanMin
	} else {
		value = o.romanMin + random.IntN(o.romanMax-o.romanMin+1)
	}

	// Convert to Roman numeral
	return intToRoman(value)
}

// intToRoman converts an integer to a Roman numeral
func intToRoman(num int) string {
	if num < 1 || num > 3999 {
		return ""
	}

	var result strings.Builder

	for _, numeral := range romanNumerals {
		for num >= numeral.Value {
			result.WriteString(numeral.Symbol)
			num -= numeral.Value
		}
	}

	return result.String()
}
