package patterns

import(
	"github.com/factorion/graytracer/pkg/primitives"
)

// PatternBase Base class for Pattern objects
type PatternBase struct {
	transform primitives.Matrix
}

// MakePatternBase Make an empty PatternBase Object
func MakePatternBase() PatternBase {
	return PatternBase{transform:primitives.MakeIdentityMatrix(4)}
}

// SetTransform Parent class for pattern interface
func (pb *PatternBase) SetTransform(transform primitives.Matrix) {
	pb.transform = transform
}

// PatternPoint Convert a point to pattern space
func (pb PatternBase) PatternPoint(point primitives.PV) primitives.PV {
	inverse, _ := pb.transform.Inverse()
	return point.Transform(inverse)
}

// Pattern Patterns are represented with a GetColorAt function
type Pattern interface {
	ColorAt(primitives.PV) RGB
	SetTransform(primitives.Matrix)
}

// TestPattern Basic pattern used for testing
type TestPattern struct {
	PatternBase
}

// ColorAt Return the points as the color
func (tp TestPattern) ColorAt(point primitives.PV) RGB {
	patternPoint := tp.PatternPoint(point)
	rgb := MakeRGB(patternPoint.X, patternPoint.Y, patternPoint.Z)
	return *rgb
}
