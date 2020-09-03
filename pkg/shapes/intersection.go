package shapes

// Intersections Map of intersections keyed to their distance
type Intersections map[float64]Intersection

// Intersection Structure to hold intersection information
type Intersection struct {
	distance float64
	shape Shape
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