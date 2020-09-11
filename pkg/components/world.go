package components

import (
	"sort"
	"github.com/factorion/graytracer/pkg/primitives"
	"github.com/factorion/graytracer/pkg/shapes"
)

// World Container for objects
type World struct {
	objects []shapes.Shape
	lights []PointLight
}

// AddObject Add a shape object to the world
func (w *World) AddObject(shape shapes.Shape) {
	w.objects = append(w.objects, shape)
}

// AddLight Add a light object to the world
func (w *World) AddLight(light PointLight) {
	w.lights = append(w.lights, light)
}

// Intersect Calculate the intersections from the ray to world objects
func (w World) Intersect(ray primitives.Ray) Intersections {
	var i Intersections
	for _, s := range w.objects {
		for _, h := range s.Intersect(ray) {
			i = append(i, Intersection{Distance:h, Obj:s})
		}
	}
	sort.Sort(i)
	return i
}

// ColorAt Calculate the color of a possible intersection hit
func (w World) ColorAt(ray primitives.Ray) primitives.RGB {
	result := primitives.MakeRGB(0, 0, 0)
	intersections := w.Intersect(ray)
	intersection, hit := intersections.Hit()
	if !hit {
		return result
	}
	comp := intersection.PrepareComputations(ray)
	for _, light := range w.lights {
		result = result.Add(Lighting(comp.Obj.Material(), light, comp.Point, comp.EyeVector, comp.NormalVector))
	}
	return result
}