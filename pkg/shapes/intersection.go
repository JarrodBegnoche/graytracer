package shapes

import (
	"math"
	"sort"
	"github.com/factorion/graytracer/pkg/primitives"
)

// Intersection Structure to hold intersection information
type Intersection struct {
	Distance float64
	Obj Shape
}

// Intersections Sortable list of intersection structs
type Intersections []Intersection

// IntersectEquals Check if two slices are equal
func IntersectEquals(xs Intersections, dists []float64) bool {
	if len(xs) != len(dists) {
        return false
	}
	sort.Sort(xs)
	sort.Float64s(dists)
    for i := range xs {
        if math.Abs(xs[i].Distance - dists[i]) > primitives.EPSILON {
            return false
        }
    }
    return true
}

// Necessary functions to make Intersections sortable
func (i Intersections) Len() int { return len(i) }

func (i Intersections) Less(j, k int) bool { return i[j].Distance < i[k].Distance }

func (i Intersections) Swap(j, k int) { i[j], i[k] = i[k], i[j] }

// Hit Get the closest hit from intersections, assumes i is sorted
func (i Intersections) Hit() (Intersection, bool) {
	for _, v := range i {
		if v.Distance >= 0 {
			return v, true
		}
	}
	return Intersection{}, false
}
