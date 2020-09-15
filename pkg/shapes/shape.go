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
    for i, v := range a {
        if math.Abs(v - b[i]) > primitives.EPSILON {
            return false
        }
    }
    return true
}

// ShapeBase Base struct to be embedded in shape objects
type ShapeBase struct {
	transform primitives.Matrix
	material primitives.Material
}

// MakeShapeBase Make a regular sphere with an identity matrix for transform
func MakeShapeBase() ShapeBase {
	return ShapeBase{transform:primitives.MakeIdentityMatrix(4), material:primitives.MakeDefaultMaterial()}
}

// SetTransform Set the transform matrix
func (s *ShapeBase) SetTransform(m primitives.Matrix) {
	s.transform = m
}

// Transform Get the transform matrix
func (s *ShapeBase) Transform() primitives.Matrix {
	return s.transform
}

// SetMaterial Set the material for the sphere
func (s *ShapeBase) SetMaterial(mat primitives.Material) {
	s.material = mat
}

// Material Get the material object
func (s *ShapeBase) Material() primitives.Material {
	return s.material
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
