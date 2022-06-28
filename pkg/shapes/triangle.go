package shapes

import (
	"math"

	"github.com/factorion/graytracer/pkg/primitives"
)

// Triangle Represents a triangle
type Triangle struct {
	ShapeBase
	Point1, Point2, Point3, Edge1, Edge2, Normal1, Normal2, Normal3 primitives.PV
	smooth                                                          bool
}

// MakeTriangle  Create a triangle from three points
func MakeTriangle(point1, point2, point3 primitives.PV) *Triangle {
	edge1 := point2.Subtract(point1)
	edge1.W = 0
	edge2 := point3.Subtract(point1)
	edge2.W = 0
	normal := edge2.CrossProduct(edge1).Normalize()
	return &Triangle{MakeShapeBase(), point1, point2, point3, edge1, edge2, normal, normal, normal, false}
}

// MakeSmoothTriangle Create a smooth triangle from three points and three normals
func MakeSmoothTriangle(point1, point2, point3, normal1, normal2, normal3 primitives.PV) *Triangle {
	edge1 := point2.Subtract(point1)
	edge1.W = 0
	edge2 := point3.Subtract(point1)
	edge2.W = 0
	return &Triangle{MakeShapeBase(), point1, point2, point3, edge1, edge2, normal1, normal2, normal3, true}
}

// GetBounds Return an axis aligned bounding box for the triangle
func (t *Triangle) GetBounds() *Bounds {
	x_min, x_max := MinMax([]float64{t.Point1.X, t.Point2.X, t.Point3.X})
	y_min, y_max := MinMax([]float64{t.Point1.Y, t.Point2.Y, t.Point3.Y})
	z_min, z_max := MinMax([]float64{t.Point1.Z, t.Point2.Z, t.Point3.Z})
	bounds := &Bounds{Min: primitives.MakePoint(x_min, y_min, z_min), Max: primitives.MakePoint(x_max, y_max, z_max)}
	return bounds.Transform(t.transform)
}

// Intersect Check if a ray intersects
func (t *Triangle) Intersect(ray primitives.Ray) Intersections {
	hits := Intersections{}
	// convert ray to object space
	oray := ray.Transform(t.inverse)
	dce2 := oray.Direction.CrossProduct(t.Edge2) // Direction crossed with edge 2
	det := t.Edge1.DotProduct(dce2)
	if math.Abs(det) < primitives.EPSILON {
		return hits
	}
	f := 1.0 / det
	p1too := oray.Origin.Subtract(t.Point1) // origin to point 1
	u := f * p1too.DotProduct(dce2)
	if (u < 0) || (u > 1) {
		return hits
	}
	oce1 := p1too.CrossProduct(t.Edge1) // Cross product of origin and edge
	v := f * oray.Direction.DotProduct(oce1)
	if (v < 0) || ((u + v) > 1) {
		return hits
	}
	hits = append(hits, Intersection{Distance: (f * t.Edge2.DotProduct(oce1)), Obj: t, U: u, V: v})
	return hits
}

// Normal Calculate the normal at a given point on the sphere
func (t *Triangle) Normal(worldPoint primitives.PV, u, v float64) primitives.PV {
	var normal primitives.PV
	if !t.smooth {
		normal = t.Normal1
	} else {
		normal = t.Normal2.Scalar(u).Add(t.Normal3.Scalar(v)).Add(t.Normal1.Scalar(1 - u - v))
	}
	worldNormal := t.ObjectToWorldPV(normal)
	worldNormal.W = 0.0
	return worldNormal.Normalize()
}

// UVMapping Return the 2D coordinates of an intersection point
func (t *Triangle) UVMapping(point primitives.PV) primitives.PV {
	return primitives.MakePoint(point.X, point.Y, 0)
}
