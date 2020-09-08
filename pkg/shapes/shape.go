package shapes

import (
	"github.com/factorion/graytracer/pkg/primitives"
)

// Shape Interface for different 3D and 2D shape modules
type Shape interface {
	Intersect(r primitives.Ray) []float64
	Normal(worldPoint primitives.PV) primitives.PV
	SetTransform(m primitives.Matrix)
	Transform() primitives.Matrix
	SetMaterial(mat primitives.Material)
	Material() primitives.Material
}
