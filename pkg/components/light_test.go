package components_test

import (
	"testing"
	"github.com/factorion/graytracer/pkg/components"
	"github.com/factorion/graytracer/pkg/primitives"
	"github.com/factorion/graytracer/pkg/patterns"
	"github.com/factorion/graytracer/pkg/shapes"
)

func TestLighting(t *testing.T) {
	tables := []struct {
		mat patterns.Material
		position, eyev, normalv primitives.PV
		light components.PointLight
		shade float64
		result *patterns.RGB
	}{
		{patterns.Material{Pat:patterns.MakeRGB(1, 1, 1), Ambient:0.1, Diffuse:0.9,
							 Specular:0.9, Shininess:200, Reflective:0},
		 primitives.MakePoint(0, 0, 0),
		 primitives.MakeVector(0, 0, -1),
		 primitives.MakeVector(0, 0, -1),
		 components.PointLight{Intensity:patterns.MakeRGB(1, 1, 1), Position:primitives.MakePoint(0, 0, -10)},
		 1, patterns.MakeRGB(1.9, 1.9, 1.9)},

		{patterns.Material{Pat:patterns.MakeRGB(1, 1, 1), Ambient:0.1, Diffuse:0.9,
							 Specular:0.9, Shininess:200, Reflective:0},
		 primitives.MakePoint(0, 0, 0),
		 primitives.MakeVector(0, 0.7071067811865476, -0.7071067811865476),
		 primitives.MakeVector(0, 0, -1),
		 components.PointLight{Intensity:patterns.MakeRGB(1, 1, 1), Position:primitives.MakePoint(0, 0, -10)},
		 1, patterns.MakeRGB(1.0, 1.0, 1.0)},

		{patterns.Material{Pat:patterns.MakeRGB(1, 1, 1), Ambient:0.1, Diffuse:0.9,
							 Specular:0.9, Shininess:200, Reflective:0},
		 primitives.MakePoint(0, 0, 0),
		 primitives.MakeVector(0, 0, -1),
		 primitives.MakeVector(0, 0, -1),
		 components.PointLight{Intensity:patterns.MakeRGB(1, 1, 1), Position:primitives.MakePoint(0, 10, -10)},
		 1, patterns.MakeRGB(0.7363961030678927, 0.7363961030678927, 0.7363961030678927)},

		{patterns.Material{Pat:patterns.MakeRGB(1, 1, 1), Ambient:0.1, Diffuse:0.9,
							 Specular:0.9, Shininess:200, Reflective:0},
		 primitives.MakePoint(0, 0, 0),
		 primitives.MakeVector(0, -0.7071067811865476, -0.7071067811865476),
		 primitives.MakeVector(0, 0, -1),
		 components.PointLight{Intensity:patterns.MakeRGB(1, 1, 1), Position:primitives.MakePoint(0, 10, -10)},
		 1, patterns.MakeRGB(1.6363961030678928, 1.6363961030678928, 1.6363961030678928)},
		
		{patterns.Material{Pat:patterns.MakeRGB(1, 1, 1), Ambient:0.1, Diffuse:0.9,
							 Specular:0.9, Shininess:200, Reflective:0},
		 primitives.MakePoint(0, 0, 0),
		 primitives.MakeVector(0, 0, -1),
		 primitives.MakeVector(0, 0, -1),
		 components.PointLight{Intensity:patterns.MakeRGB(1, 1, 1), Position:primitives.MakePoint(0, 0, 10)},
		 1, patterns.MakeRGB(0.1, 0.1, 0.1)},

		{patterns.Material{Pat:patterns.MakeRGB(1, 1, 1), Ambient:0.1, Diffuse:0.9,
		 Specular:0.9, Shininess:200, Reflective:0},
		 primitives.MakePoint(0, 0, 0),
		 primitives.MakeVector(0, 0, -1),
		 primitives.MakeVector(0, 0, -1),
	 	 components.PointLight{Intensity:patterns.MakeRGB(1, 1, 1), Position:primitives.MakePoint(0, 0, -10)},
		 0, patterns.MakeRGB(0.1, 0.1, 0.1)},
	}
	for _, table := range tables {
		sphere := shapes.MakeSphere()
		sphere.SetMaterial(table.mat)
		result := components.Lighting(sphere, table.light, table.position, table.eyev, table.normalv, table.shade)
		if !result.Equals(*table.result) {
			t.Errorf("Expect %v, got %v", table.result, result)
		}
	}
}
