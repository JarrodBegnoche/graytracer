package shapes

import (
	"math"
	"github.com/factorion/graytracer/pkg/primitives"
)

// Cube Basic cube representation
type Cube struct {
	ShapeBase
}

func checkAxis(origin, direction float64) (float64, float64) {
	// We can do just this because go handles division by zero
	tmin := (-1 - origin) / direction
	tmax := (1 - origin) / direction
	if tmin > tmax {
		tmin, tmax = tmax, tmin
	}
	return tmin, tmax
}

// MakeCube Make a regular cube with an identity matrix for transform
func MakeCube() *Cube {
	return &Cube{MakeShapeBase()}
}

// Intersect Check for intersection along one of the six sides of the cube
func (c *Cube) Intersect(r primitives.Ray) []float64 {
	// convert ray to object space
	ray2 := r.Transform(c.Inverse())
	xtmin, xtmax := checkAxis(ray2.Origin.X, ray2.Direction.X)
	ytmin, ytmax := checkAxis(ray2.Origin.Y, ray2.Direction.Y)
	ztmin, ztmax := checkAxis(ray2.Origin.Z, ray2.Direction.Z)
	tmin := math.Max(math.Max(xtmin, ytmin), ztmin)
	tmax := math.Min(math.Min(xtmax, ytmax), ztmax)
	if tmin > tmax {
		return []float64{}
	}
	return []float64{tmin, tmax}
}

// Normal Calculate the normal at a given point on the cube
func (c *Cube) Normal(worldPoint primitives.PV) primitives.PV {
	objectPoint := worldPoint.Transform(c.Inverse())
	absx := math.Abs(objectPoint.X)
	absy := math.Abs(objectPoint.Y)
	absz := math.Abs(objectPoint.Z)
	max := math.Max(absx, math.Max(absy, absz))
	var objectNormal primitives.PV
	if max == absx {
		objectNormal = primitives.MakeVector(objectPoint.X, 0, 0)
	} else if max == absy {
		objectNormal = primitives.MakeVector(0, objectPoint.Y, 0)
	} else {
		objectNormal = primitives.MakeVector(0, 0, objectPoint.Z)
	}
	worldNormal := objectNormal.Transform(c.Inverse().Transpose())
	worldNormal.W = 0.0
	return worldNormal.Normalize()
}

// UVMapping Return the 2D coordinates of an intersection point
func (c *Cube) UVMapping(point primitives.PV) primitives.PV {
	return primitives.MakePoint(point.X, point.Y, 0)
}
