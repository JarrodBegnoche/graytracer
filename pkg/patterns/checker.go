package patterns

import (
	"math"
	"github.com/factorion/graytracer/pkg/primitives"
)

// Checker Basic checker pattern based on x-axis value
type Checker struct {
	PatternBase
	pattern1, pattern2 Pattern
}

// MakeChecker Return a basic Checker pattern with two colors
func MakeChecker(p1, p2 Pattern) *Checker {
	return &Checker{PatternBase:PatternBase{transform:primitives.MakeIdentityMatrix(4)}, pattern1:p1, pattern2:p2}
}

// ColorAt Return color at specific point
func (c Checker) ColorAt(point primitives.PV) RGB {
	patternPoint := c.PatternPoint(point)
	if (int(math.Floor(patternPoint.X)) + int(math.Floor(patternPoint.Y))) % 2 == 0 {
		return c.pattern1.ColorAt(patternPoint)
	}
	return c.pattern2.ColorAt(patternPoint)
}
