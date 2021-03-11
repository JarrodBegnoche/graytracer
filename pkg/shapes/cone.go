package shapes

import (
	"math"
	"github.com/factorion/graytracer/pkg/primitives"
)

// Cone Represents a cone
type Cone struct {
	ShapeBase
	closed bool
}

// MakeCone Make a regular cone with an identity matrix for transform
func MakeCone(closed bool) *Cone {
	return &Cone{MakeShapeBase(), closed}
}

// Intersect Check if a ray intersects
func (cone *Cone) Intersect(r primitives.Ray) Intersections {
	hits := Intersections{}
	// convert ray to object space
	oray := r.Transform(cone.Inverse())
	a := (oray.Direction.X * oray.Direction.X) - (oray.Direction.Y * oray.Direction.Y) +
		 (oray.Direction.Z * oray.Direction.Z)
	b := (2.0 * oray.Origin.X * oray.Direction.X) - (2.0 * oray.Origin.Y * oray.Direction.Y) +
		 (2.0 * oray.Origin.Z * oray.Direction.Z)
	c := (oray.Origin.X * oray.Origin.X) - (oray.Origin.Y * oray.Origin.Y) + (oray.Origin.Z * oray.Origin.Z)
	if (math.Abs(a) < primitives.EPSILON) && (math.Abs(b) > primitives.EPSILON) {
		t0 := -c / (2.0 * b)
		y0 := oray.Origin.Y + (t0 * oray.Direction.Y)
		if ( -1 < y0) && (y0 < 0) {
			hits = append(hits, Intersection{Distance:t0, Obj:cone})
		}
	} else if (math.Abs(a) > primitives.EPSILON) {
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
		if (-1 < y0) && (y0 < 0) {
			hits = append(hits, Intersection{Distance:t0, Obj:cone})
		}

		y1 := oray.Origin.Y + (t1 * oray.Direction.Y)
		if (-1 < y1) && (y1 < 0) {
			hits = append(hits, Intersection{Distance:t1, Obj:cone})
		}
	}

	// Cap checking only matters if Cone is closed
	if !cone.closed || math.Abs(oray.Direction.Y) < primitives.EPSILON {
		return hits
	}

	// Check bottom and top caps
	t := (-1 - oray.Origin.Y) / oray.Direction.Y
	if CheckCap(oray, t) {
		hits = append(hits, Intersection{Distance:t, Obj:cone})
	}

	return hits
}

// Normal Calculate the normal at a given point on the Cone
func (cone *Cone) Normal(worldPoint primitives.PV) primitives.PV {
	var objectNormal primitives.PV
	objectPoint := cone.WorldToObjectPV(worldPoint)
	distance := (objectPoint.X * objectPoint.X) + (objectPoint.Z * objectPoint.Z)
	if (distance < 1) && (objectPoint.Y <= (-1.0 + primitives.EPSILON)) {
		objectNormal = primitives.MakeVector(0, -1, 0)
	} else {
		objectNormal = primitives.MakeVector(objectPoint.X, math.Sqrt(distance), objectPoint.Z)
	}	
	worldNormal := cone.ObjectToWorldPV(objectNormal)
	worldNormal.W = 0.0
	return worldNormal.Normalize()
}

// UVMapping Return the 2D coordinates of an intersection point
func (cone *Cone) UVMapping(point primitives.PV) primitives.PV {
	objectPoint := cone.WorldToObjectPV(point)
	d := primitives.MakePoint(0, 0, 0).Subtract(objectPoint)
	return primitives.MakePoint(0.5 + math.Atan2(d.X, d.Z) / (2 * math.Pi), objectPoint.Y + 1, 0)
}
