package shapes

import (
	"math"
	"github.com/factorion/graytracer/pkg/primitives"
)

// Cylinder Represents a cylinder
type Cylinder struct {
	ShapeBase
	closed bool
}

// MakeCylinder Make a regular cylinder with an identity matrix for transform
func MakeCylinder(closed bool) *Cylinder {
	return &Cylinder{MakeShapeBase(), closed}
}

// checkCap Checks if an intersection happens at the cap
func (cyl Cylinder) checkCap(r primitives.Ray, t float64) bool {
	x := r.Origin.X + (t * r.Direction.X)
	z := r.Origin.Z + (t * r.Direction.Z)
	return ((x * x) + (z * z)) <= 1
}

// Intersect Check if a ray intersects
func (cyl *Cylinder) Intersect(r primitives.Ray) []float64 {
	hits := []float64{}
	// convert ray to object space
	inverse, _ := cyl.transform.Inverse()
	oray := r.Transform(inverse)
	direction := oray.Direction.Normalize()
	// Vector from the Cylinder's center
	//sray := oray.Origin.Subtract(primitives.MakePoint(0, 0, 0))
	a := (direction.X * direction.X) + (direction.Z * direction.Z)
	if math.Abs(a) > primitives.EPSILON {
		b := (2.0 * oray.Origin.X * direction.X) + (2.0 * oray.Origin.Z * direction.Z)
		c := (oray.Origin.X * oray.Origin.X) + (oray.Origin.Z * oray.Origin.Z) - 1.0
		discriminant := (b * b) - (4.0 * a * c)
		if discriminant < 0 {
			return hits
		}
		t0 := (-b - math.Sqrt(discriminant)) / (2.0 * a)
		t1 := (-b + math.Sqrt(discriminant)) / (2.0 * a)

		if t0 > t1 {
			t0, t1 = t1, t0
		}

		// Verify hits are within height of cone
		y0 := oray.Origin.Y + (t0 * oray.Direction.Y)
		if (0 < y0) && (y0 < 1) {
			hits = append(hits, t0)
		}

		y1 := oray.Origin.Y + (t1 * oray.Direction.Y)
		if (0 < y1) && (y1 < 1) {
			hits = append(hits, t1)
		}
	}	

	// Cap checking only matters if cylinder is closed
	if !cyl.closed || math.Abs(oray.Direction.Y) < primitives.EPSILON {
		return hits
	}

	// Check bottom and top caps
	t := -oray.Origin.Y / oray.Direction.Y
	if cyl.checkCap(oray, t) {
		hits = append(hits, t)
	}

	t = (1.0 / oray.Direction.Y) + t
	if cyl.checkCap(oray, t) {
		hits = append(hits, t)
	}

	return hits
}

// Normal Calculate the normal at a given point on the Cylinder
func (cyl *Cylinder) Normal(worldPoint primitives.PV) primitives.PV {
	var objectNormal primitives.PV
	inverse, _ := cyl.transform.Inverse()
	objectPoint := worldPoint.Transform(inverse)
	distance := (objectPoint.X * objectPoint.X) + (objectPoint.Z * objectPoint.Z)
	if (distance < 1) && (objectPoint.Y >= (1.0 - primitives.EPSILON)) {
		objectNormal = primitives.MakeVector(0, 1, 0)
	} else if (distance < 1) && (objectPoint.Y <= primitives.EPSILON) {
		objectNormal = primitives.MakeVector(0, -1, 0)
	} else {
		objectNormal = primitives.MakeVector(objectPoint.X, 0, objectPoint.Z)
	}	
	worldNormal := objectNormal.Transform(inverse.Transpose())
	worldNormal.W = 0.0
	return worldNormal.Normalize()
}

// UVMapping Return the 2D coordinates of an intersection point
func (cyl *Cylinder) UVMapping(point primitives.PV) primitives.PV {
	inverse, _ := cyl.transform.Inverse()
	objectPoint := point.Transform(inverse)
	d := primitives.MakePoint(0, 0, 0).Subtract(objectPoint)
	return primitives.MakePoint(0.5 + math.Atan2(d.X, d.Z) / (2 * math.Pi), objectPoint.Y, 0)
}
