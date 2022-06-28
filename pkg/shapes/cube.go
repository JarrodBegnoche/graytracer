package shapes

import (
	"math"

	"github.com/factorion/graytracer/pkg/primitives"
)

// Cube Basic cube representation
type Cube struct {
	ShapeBase
}

// CheckAxis Checks two sides of a cube on an axis from a ray's path on that axis
func CheckAxis(origin, direction, minimum, maximum float64) (float64, float64) {
	// We can do just this because go handles division by zero
	tmin := (minimum - origin) / direction
	tmax := (maximum - origin) / direction
	if tmin > tmax {
		tmin, tmax = tmax, tmin
	}
	return tmin, tmax
}

// MakeCube Make a regular cube with an identity matrix for transform
func MakeCube() *Cube {
	return &Cube{MakeShapeBase()}
}

// GetBounds Return an axis aligned bounding box for the sphere
func (c *Cube) GetBounds() *Bounds {
	bounds := Bounds{Min: primitives.MakePoint(-1, -1, -1), Max: primitives.MakePoint(1, 1, 1)}
	return bounds.Transform(c.transform)
}

// Intersect Check for intersection along one of the six sides of the cube
func (c *Cube) Intersect(r primitives.Ray) Intersections {
	// convert ray to object space
	oray := r.Transform(c.Inverse())
	xtmin, xtmax := CheckAxis(oray.Origin.X, oray.Direction.X, -1, 1)
	ytmin, ytmax := CheckAxis(oray.Origin.Y, oray.Direction.Y, -1, 1)
	ztmin, ztmax := CheckAxis(oray.Origin.Z, oray.Direction.Z, -1, 1)
	tmin := math.Max(math.Max(xtmin, ytmin), ztmin)
	tmax := math.Min(math.Min(xtmax, ytmax), ztmax)
	if tmin > tmax {
		return Intersections{}
	}
	return Intersections{Intersection{Distance: tmin, Obj: c}, Intersection{Distance: tmax, Obj: c}}
}

// Normal Calculate the normal at a given point on the cube
func (c *Cube) Normal(worldPoint primitives.PV, u, v float64) primitives.PV {
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
