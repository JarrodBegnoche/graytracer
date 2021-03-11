package components

import (
	"math"
	"github.com/factorion/graytracer/pkg/primitives"
	"github.com/factorion/graytracer/pkg/shapes"
)

// Computations Set of pre-computed values used for point detection
type Computations struct {
	shapes.Intersection
	Point, OverPoint, UnderPoint, EyeVector, NormalVector, ReflectVector primitives.PV
	Index1, Index2 float64
	Inside bool
}

// Schlick Calculate an approximation of the Fresnel effect
func (c Computations) Schlick() float64 {
	cos := c.EyeVector.DotProduct(c.NormalVector)
	if c.Index1 > c.Index2 {
		ratio := c.Index1 / c.Index2
		sin2Theta := math.Pow(ratio, 2) * (1.0 - math.Pow(cos, 2))
		if sin2Theta > 1.0 {
			return 1.0
		}
		cosTheta := math.Sqrt(1.0 - sin2Theta)
		cos = cosTheta
	}
	r0 := math.Pow((c.Index1 - c.Index2) / (c.Index1 + c.Index2), 2)
	return r0 + ((1 - r0) * math.Pow(1 - cos, 5))
}



// PrepareComputations Calculates the vectors at the point on the object
func PrepareComputations(i shapes.Intersection, ray primitives.Ray, xs shapes.Intersections) Computations {
	comp := Computations{Intersection:i}
	comp.Point = ray.Position(comp.Distance)
	comp.EyeVector = ray.Direction.Negate()
	comp.NormalVector = comp.Obj.Normal(comp.Point)
	if comp.NormalVector.DotProduct(comp.EyeVector) < 0 {
		comp.NormalVector = comp.NormalVector.Negate()
		comp.Inside = true
	} else{
		comp.Inside = false
	}
	scaledNormal := comp.NormalVector.Scalar(primitives.EPSILON)
	comp.OverPoint = comp.Point.Add(scaledNormal)
	comp.UnderPoint = comp.Point.Subtract(scaledNormal)
	comp.ReflectVector = ray.Direction.Reflect(comp.NormalVector)
	var stack []shapes.Shape
	for _, inter := range xs {
		if len(stack) == 0 {
			comp.Index1 = 1.0
		} else {
			comp.Index1 = stack[len(stack) - 1].Material().RefractiveIndex
		}
		if index := contains(stack, inter.Obj); index >= 0 {
			stack = append(stack[:index], stack[index + 1:]...)
		} else {
			stack = append(stack, inter.Obj)
		}
		if i == inter {
			if len(stack) == 0 {
				comp.Index2 = 1.0
			} else {
				comp.Index2 = stack[len(stack) - 1].Material().RefractiveIndex
			}
			break
		}
	}
	return comp
}

func contains(s []shapes.Shape, e shapes.Shape) int {
    for i, a := range s {
        if a == e {
            return i
        }
    }
    return -1
}
