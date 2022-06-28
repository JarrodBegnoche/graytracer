package shapes

import (
	"math"

	"github.com/factorion/graytracer/pkg/primitives"
)

// Plane Plane along the XZ axis
type Plane struct {
	ShapeBase
}

// MakePlane Make a default plane
func MakePlane() *Plane {
	return &Plane{MakeShapeBase()}
}

// GetBounds Return an axis aligned bounding box for the sphere
func (p *Plane) GetBounds() *Bounds {
	bounds := &Bounds{Min: primitives.MakePoint(math.Inf(-1), -primitives.EPSILON, math.Inf(-1)),
		Max: primitives.MakePoint(math.Inf(1), primitives.EPSILON, math.Inf(1))}
	return bounds.Transform(p.transform)
}

// Intersect Check if a ray intersects
func (p *Plane) Intersect(r primitives.Ray) Intersections {
	hits := Intersections{}
	// convert ray to object space
	objectRay := r.Transform(p.Inverse())
	if math.Abs(objectRay.Direction.Y) > primitives.EPSILON {
		hits = append(hits, Intersection{Distance: (-objectRay.Origin.Y / objectRay.Direction.Y), Obj: p})
	}
	return hits
}

// Normal Calculate the normal at a given point on the sphere
func (p *Plane) Normal(worldPoint primitives.PV, u, v float64) primitives.PV {
	objectNormal := primitives.MakeVector(0, 1, 0)
	worldNormal := p.ObjectToWorldPV(objectNormal)
	worldNormal.W = 0.0
	return worldNormal.Normalize()
}

// UVMapping Return the 2D coordinates of an intersected point
func (p *Plane) UVMapping(point primitives.PV) primitives.PV {
	objectPoint := p.WorldToObjectPV(point)
	return primitives.MakePoint(objectPoint.X, objectPoint.Z, 0)
}
