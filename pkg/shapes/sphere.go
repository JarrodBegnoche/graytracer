package shapes

import (
	"math"
	"github.com/factorion/graytracer/pkg/primitives"
)

// Sphere Represents a sphere
type Sphere struct {
	ShapeBase
}

// MakeSphere Make a regular sphere with an identity matrix for transform
func MakeSphere() *Sphere {
	return &Sphere{MakeShapeBase()}
}

// Intersect Check if a ray intersects
func (s *Sphere) Intersect(r primitives.Ray) []float64 {
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

// Normal Calculate the normal at a given point on the sphere
func (s *Sphere) Normal(worldPoint primitives.PV) primitives.PV {
	inverse, _ := s.transform.Inverse()
	objectPoint := worldPoint.Transform(inverse)
	objectNormal := objectPoint.Subtract(primitives.MakePoint(0, 0, 0))
	worldNormal := objectNormal.Transform(inverse.Transpose())
	worldNormal.W = 0.0
	return worldNormal.Normalize()
}

// UVMapping Return the 2D coordinates of an intersection point
func (s *Sphere) UVMapping(point primitives.PV) primitives.PV {
	inverse, _ := s.transform.Inverse()
	d := primitives.MakePoint(0, 0, 0).Subtract(point.Transform(inverse))
	return primitives.MakePoint(0.5 + math.Atan2(d.X, d.Z) / (2 * math.Pi), 0.5 - math.Asin(d.Y) / math.Pi, 0)
}
