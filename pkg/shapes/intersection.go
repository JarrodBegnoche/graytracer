package shapes

import (
	"math"
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

// Intersections Map of intersections keyed to their distance
type Intersections map[float64]Intersection

// Intersection Structure to hold intersection information
type Intersection struct {
	Distance float64
	Obj Shape
}

// Hit Get the closest hit from intersections
func Hit(inters Intersections) float64 {
	intersection := float64(-1)
	for k := range inters {
		if k >= 0  && (k < intersection || intersection == -1) {
			intersection = k
		}
	}
	return intersection
}