package components

import (
	"github.com/factorion/graytracer/pkg/primitives"
	"github.com/factorion/graytracer/pkg/shapes"
)

// Computations Set of pre-computed values used for point detection
type Computations struct {
	Intersection
	Point, EyeVector, NormalVector primitives.PV
	Inside bool
}

// Intersection Structure to hold intersection information
type Intersection struct {
	Distance float64
	Obj shapes.Shape
}

// Intersections Sortable list of intersection structs
type Intersections []Intersection

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

// PrepareComputations Calculates the vectors at the point on the object
func (i Intersection) PrepareComputations(ray primitives.Ray) Computations {
	comp := Computations{Intersection:i}
	comp.Point = ray.Position(comp.Distance)
	comp.EyeVector = ray.Direction.Negate()
	comp.NormalVector = comp.Obj.Normal(comp.Point)
	if comp.NormalVector.DotProduct(comp.EyeVector) < 0 {
		comp.NormalVector = comp.NormalVector.Negate()
		comp.Inside = true
	} else{
		comp.Inside = false
	}
	return comp
}
