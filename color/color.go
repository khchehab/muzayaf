// Package color provides functionality for generating random colors in different color spaces.
// It supports RGB, HSL, CMYK color spaces and named colors with localization options.
package color

var (
	fallbackValues = map[string]map[string]string{
		"en": {
			"name": "white",
		},
	}
)
