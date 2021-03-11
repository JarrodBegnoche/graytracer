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
func (c *Cube) Intersect(r primitives.Ray) Intersections {
	// convert ray to object space
	oray := r.Transform(c.Inverse())
	xtmin, xtmax := checkAxis(oray.Origin.X, oray.Direction.X)
	ytmin, ytmax := checkAxis(oray.Origin.Y, oray.Direction.Y)
	ztmin, ztmax := checkAxis(oray.Origin.Z, oray.Direction.Z)
	tmin := math.Max(math.Max(xtmin, ytmin), ztmin)
	tmax := math.Min(math.Min(xtmax, ytmax), ztmax)
	if tmin > tmax {
		return Intersections{}
	}
	return Intersections{Intersection{Distance:tmin, Obj:c}, Intersection{Distance:tmax, Obj:c}}
}

// Normal Calculate the normal at a given point on the cube
func (c *Cube) Normal(worldPoint primitives.PV) primitives.PV {
	objectPoint := c.WorldToObjectPV(worldPoint)
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
	worldNormal := c.ObjectToWorldPV(objectNormal)
	worldNormal.W = 0.0
	return worldNormal.Normalize()
}

// UVMapping Return the 2D coordinates of an intersection point
func (c *Cube) UVMapping(point primitives.PV) primitives.PV {
	return primitives.MakePoint(point.X, point.Y, 0)
}
