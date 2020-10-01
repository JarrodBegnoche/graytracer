package patterns

import(
	"github.com/factorion/graytracer/pkg/primitives"
)

// PatternBase Base class for Pattern objects
type PatternBase struct {
	transform primitives.Matrix
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
