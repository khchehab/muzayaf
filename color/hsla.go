package color

import (
	"fmt"
	"github.com/khchehab/muzayaf/random"
)

// HSLAColor represents a color in the HSLA color space
type HSLAColor struct {
	Hue        int
	Saturation int
	Lightness  int
	Alpha      float64
}

// String returns a string representation of the HSL color
func (h HSLAColor) String() string {
	if h.Alpha == 1.0 {
		return fmt.Sprintf("hsl(%d°, %d%%, %d%%)",
			h.Hue, h.Saturation, h.Lightness)
	}
	return fmt.Sprintf("hsla(%d°, %d%%, %d%%, %.2f)",
		h.Hue, h.Saturation, h.Lightness, h.Alpha)
}

// HSLA generates a random HSLA color
func HSLA() HSLAColor {
	return HSLAColor{
		Hue:        random.IntN(361),                  // 0-360 degrees
		Saturation: random.IntN(101),                  // 0-100 percent
		Lightness:  random.IntN(101),                  // 0-100 percent
		Alpha:      float64(random.IntN(101)) / 100.0, // 0-1 inclusive
	}
}

// HSL generates a random HSL color
func HSL() HSLAColor {
	return HSLAColor{
		Hue:        random.IntN(361), // 0-360 degrees
		Saturation: random.IntN(101), // 0-100 percent
		Lightness:  random.IntN(101), // 0-100 percent
		Alpha:      1.0,
	}
}
