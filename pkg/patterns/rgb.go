package patterns

import (
	"image/color"
	"math"
	"github.com/factorion/graytracer/pkg/primitives"
)

// RGB represents red, green, and blue values for a color object
type RGB struct {
	PatternBase
	red, green, blue float64
}

// MakeRGB Factory method for RGB object
func MakeRGB(red, green, blue float64) *RGB {
	return &RGB{PatternBase:PatternBase{transform:primitives.MakeIdentityMatrix(4)}, red:red, green:green, blue:blue}
}

// Equals Compares two RGB color objects with an amount for approximation
func (r RGB) Equals(g RGB) bool {
	if math.Abs(r.red - g.red) > primitives.EPSILON {
		return false
	}
	if math.Abs(r.green - g.green) > primitives.EPSILON {
		return false
	}
	if math.Abs(r.blue - g.blue) > primitives.EPSILON {
		return false
	}
	return true
}

// Add Adds one RGB color to another and returns as a new RGB object
func (r RGB) Add(g RGB) RGB {
	return RGB{PatternBase{transform:primitives.MakeIdentityMatrix(4)}, r.red + g.red, r.green + g.green, r.blue + g.blue}
}

// Subtract Subtracts one RGB color from another and returns as a new RGB object
func (r RGB) Subtract(g RGB) RGB {
	return RGB{PatternBase{transform:primitives.MakeIdentityMatrix(4)}, r.red - g.red, r.green - g.green, r.blue - g.blue}
}

// Multiply Multiples one RGB color to another and returns as a new RGB object
func (r RGB) Multiply(g RGB) RGB {
	return RGB{PatternBase{transform:primitives.MakeIdentityMatrix(4)}, r.red * g.red, r.green * g.green, r.blue * g.blue}
}

// Scale Scale an RGB color by a single value and return as a new RGB object
func (r RGB) Scale(s float64) RGB {
	return RGB{PatternBase{transform:primitives.MakeIdentityMatrix(4)}, r.red * s, r.green * s, r.blue * s}
}

// ToImageRGBA Convert to an RGBA image format
func (r RGB) ToImageRGBA() color.RGBA {
	return color.RGBA{byte(math.Min(1.0, math.Max(0, r.red)) * 255),
					  byte(math.Min(1.0, math.Max(0, r.green)) * 255),
					  byte(math.Min(1.0, math.Max(0, r.blue)) * 255), 0xff}
}

// ColorAt Pattern interface function
func (r RGB) ColorAt(point primitives.PV) RGB {
	return r
}
