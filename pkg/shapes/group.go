package shapes

import (
	"github.com/factorion/graytracer/pkg/primitives"
)

// Group Represents a group of other shapes
type Group struct {
	ShapeBase
	shapes []Shape
}

// MakeGroup Make an empty set of shapes
func MakeGroup() *Group {
	return &Group{MakeShapeBase(), []Shape{}}
}

// AddShape Add a shape to the group and set its parent
func (g *Group) AddShape(shape Shape) {
	g.shapes = append(g.shapes, shape)
	shape.SetParent(g)
}

// Intersect Check if a ray intersects
func (g *Group) Intersect(r primitives.Ray) Intersections {
	hits := Intersections{}
	// convert ray to object space
	oray := r.Transform(g.inverse)
	for _, shape := range g.shapes {
		hits = append(hits, shape.Intersect(oray)...)
	}
	return hits
}

// Normal Calculate the normal at a given point on the sphere
func (g *Group) Normal(worldPoint primitives.PV) primitives.PV {
	// Only exists for Interface, should never be called
	return primitives.MakeVector(0, 1, 0)
}

// UVMapping Return the 2D coordinates of an intersection point
func (g *Group) UVMapping(point primitives.PV) primitives.PV {
	// Only exists for Interface, should never be called
	return primitives.MakePoint(point.X, point.Y, 0)
}
