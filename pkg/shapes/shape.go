package shapes

import (
	"github.com/factorion/graytracer/pkg/primitives"
)

// Shape Interface for different 3D and 2D shape modules
type Shape interface {
	Intersect(r primitives.Ray) []float64
	//SetTransform(m primitives.Matrix)
	Transform() primitives.Matrix
}
