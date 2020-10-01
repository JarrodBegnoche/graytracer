package patterns

import (
	"math"
	"github.com/factorion/graytracer/pkg/primitives"
)

// Gradient Basic gradient pattern based on x-axis value
type Gradient struct {
	PatternBase
	pattern1, pattern2 Pattern
}

// MakeGradient Return a basic gradient pattern with two colors
func MakeGradient(p1, p2 Pattern) *Gradient {
	return &Gradient{PatternBase:PatternBase{transform:primitives.MakeIdentityMatrix(4)}, pattern1:p1, pattern2:p2}
}

// ColorAt Return color at specific point
func (g Gradient) ColorAt(point primitives.PV) RGB {
	patternPoint := g.PatternPoint(point)
	pX := patternPoint.X
	color1 := g.pattern1.ColorAt(patternPoint)
	color2 := g.pattern2.ColorAt(patternPoint)
	return color1.Add(color2.Subtract(color1).Scale(pX - math.Floor(pX)))
}
