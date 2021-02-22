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

// Intersect Check if a ray intersects
func (cyl *Cylinder) Intersect(r primitives.Ray) []float64 {
	hits := []float64{}
	// convert ray to object space
	oray := r.Transform(cyl.Inverse())
	//oray.Direction = oray.Direction.Normalize()
	a := (oray.Direction.X * oray.Direction.X) + (oray.Direction.Z * oray.Direction.Z)
	if math.Abs(a) > primitives.EPSILON {
		b := 2.0 * ((oray.Origin.X * oray.Direction.X) + (oray.Origin.Z * oray.Direction.Z))
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
	if !cyl.closed || math.Abs(oray.Direction.Y) < 0 {
		return hits
	}

	// Check bottom and top caps
	t := -oray.Origin.Y / oray.Direction.Y
	if CheckCap(oray, t) {
		hits = append(hits, t)
	}

	t = (1.0 / oray.Direction.Y) + t
	if CheckCap(oray, t) {
		hits = append(hits, t)
	}

	return hits
}

// Normal Calculate the normal at a given point on the Cylinder
func (cyl *Cylinder) Normal(worldPoint primitives.PV) primitives.PV {
	var objectNormal primitives.PV
	objectPoint := worldPoint.Transform(cyl.Inverse())
	distance := (objectPoint.X * objectPoint.X) + (objectPoint.Z * objectPoint.Z)
	if (distance < 1) && (objectPoint.Y >= (1.0 - primitives.EPSILON)) {
		objectNormal = primitives.MakeVector(0, 1, 0)
	} else if (distance < 1) && (objectPoint.Y <= primitives.EPSILON) {
		objectNormal = primitives.MakeVector(0, -1, 0)
	} else {
		objectNormal = primitives.MakeVector(objectPoint.X, 0, objectPoint.Z)
	}	
	worldNormal := objectNormal.Transform(cyl.Inverse().Transpose())
	worldNormal.W = 0.0
	return worldNormal.Normalize()
}

// UVMapping Return the 2D coordinates of an intersection point
func (cyl *Cylinder) UVMapping(point primitives.PV) primitives.PV {
	objectPoint := point.Transform(cyl.Inverse())
	d := primitives.MakePoint(0, 0, 0).Subtract(objectPoint)
	return primitives.MakePoint(0.5 + math.Atan2(d.X, d.Z) / (2 * math.Pi), objectPoint.Y, 0)
}
