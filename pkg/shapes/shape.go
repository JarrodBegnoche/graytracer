package shapes

import (
	"github.com/factorion/graytracer/pkg/primitives"
	"github.com/factorion/graytracer/pkg/patterns"
)

// ShapeBase Base struct to be embedded in shape objects
type ShapeBase struct {
	transform primitives.Matrix
	inverse primitives.Matrix
	material patterns.Material
	parent Shape
}

// MakeShapeBase Make a regular sphere with an identity matrix for transform
func MakeShapeBase() ShapeBase {
	return ShapeBase{transform:primitives.MakeIdentityMatrix(4),
					 inverse:primitives.MakeIdentityMatrix(4),
					 material:patterns.MakeDefaultMaterial(),
					 parent:nil}
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

// SetMaterial Set the material for the shape
func (s *ShapeBase) SetMaterial(mat patterns.Material) {
	s.material = mat
}

// Material Get the material object
func (s *ShapeBase) Material() patterns.Material {
	return s.material
}

// SetParent Set the parent object of the shape
func (s *ShapeBase) SetParent(parent Shape) {
	s.parent = parent
}

// Parent Get the parent object of the shape
func (s *ShapeBase) Parent() Shape {
	return s.parent
}

// WorldToObjectPV Convert a Point/Vector from world to object-space
func (s *ShapeBase) WorldToObjectPV(pv primitives.PV) primitives.PV {
	if (s.parent != nil) {
		pv = s.parent.WorldToObjectPV(pv)
	}
	return pv.Transform(s.Inverse())
}

// ObjectToWorldPV Convert a Point/Vector from object to world-space
func (s *ShapeBase) ObjectToWorldPV(pv primitives.PV) primitives.PV {
	result := pv.Transform(s.Inverse().Transpose())
	if (s.parent != nil) {
		result = s.parent.ObjectToWorldPV(result)
	}
	return result
}

// Shape Interface for different 3D and 2D shape modules
type Shape interface {
	Intersect(primitives.Ray) Intersections
	Normal(primitives.PV) primitives.PV
	SetTransform(primitives.Matrix)
	Transform() primitives.Matrix
	SetMaterial(patterns.Material)
	Material() patterns.Material
	SetParent(Shape)
	Parent() Shape
	GetBounds() *Bounds
	UVMapping(primitives.PV) primitives.PV
	WorldToObjectPV(primitives.PV) primitives.PV
	ObjectToWorldPV(primitives.PV) primitives.PV
}
