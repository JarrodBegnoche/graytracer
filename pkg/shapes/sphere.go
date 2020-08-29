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
func (s Sphere) Intersect(r primitives.Ray) (bool, float64, float64) {
	// Vector from the sphere's center
	sray := r.Origin.Subtract(primitives.MakePoint(0, 0, 0))
	a := r.Direction.DotProduct(r.Direction)
	b := 2 * r.Direction.DotProduct(sray)
	c := sray.DotProduct(sray) - 1
	discriminant := (b * b) - (4 * a * c)
	if discriminant < 0 {
		return false, 0, 0
	}
	return true, (-b - math.Sqrt(discriminant)) / (2 * a), (-b + math.Sqrt(discriminant)) / (2 * a)
}