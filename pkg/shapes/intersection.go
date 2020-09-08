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

// Intersection Structure to hold intersection information
type Intersection struct {
	Distance float64
	Obj Shape
}

// Hit Get the closest hit from intersections
func Hit(inters []Intersection) (Intersection, bool) {
	var intersection Intersection
	hit := false
	for _, v := range inters {
		if v.Distance >= 0  && (v.Distance < intersection.Distance || !hit) {
			intersection = v
			hit = true
		}
	}
	return intersection, hit
}