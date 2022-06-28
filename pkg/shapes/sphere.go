package shapes

import (
	"math"

	"github.com/factorion/graytracer/pkg/primitives"
)

// Sphere Represents a sphere of radius 1
type Sphere struct {
	ShapeBase
}

// MakeSphere Make a regular sphere with an identity matrix for transform
func MakeSphere() *Sphere {
	return &Sphere{MakeShapeBase()}
}

// GetBounds Return an axis aligned bounding box for the sphere
func (s *Sphere) GetBounds() *Bounds {
	bounds := &Bounds{Min: primitives.MakePoint(-1, -1, -1), Max: primitives.MakePoint(1, 1, 1)}
	return bounds.Transform(s.transform)
}

// Intersect Check if a ray intersects
func (s *Sphere) Intersect(r primitives.Ray) Intersections {
	hits := Intersections{}
	// convert ray to object space
	oray := r.Transform(s.Inverse())
	// Vector from the sphere's center
	sray := oray.Origin.Subtract(primitives.MakePoint(0, 0, 0))
	a := oray.Direction.DotProduct(oray.Direction)
	b := 2 * oray.Direction.DotProduct(sray)
	c := sray.DotProduct(sray) - 1
	discriminant := (b * b) - (4 * a * c)
	if discriminant < 0 {
		return hits
	}
	hits = append(hits, Intersection{Distance: ((-b - math.Sqrt(discriminant)) / (2 * a)), Obj: s})
	if discriminant > 0 {
		hits = append(hits, Intersection{Distance: ((-b + math.Sqrt(discriminant)) / (2 * a)), Obj: s})
	}
	return hits
}

// Normal Calculate the normal at a given point on the sphere
func (s *Sphere) Normal(worldPoint primitives.PV, u, v float64) primitives.PV {
	objectPoint := s.WorldToObjectPV(worldPoint)
	objectNormal := objectPoint.Subtract(primitives.MakePoint(0, 0, 0))
	worldNormal := s.ObjectToWorldPV(objectNormal)
	worldNormal.W = 0.0
	return worldNormal.Normalize()
}

// UVMapping Return the 2D coordinates of an intersection point
func (s *Sphere) UVMapping(point primitives.PV) primitives.PV {
	objectPoint := s.WorldToObjectPV(point)
	d := primitives.MakePoint(0, 0, 0).Subtract(objectPoint)
	return primitives.MakePoint(0.5+math.Atan2(d.X, d.Z)/(2*math.Pi), 0.5-math.Asin(d.Y)/math.Pi, 0)
}
