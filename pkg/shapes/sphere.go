package shapes

import (
	"math"
	"github.com/factorion/graytracer/pkg/primitives"
)

// Sphere Represents a sphere
type Sphere struct {
	center primitives.PV
	radius float64
}

// Intersect Check if a ray intersects
func (s Sphere) Intersect(r primitives.Ray) []float64 {
	hits := []float64{}
	// Vector from the sphere's center
	sray := r.Origin.Subtract(primitives.MakePoint(0, 0, 0))
	a := r.Direction.DotProduct(r.Direction)
	b := 2 * r.Direction.DotProduct(sray)
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