package color

import (
	"fmt"
	"github.com/khchehab/muzayaf/random"
)

// CMYKColor represents a color in the CMYK color space
type CMYKColor struct {
	Cyan    int
	Magenta int
	Yellow  int
	Key     int
}

// String returns a string representation of the CMYK color
func (c CMYKColor) String() string {
	return fmt.Sprintf("cmyk(%d%%, %d%%, %d%%, %d%%)",
		c.Cyan, c.Magenta, c.Yellow, c.Key)
}

// CMYK generates a random CMYK color
func CMYK() CMYKColor {
	return CMYKColor{
		Cyan:    random.IntN(101), // 0-100 percent
		Magenta: random.IntN(101), // 0-100 percent
		Yellow:  random.IntN(101), // 0-100 percent
		Key:     random.IntN(101), // 0-100 percent
	}
}
