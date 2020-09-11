package components

import (
	"testing"
	"github.com/factorion/graytracer/pkg/primitives"
)

func TestLighting(t *testing.T) {
	tables := []struct {
		mat primitives.Material
		position, eyev, normalv primitives.PV
		light PointLight
		result primitives.RGB
	}{
		{primitives.Material{Color:primitives.MakeRGB(1, 1, 1), Ambient:0.1, Diffuse:0.9,
							 Specular:0.9, Shininess:200},
		 primitives.MakePoint(0, 0, 0),
		 primitives.MakeVector(0, 0, -1),
		 primitives.MakeVector(0, 0, -1),
		 PointLight{Intensity:primitives.MakeRGB(1, 1, 1), Position:primitives.MakePoint(0, 0, -10)},
		 primitives.MakeRGB(1.9, 1.9, 1.9)},

		{primitives.Material{Color:primitives.MakeRGB(1, 1, 1), Ambient:0.1, Diffuse:0.9,
							 Specular:0.9, Shininess:200},
		 primitives.MakePoint(0, 0, 0),
		 primitives.MakeVector(0, 0.7071067811865476, -0.7071067811865476),
		 primitives.MakeVector(0, 0, -1),
		 PointLight{Intensity:primitives.MakeRGB(1, 1, 1), Position:primitives.MakePoint(0, 0, -10)},
		 primitives.MakeRGB(1.0, 1.0, 1.0)},

		{primitives.Material{Color:primitives.MakeRGB(1, 1, 1), Ambient:0.1, Diffuse:0.9,
							 Specular:0.9, Shininess:200},
		 primitives.MakePoint(0, 0, 0),
		 primitives.MakeVector(0, 0, -1),
		 primitives.MakeVector(0, 0, -1),
		 PointLight{Intensity:primitives.MakeRGB(1, 1, 1), Position:primitives.MakePoint(0, 10, -10)},
		 primitives.MakeRGB(0.7363961030678927, 0.7363961030678927, 0.7363961030678927)},

		{primitives.Material{Color:primitives.MakeRGB(1, 1, 1), Ambient:0.1, Diffuse:0.9,
							 Specular:0.9, Shininess:200},
		 primitives.MakePoint(0, 0, 0),
		 primitives.MakeVector(0, -0.7071067811865476, -0.7071067811865476),
		 primitives.MakeVector(0, 0, -1),
		 PointLight{Intensity:primitives.MakeRGB(1, 1, 1), Position:primitives.MakePoint(0, 10, -10)},
		 primitives.MakeRGB(1.6363961030678928, 1.6363961030678928, 1.6363961030678928)},
		
		{primitives.Material{Color:primitives.MakeRGB(1, 1, 1), Ambient:0.1, Diffuse:0.9,
							 Specular:0.9, Shininess:200},
		 primitives.MakePoint(0, 0, 0),
		 primitives.MakeVector(0, 0, -1),
		 primitives.MakeVector(0, 0, -1),
		 PointLight{Intensity:primitives.MakeRGB(1, 1, 1), Position:primitives.MakePoint(0, 0, 10)},
		 primitives.MakeRGB(0.1, 0.1, 0.1)},
	}
	for _, table := range tables {
		result := Lighting(table.mat, table.light, table.position, table.eyev, table.normalv)
		if !result.Equals(table.result) {
			t.Errorf("Expect %v, got %v", table.result, result)
		}
	}
}