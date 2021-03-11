package components

import (
	"math"
	"sort"
	"github.com/factorion/graytracer/pkg/primitives"
	"github.com/factorion/graytracer/pkg/patterns"
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
func (w World) Intersect(ray primitives.Ray) shapes.Intersections {
	var i shapes.Intersections
	for _, s := range w.objects {
		i = append(i, s.Intersect(ray)...)
	}
	sort.Sort(i)
	return i
}

// ReflectedColor Calculate the color of the reflected ray
func (w World) ReflectedColor(comps Computations, remaining int) patterns.RGB {
	reflective := comps.Obj.Material().Reflective
	if reflective == 0 {
		return *patterns.MakeRGB(0, 0, 0)
	}
	reflectRay := primitives.Ray{Origin:comps.OverPoint, Direction:comps.ReflectVector}
	return w.ColorAt(reflectRay, remaining - 1).Scale(reflective)
}

// RefractedColor Calculate the color of the refracted ray
func (w World) RefractedColor(comps Computations, remaining int) patterns.RGB {
	transparency := comps.Obj.Material().Transparency
	if transparency == 0 {
		return *patterns.MakeRGB(0, 0, 0)
	}
	nRatio := comps.Index1 / comps.Index2
	cosi := comps.EyeVector.DotProduct(comps.NormalVector)
	sin2t := math.Pow(nRatio, 2) * (1 - math.Pow(cosi, 2))
	if sin2t > 1 {
		// Total internal reflection
		return *patterns.MakeRGB(0, 0, 0)
	}
	cost := math.Sqrt(1 - sin2t)
	direction := comps.NormalVector.Scalar((nRatio * cosi) - cost).Subtract(comps.EyeVector.Scalar(nRatio))
	refractRay := primitives.Ray{Origin:comps.UnderPoint, Direction:direction}
	return w.ColorAt(refractRay, remaining - 1).Scale(transparency)
}

// ColorAt Calculate the color of a possible intersection hit
func (w World) ColorAt(ray primitives.Ray, remaining int) patterns.RGB {
	surface := *patterns.MakeRGB(0, 0, 0)
	if remaining <= 0 {
		return surface
	}
	intersections := w.Intersect(ray)
	intersection, hit := intersections.Hit()
	if !hit {
		return surface
	}
	comp := PrepareComputations(intersection, ray, intersections)
	for _, light := range w.lights {
		shade := 1.0
		shadowVector := light.Position.Subtract(comp.OverPoint)
		distance := shadowVector.Magnitude()
		shadowRay := primitives.Ray{Origin:comp.OverPoint,
									Direction:shadowVector.Normalize()}
		shadowIntersections := w.Intersect(shadowRay)
		_, shadowHit := shadowIntersections.Hit()
		if shadowHit {
			shadowShapes := make(map[shapes.Shape]bool)
			for _, shadeIntersection := range shadowIntersections {
				if shadeIntersection.Distance > distance {
					break
				}
				if _, exists := shadowShapes[shadeIntersection.Obj]; !exists && shadeIntersection.Distance > 0 {
					shadowShapes[shadeIntersection.Obj] = true
					shade *= shadeIntersection.Obj.Material().Transparency
				}
			}
		}
		surface = surface.Add(Lighting(comp.Obj, light, comp.Point,
							  comp.EyeVector, comp.NormalVector, shade))
	}
	reflected := w.ReflectedColor(comp, remaining)
	refracted := w.RefractedColor(comp, remaining)
	material := comp.Obj.Material()
	if material.Reflective > 0 && material.Transparency > 0 {
		reflectance := comp.Schlick()
		return surface.Add(reflected.Scale(reflectance)).Add(refracted.Scale(1 - reflectance))
	}
	return surface.Add(reflected).Add(refracted)
}
