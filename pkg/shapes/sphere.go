package shapes

import (
	"math"
	"github.com/factorion/graytracer/pkg/primitives"
)

// Sphere Represents a sphere
type Sphere struct {
	center primitives.PV
	radius float64
	transform primitives.Matrix
}

// MakeSphere Make a regular sphere with an identity matrix for transform
func MakeSphere(x, y, z, radius float64) Sphere {
	return Sphere{center:primitives.MakePoint(x, y, z), radius:radius, transform:primitives.MakeIdentityMatrix(4)}
}

// MakeTransformedSphere Make a sphere with a custom transform matrix
func MakeTransformedSphere(x, y, z, radius float64, m primitives.Matrix) Sphere {
	return Sphere{center:primitives.MakePoint(x, y, z), radius:radius, transform:m}
}

// SetTransform Set the transform matrix
func (s Sphere) SetTransform(m primitives.Matrix) {
	s.transform = m
}

// Transform Get the transform matrix
func (s Sphere) Transform() primitives.Matrix {
	return s.transform
}

// Intersect Check if a ray intersects
func (s Sphere) Intersect(r primitives.Ray) []float64 {
	hits := []float64{}
	// convert ray to object space
	inverse, _ := s.transform.Inverse()
	ray2 := r.Transform(inverse)
	// Vector from the sphere's center
	sray := ray2.Origin.Subtract(primitives.MakePoint(0, 0, 0))
	a := ray2.Direction.DotProduct(ray2.Direction)
	b := 2 * ray2.Direction.DotProduct(sray)
	c := sray.DotProduct(sray) - 1
	discriminant := (b * b) - (4 * a * c)
	if discriminant < 0 {
		return hits
	}
	hits = append(hits, (-b - math.Sqrt(discriminant)) / (2 * a))
	if discriminant > 0 {
		hits = append(hits, (-b + math.Sqrt(discriminant)) / (2 * a))
	}
	return hits
}