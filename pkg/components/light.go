package components

import (
	"math"
	"github.com/factorion/graytracer/pkg/primitives"
)

// PointLight Basic light object a specific point
type PointLight struct {
	Intensity primitives.RGB
	Position primitives.PV
}

// Lighting Basic lighting calculation function
func Lighting(mat primitives.Material, light PointLight,
			  point, eyev, normalv primitives.PV) primitives.RGB {
	effectiveColor := mat.Color.Multiply(light.Intensity)
	var diffuse, specular primitives.RGB
	lightv := light.Position.Subtract(point).Normalize()
	ambient := effectiveColor.Scale(mat.Ambient)
	if lightDotNormal := lightv.DotProduct(normalv); lightDotNormal >= 0 {
		diffuse = effectiveColor.Scale(mat.Diffuse * lightDotNormal)
		reflectv := lightv.Negate().Reflect(normalv)
		reflectDotEye := reflectv.DotProduct(eyev)
		if reflectDotEye >= 0 {
			factor := math.Pow(reflectDotEye, mat.Shininess)
			specular = light.Intensity.Scale(mat.Specular * factor)
		}
	}
	return ambient.Add(diffuse.Add(specular))
}
