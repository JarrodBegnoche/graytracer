package shapes

import (
	"math"
	"github.com/factorion/graytracer/pkg/primitives"
)

// SliceEquals Check if two slices are equal
func SliceEquals(a, b []float64) bool {
	if len(a) != len(b) {
        return false
	}
	EPSILON := 0.00000001
    for i, v := range a {
        if math.Abs(v - b[i]) > EPSILON {
            return false
        }
    }
    return true
}

// Shape Interface for different 3D and 2D shape modules
type Shape interface {
	Intersect(r primitives.Ray) []float64
	Normal(worldPoint primitives.PV) primitives.PV
	SetTransform(m primitives.Matrix)
	Transform() primitives.Matrix
	SetMaterial(mat primitives.Material)
	Material() primitives.Material
}
