package primitives

import (
	"testing"
)

func TestLighting(t *testing.T) {
	tables := []struct {
		mat Material
		position, eyev, normalv PV
		light PointLight
		result RGB
	}{
		{Material{Color:MakeRGB(1, 1, 1), Ambient:0.1, Diffuse:0.9, Specular:0.9, Shininess:200},
		 MakePoint(0, 0, 0), MakeVector(0, 0, -1), MakeVector(0, 0, -1),
		 PointLight{Intensity:MakeRGB(1, 1, 1), Position:MakePoint(0, 0, -10)},
		 MakeRGB(1.9, 1.9, 1.9)},

		{Material{Color:MakeRGB(1, 1, 1), Ambient:0.1, Diffuse:0.9, Specular:0.9, Shininess:200},
		 MakePoint(0, 0, 0), MakeVector(0, 0.7071067811865476, -0.7071067811865476), MakeVector(0, 0, -1),
		 PointLight{Intensity:MakeRGB(1, 1, 1), Position:MakePoint(0, 0, -10)},
		 MakeRGB(1.0, 1.0, 1.0)},

		{Material{Color:MakeRGB(1, 1, 1), Ambient:0.1, Diffuse:0.9, Specular:0.9, Shininess:200},
		 MakePoint(0, 0, 0), MakeVector(0, 0, -1), MakeVector(0, 0, -1),
		 PointLight{Intensity:MakeRGB(1, 1, 1), Position:MakePoint(0, 10, -10)},
		 MakeRGB(0.7363961030678927, 0.7363961030678927, 0.7363961030678927)},

		{Material{Color:MakeRGB(1, 1, 1), Ambient:0.1, Diffuse:0.9, Specular:0.9, Shininess:200},
		 MakePoint(0, 0, 0), MakeVector(0, -0.7071067811865476, -0.7071067811865476), MakeVector(0, 0, -1),
		 PointLight{Intensity:MakeRGB(1, 1, 1), Position:MakePoint(0, 10, -10)},
		 MakeRGB(1.6363961030678928, 1.6363961030678928, 1.6363961030678928)},
		
		{Material{Color:MakeRGB(1, 1, 1), Ambient:0.1, Diffuse:0.9, Specular:0.9, Shininess:200},
		 MakePoint(0, 0, 0), MakeVector(0, 0, -1), MakeVector(0, 0, -1),
		 PointLight{Intensity:MakeRGB(1, 1, 1), Position:MakePoint(0, 0, 10)},
		 MakeRGB(0.1, 0.1, 0.1)},
	}
	for _, table := range tables {
		result := Lighting(table.mat, table.light, table.position, table.eyev, table.normalv)
		if !result.Equals(table.result) {
			t.Errorf("Expect %v, got %v", table.result, result)
		}
	}
}