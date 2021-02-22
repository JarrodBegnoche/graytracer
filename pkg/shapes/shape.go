package shapes

import (
	"math"
	"github.com/factorion/graytracer/pkg/primitives"
	"github.com/factorion/graytracer/pkg/patterns"
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
	inverse primitives.Matrix
	material patterns.Material
}

// MakeShapeBase Make a regular sphere with an identity matrix for transform
func MakeShapeBase() ShapeBase {
	return ShapeBase{transform:primitives.MakeIdentityMatrix(4),
					 inverse:primitives.MakeIdentityMatrix(4),
					 material:patterns.MakeDefaultMaterial()}
}

// SetTransform Set the transform matrix
func (s *ShapeBase) SetTransform(m primitives.Matrix) {
	inverse, _ := m.Inverse()
	s.transform = m
	s.inverse = inverse
}

// Transform Get the transform matrix
func (s *ShapeBase) Transform() primitives.Matrix {
	return s.transform
}

// Inverse Get the Inverse of the transform matrix
func (s *ShapeBase) Inverse() primitives.Matrix {
	return s.inverse
}

// SetMaterial Set the material for the sphere
func (s *ShapeBase) SetMaterial(mat patterns.Material) {
	s.material = mat
}

// Material Get the material object
func (s *ShapeBase) Material() patterns.Material {
	return s.material
}

// Shape Interface for different 3D and 2D shape modules
type Shape interface {
	Intersect(r primitives.Ray) []float64
	Normal(worldPoint primitives.PV) primitives.PV
	SetTransform(m primitives.Matrix)
	Transform() primitives.Matrix
	SetMaterial(mat patterns.Material)
	Material() patterns.Material
	UVMapping(primitives.PV) primitives.PV
}
