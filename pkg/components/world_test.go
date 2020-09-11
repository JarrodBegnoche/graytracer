package components

import (
	"testing"
	"github.com/factorion/graytracer/pkg/primitives"
	"github.com/factorion/graytracer/pkg/shapes"
)

func TestWorldIntersect(t *testing.T) {
	tables := []struct {
		shapes []shapes.Shape
		mats []primitives.Material
		transforms []primitives.Matrix
		light PointLight
		ray primitives.Ray
		distances []float64
	}{
		{[]shapes.Shape{shapes.MakeSphere(), shapes.MakeSphere()},
		 []primitives.Material{primitives.Material{Color:primitives.MakeRGB(0.8, 1.0, 0.6)},
			                   primitives.Material{}},
		 []primitives.Matrix{primitives.MakeIdentityMatrix(4), primitives.Scaling(0.5, 0.5, 0.5)},
		 PointLight{Intensity:primitives.MakeRGB(1, 1, 1), Position:primitives.MakePoint(-10, 10, -10)},
		 primitives.Ray{Origin:primitives.MakePoint(0, 0, -5), Direction:primitives.MakeVector(0, 0, 1)},
		 []float64{4, 4.5, 5.5, 6}},
	}
	for _, table := range tables {
		world := &World{}
		for i := 0; i < len(table.shapes); i++ {
			table.shapes[i].SetMaterial(table.mats[i])
			table.shapes[i].SetTransform(table.transforms[i])
			world.AddObject(table.shapes[i])
		}
		world.AddLight(table.light)
		hits := world.Intersect(table.ray)
		if len(hits) != len(table.distances) {
			t.Errorf("Incorrect number of intersections: %v", len(hits))
		}
		for i, v := range hits {
			if v.Distance != table.distances[i] {
				t.Errorf("Incorrect hit distance, expect %v, got %v", table.distances[i], v.Distance)
			}
		}
	}
}

func TestWorldColorAt(t *testing.T) {
	tables := []struct {
		shapes []shapes.Shape
		mats []primitives.Material
		transforms []primitives.Matrix
		light PointLight
		ray primitives.Ray
		result primitives.RGB
	}{
		{[]shapes.Shape{shapes.MakeSphere(), shapes.MakeSphere()},
		 []primitives.Material{primitives.Material{Color:primitives.MakeRGB(0.8, 1.0, 0.6),
												   Ambient:0.1, Diffuse:0.7, Specular:0.2, Shininess:200},
			                   primitives.Material{Color:primitives.MakeRGB(1, 1, 1),
												   Ambient:0.1, Diffuse:0.9, Specular:0.9, Shininess:200}},
		 []primitives.Matrix{primitives.MakeIdentityMatrix(4), primitives.Scaling(0.5, 0.5, 0.5)},
		 PointLight{Intensity:primitives.MakeRGB(1, 1, 1), Position:primitives.MakePoint(-10, 10, -10)},
		 primitives.Ray{Origin:primitives.MakePoint(0, 0, -5), Direction:primitives.MakeVector(0, 0, 1)},
		 primitives.MakeRGB(0.38066119308103435, 0.47582649135129296, 0.28549589481077575)},

		{[]shapes.Shape{shapes.MakeSphere(), shapes.MakeSphere()},
		 []primitives.Material{primitives.Material{Color:primitives.MakeRGB(0.8, 1.0, 0.6),
		    									   Ambient:0.1, Diffuse:0.7, Specular:0.2, Shininess:200},
							   primitives.Material{Color:primitives.MakeRGB(1, 1, 1),
												   Ambient:0.1, Diffuse:0.9, Specular:0.9, Shininess:200}},
		 []primitives.Matrix{primitives.MakeIdentityMatrix(4), primitives.Scaling(0.5, 0.5, 0.5)},
		 PointLight{Intensity:primitives.MakeRGB(1, 1, 1), Position:primitives.MakePoint(0, 0.25, 0)},
		 primitives.Ray{Origin:primitives.MakePoint(0, 0, 0), Direction:primitives.MakeVector(0, 0, 1)},
		 primitives.MakeRGB(0.9049844720832575, 0.9049844720832575, 0.9049844720832575)},

		{[]shapes.Shape{shapes.MakeSphere(), shapes.MakeSphere()},
		 []primitives.Material{primitives.Material{Color:primitives.MakeRGB(0.8, 1.0, 0.6),
		    									   Ambient:0.1, Diffuse:0.9, Specular:0.9, Shininess:200},
							   primitives.Material{Color:primitives.MakeRGB(1, 1, 1),
												   Ambient:0.1, Diffuse:0.9, Specular:0.9, Shininess:200}},
		 []primitives.Matrix{primitives.MakeIdentityMatrix(4), primitives.Scaling(0.5, 0.5, 0.5)},
		 PointLight{Intensity:primitives.MakeRGB(1, 1, 1), Position:primitives.MakePoint(-10, 10, -10)},
		 primitives.Ray{Origin:primitives.MakePoint(0, 0, -5), Direction:primitives.MakeVector(0, 1, 0)},
		 primitives.MakeRGB(0, 0, 0)},

		{[]shapes.Shape{shapes.MakeSphere(), shapes.MakeSphere()},
		 []primitives.Material{primitives.Material{Color:primitives.MakeRGB(0.8, 1.0, 0.6),
		    									   Ambient:1, Diffuse:0.9, Specular:0.9, Shininess:200},
							   primitives.Material{Color:primitives.MakeRGB(1, 1, 1),
												   Ambient:1, Diffuse:0.9, Specular:0.9, Shininess:200}},
		 []primitives.Matrix{primitives.MakeIdentityMatrix(4), primitives.Scaling(0.5, 0.5, 0.5)},
		 PointLight{Intensity:primitives.MakeRGB(1, 1, 1), Position:primitives.MakePoint(0, 0.25, 0)},
		 primitives.Ray{Origin:primitives.MakePoint(0, 0, 0.75), Direction:primitives.MakeVector(0, 0, -1)},
		 primitives.MakeRGB(1, 1, 1)},
	}
	for _, table := range tables {
		world := &World{}
		for i := 0; i < len(table.shapes); i++ {
			table.shapes[i].SetMaterial(table.mats[i])
			table.shapes[i].SetTransform(table.transforms[i])
			world.AddObject(table.shapes[i])
		}
		world.AddLight(table.light)
		result := world.ColorAt(table.ray)
		if !result.Equals(table.result) {
			t.Errorf("Expected result color %v, got %v", table.result, result)
		}
	}
}
