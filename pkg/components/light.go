package components

import (
	"math"
	"github.com/factorion/graytracer/pkg/primitives"
	"github.com/factorion/graytracer/pkg/patterns"
	"github.com/factorion/graytracer/pkg/shapes"
)

// PointLight Basic light object a specific point
type PointLight struct {
	Intensity *patterns.RGB
	Position primitives.PV
}

// Lighting Basic lighting calculation function
func Lighting(shape shapes.Shape, light PointLight, point, eyeVector,
			  normalVector primitives.PV, shade float64) patterns.RGB {
	mat := shape.Material()
	effectiveColor := mat.Pat.ColorAt(shape.UVMapping(point)).Multiply(*light.Intensity)
	var diffuse, specular patterns.RGB
	lightv := light.Position.Subtract(point).Normalize()
	ambient := effectiveColor.Scale(mat.Ambient)
	if shade > 0 {
		if lightDotNormal := lightv.DotProduct(normalVector); lightDotNormal >= 0 {
			diffuse = effectiveColor.Scale(mat.Diffuse * lightDotNormal)
			reflectv := lightv.Negate().Reflect(normalVector)
			reflectDotEye := reflectv.DotProduct(eyeVector)
			if reflectDotEye >= 0 {
				factor := math.Pow(reflectDotEye, mat.Shininess)
				specular = light.Intensity.Scale(mat.Specular * factor)
			}
		}
	}
	return ambient.Add(diffuse.Add(specular).Scale(shade))
}
