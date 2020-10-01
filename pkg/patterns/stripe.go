package patterns

import (
	"math"
	"github.com/factorion/graytracer/pkg/primitives"
)

// Stripe Basic stripe pattern based on x-axis value
type Stripe struct {
	PatternBase
	pattern1, pattern2 Pattern
}

// MakeStripe Make a stripe pattern from two patterns
func MakeStripe(p1, p2 Pattern) *Stripe {
	return &Stripe{PatternBase:PatternBase{transform:primitives.MakeIdentityMatrix(4)}, pattern1:p1, pattern2:p2}
}

// ColorAt Calculate which stripe and return the color
func (s Stripe) ColorAt(point primitives.PV) RGB {
	patternPoint := s.PatternPoint(point)
	if int(math.Floor(patternPoint.X)) % 2 == 0 {
		return s.pattern1.ColorAt(patternPoint)
	}
	return s.pattern2.ColorAt(patternPoint)
}
