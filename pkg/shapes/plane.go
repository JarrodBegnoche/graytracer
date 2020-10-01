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

// Intersect Check if a ray intersects
func (p *Plane) Intersect(r primitives.Ray) []float64 {
	hits := []float64{}
	// convert ray to object space
	inverse, _ := p.transform.Inverse()
	objectRay := r.Transform(inverse)
	if math.Abs(objectRay.Direction.Y) > primitives.EPSILON {
		hits = append(hits, -objectRay.Origin.Y / objectRay.Direction.Y)
	}
	return hits
}

// Normal Calculate the normal at a given point on the sphere
func (p *Plane) Normal(worldPoint primitives.PV) primitives.PV {
	inverse, _ := p.transform.Inverse()
	objectNormal := primitives.MakeVector(0, 1, 0)
	worldNormal := objectNormal.Transform(inverse.Transpose())
	worldNormal.W = 0.0
	return worldNormal.Normalize()
}

// UVMapping Return the 2D coordinates of an intersected point
func (p *Plane) UVMapping(point primitives.PV) primitives.PV {
	inverse, _ := p.transform.Inverse()
	objectPoint := point.Transform(inverse)
	return primitives.MakePoint(objectPoint.X, objectPoint.Z, 0)
}
