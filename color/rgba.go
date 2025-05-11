package color

import (
	"fmt"
	"github.com/khchehab/muzayaf/random"
)

// RGBAColor represents a color in the RGBA color space
type RGBAColor struct {
	Red   int
	Green int
	Blue  int
	Alpha float64
}

// String returns a string representation of the RGBA color
func (r RGBAColor) String() string {
	if r.Alpha == 1.0 {
		return fmt.Sprintf("rgb(%d, %d, %d)",
			r.Red, r.Green, r.Blue)
	}
	return fmt.Sprintf("rgba(%d, %d, %d, %.2f)",
		r.Red, r.Green, r.Blue, r.Alpha)
}

// RGBA generates a random RGBA color
func RGBA() RGBAColor {
	return RGBAColor{
		Red:   random.IntN(256),                  // 0-255
		Green: random.IntN(256),                  // 0-255
		Blue:  random.IntN(256),                  // 0-255
		Alpha: float64(random.IntN(101)) / 100.0, // 0-1 inclusive
	}
}

// RGB generates a random RGB color with full alpha (1.0)
func RGB() RGBAColor {
	return RGBAColor{
		Red:   random.IntN(256), // 0-255
		Green: random.IntN(256), // 0-255
		Blue:  random.IntN(256), // 0-255
		Alpha: 1.0,              // Full alpha
	}
}
