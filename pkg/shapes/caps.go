package shapes

import (
	"github.com/factorion/graytracer/pkg/primitives"
)

// CheckCap Checks if an intersection happens at the cap
func CheckCap(r primitives.Ray, t float64) bool {
	x := r.Origin.X + (t * r.Direction.X)
	z := r.Origin.Z + (t * r.Direction.Z)
	return ((x * x) + (z * z)) <= 1
}
