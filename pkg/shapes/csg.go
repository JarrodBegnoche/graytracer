package shapes

import (
	"sort"

	"github.com/factorion/graytracer/pkg/primitives"
)

type Operation int

const (
	UNION Operation = iota
	INTERSECT
	DIFFERENCE
)

// CSG Constructive Solid Geometry node for two separate shapes
type CSG struct {
	ShapeBase
	op          Operation
	left, right Shape
	bounds      *Bounds
}

// MakeCSG
func MakeCSG(op Operation, shape1, shape2 Shape) *CSG {
	return &CSG{MakeShapeBase(), op, shape1, shape2, nil}
}

// GetBounds Return an axis aligned bounding box for the CSG
func (csg *CSG) GetBounds() *Bounds {
	return csg.bounds
}

// intersection_allowed Determines whether the intersection is valid or not
func (csg *CSG) IntersectionAllowed(lhit, inl, inr bool) bool {
	allowed := false
	switch csg.op {
	case UNION:
		allowed = ((lhit && !inr) || (!lhit && !inl))
	case INTERSECT:
		allowed = ((lhit && inr) || (!lhit && inl))
	case DIFFERENCE:
		allowed = ((lhit && !inr) || (!lhit && inl))
	}
	return allowed
}

// Intersect Check if a ray intersects
func (csg *CSG) Intersect(r primitives.Ray) Intersections {
	hits := Intersections{}
	if (csg.bounds == nil) || (csg.bounds.Intersect(r)) {
		// convert ray to object space
		oray := r.Transform(csg.Inverse())
		lhits := csg.left.Intersect(oray)
		rhits := csg.right.Intersect(oray)
		sort.Sort(lhits)
		sort.Sort(rhits)
		lhit := false
		inl := false
		inr := false
		lindex := 0
		rindex := 0
		for lindex < lhits.Len() || rindex < rhits.Len() {
			var i Intersection
			if (lindex >= lhits.Len()) || (rindex < rhits.Len() && rhits[rindex].Distance < lhits[lindex].Distance) {
				i = rhits[rindex]
				lhit = false
				rindex++
			} else {
				i = lhits[lindex]
				lhit = true
				lindex++
			}
			if csg.IntersectionAllowed(lhit, inl, inr) {
				hits = append(hits, i)
			}
			if lhit {
				inl = !inl
			} else {
				inr = !inr
			}
		}
	}
	return hits
}

// Normal Calculate the normal at a given point on the sphere
func (csg *CSG) Normal(worldPoint primitives.PV, u, v float64) primitives.PV {
	// Only exists for Interface, should never be called
	return primitives.MakeVector(0, 1, 0)
}

// UVMapping Return the 2D coordinates of an intersection point
func (csg *CSG) UVMapping(point primitives.PV) primitives.PV {
	// Only exists for Interface, should never be called
	return primitives.MakePoint(point.X, point.Y, 0)
}
