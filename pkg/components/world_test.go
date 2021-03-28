package components_test

import (
	"math"
	"testing"
	"github.com/factorion/graytracer/pkg/primitives"
	"github.com/factorion/graytracer/pkg/components"
	"github.com/factorion/graytracer/pkg/patterns"
	"github.com/factorion/graytracer/pkg/shapes"
)

func TestWorldIntersect(t *testing.T) {
	tables := []struct {
		shapes []shapes.Shape
		mats []patterns.Material
		transforms []primitives.Matrix
		light components.PointLight
		ray primitives.Ray
		distances []float64
	}{
		{[]shapes.Shape{shapes.MakeSphere(), shapes.MakeSphere()},
		 []patterns.Material{{Pat:patterns.MakeRGB(0.8, 1.0, 0.6)}, {}},
		 []primitives.Matrix{primitives.MakeIdentityMatrix(4), primitives.Scaling(0.5, 0.5, 0.5)},
		 components.PointLight{Intensity:patterns.MakeRGB(1, 1, 1), Position:primitives.MakePoint(-10, 10, -10)},
		 primitives.Ray{Origin:primitives.MakePoint(0, 0, -5), Direction:primitives.MakeVector(0, 0, 1)},
		 []float64{4, 4.5, 5.5, 6}},
	}
	for _, table := range tables {
		world := &components.World{}
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

func TestReflectedColor(t *testing.T) {
	tables := []struct {
		shapes []shapes.Shape
		mats []patterns.Material
		transforms []primitives.Matrix
		light components.PointLight
		ray primitives.Ray
		result *patterns.RGB
	}{
		{[]shapes.Shape{shapes.MakeSphere(), shapes.MakeSphere(), shapes.MakePlane()},
		 []patterns.Material{{Pat:patterns.MakeRGB(0.8, 1.0, 0.6), Ambient:0.1,
							  Diffuse:0.7, Specular:0.2, Shininess:200, Reflective:0,
							  Transparency:0, RefractiveIndex:1},
							 {Pat:patterns.MakeRGB(1, 1, 1), Ambient:0.1,
							  Diffuse:0.9, Specular:0.9, Shininess:200, Reflective:0,
							  Transparency:0, RefractiveIndex:1},
							 {Pat:patterns.MakeRGB(1, 1, 1), Ambient:0.1,
							  Diffuse:0.9, Specular:0.9, Shininess:200, Reflective:0.5,
							  Transparency:0, RefractiveIndex:1}},
		 []primitives.Matrix{primitives.MakeIdentityMatrix(4),
							 primitives.Scaling(0.5, 0.5, 0.5),
							 primitives.Translation(0, -1, 0)},
		 components.PointLight{Intensity:patterns.MakeRGB(1, 1, 1), Position:primitives.MakePoint(-10, 10, -10)},
		 primitives.Ray{Origin:primitives.MakePoint(0, 0, -3),
					    Direction:primitives.MakeVector(0, -0.7071067811865476, 0.7071067811865476)},
		 patterns.MakeRGB(0.19033059826435575, 0.23791324783044465, 0.1427479486982668)},
	}
	for _, table := range tables {
		world := &components.World{}
		for i := 0; i < len(table.shapes); i++ {
			table.shapes[i].SetMaterial(table.mats[i])
			table.shapes[i].SetTransform(table.transforms[i])
			world.AddObject(table.shapes[i])
		}
		world.AddLight(table.light)
		intersections := world.Intersect(table.ray)
		intersection, _ := intersections.Hit()
		comp := components.PrepareComputations(intersection, table.ray, shapes.Intersections{})
		result := world.ReflectedColor(comp, 5)
		if !result.Equals(*table.result) {
			t.Errorf("Expected result color %v, got %v", table.result, result)
		}
	}
}

func TestRefractedColor(t *testing.T) {
	tables := []struct {
		shapes []shapes.Shape
		mats []patterns.Material
		transforms []primitives.Matrix
		light components.PointLight
		ray primitives.Ray
		result *patterns.RGB
	}{
		{[]shapes.Shape{shapes.MakeSphere(), shapes.MakeSphere()},
		 []patterns.Material{{Pat:patterns.MakeRGB(0.8, 1.0, 0.6), Ambient:0.1,
							  Diffuse:0.7, Specular:0.2, Shininess:200, Reflective:0,
							  Transparency:1.0, RefractiveIndex:1.5},
							 {Pat:patterns.MakeRGB(1, 1, 1), Ambient:0.1,
							  Diffuse:0.9, Specular:0.9, Shininess:200, Reflective:0,
							  Transparency:0, RefractiveIndex:1}},
		 []primitives.Matrix{primitives.MakeIdentityMatrix(4),
							 primitives.Scaling(0.5, 0.5, 0.5)},
		 components.PointLight{Intensity:patterns.MakeRGB(1, 1, 1), Position:primitives.MakePoint(-10, 10, -10)},
		 primitives.Ray{Origin:primitives.MakePoint(0, 0, 0.7071067811865476),
					    Direction:primitives.MakeVector(0, 1, 0)},
		 patterns.MakeRGB(0, 0, 0)},
	}
	for _, table := range tables {
		world := &components.World{}
		for i := 0; i < len(table.shapes); i++ {
			table.shapes[i].SetMaterial(table.mats[i])
			table.shapes[i].SetTransform(table.transforms[i])
			world.AddObject(table.shapes[i])
		}
		world.AddLight(table.light)
		xs := world.Intersect(table.ray)
		intersection, _ := xs.Hit()
		comp := components.PrepareComputations(intersection, table.ray, xs)
		result := world.RefractedColor(comp, 5)
		if !result.Equals(*table.result) {
			t.Errorf("Expected result color %v, got %v", table.result, result)
		}
	}
}

func TestWorldColorAt(t *testing.T) {
	tables := []struct {
		shapes []shapes.Shape
		mats []patterns.Material
		transforms []primitives.Matrix
		light components.PointLight
		ray primitives.Ray
		result *patterns.RGB
	}{
		{[]shapes.Shape{shapes.MakeSphere(), shapes.MakeSphere()},
		 []patterns.Material{{Pat:patterns.MakeRGB(0.8, 1.0, 0.6), Ambient:0.1,
							  Diffuse:0.7, Specular:0.2, Shininess:200, Reflective:0,
							  Transparency:0, RefractiveIndex:1},
			                 {Pat:patterns.MakeRGB(1, 1, 1), Ambient:0.1,
							  Diffuse:0.9, Specular:0.9, Shininess:200, Reflective:0,
							  Transparency:0, RefractiveIndex:1}},
		 []primitives.Matrix{primitives.MakeIdentityMatrix(4), primitives.Scaling(0.5, 0.5, 0.5)},
		 components.PointLight{Intensity:patterns.MakeRGB(1, 1, 1), Position:primitives.MakePoint(-10, 10, -10)},
		 primitives.Ray{Origin:primitives.MakePoint(0, 0, -5), Direction:primitives.MakeVector(0, 0, 1)},
		 patterns.MakeRGB(0.38066119308103435, 0.47582649135129296, 0.28549589481077575)},

		{[]shapes.Shape{shapes.MakeSphere(), shapes.MakeSphere()},
		 []patterns.Material{{Pat:patterns.MakeRGB(0.8, 1.0, 0.6), Ambient:0.1,
							  Diffuse:0.7, Specular:0.2, Shininess:200, Reflective:0,
							  Transparency:0, RefractiveIndex:1},
							 {Pat:patterns.MakeRGB(1, 1, 1), Ambient:0.1,
							  Diffuse:0.9, Specular:0.9, Shininess:200, Reflective:0,
							  Transparency:0, RefractiveIndex:1}},
		 []primitives.Matrix{primitives.MakeIdentityMatrix(4), primitives.Scaling(0.5, 0.5, 0.5)},
		 components.PointLight{Intensity:patterns.MakeRGB(1, 1, 1), Position:primitives.MakePoint(0, 0.25, 0)},
		 primitives.Ray{Origin:primitives.MakePoint(0, 0, 0), Direction:primitives.MakeVector(0, 0, 1)},
		 patterns.MakeRGB(0.9049844720832575, 0.9049844720832575, 0.9049844720832575)},

		{[]shapes.Shape{shapes.MakeSphere(), shapes.MakeSphere()},
		 []patterns.Material{patterns.MakeDefaultMaterial(), patterns.MakeDefaultMaterial()},
		 []primitives.Matrix{primitives.MakeIdentityMatrix(4), primitives.Translation(0, 0, 10)},
		 components.PointLight{Intensity:patterns.MakeRGB(1, 1, 1), Position:primitives.MakePoint(0, 0, -10)},
		 primitives.Ray{Origin:primitives.MakePoint(0, 0, 5), Direction:primitives.MakeVector(0, 0, 1)},
		 patterns.MakeRGB(0.1, 0.1, 0.1)},

		{[]shapes.Shape{shapes.MakeSphere(), shapes.MakeSphere()},
		 []patterns.Material{{Pat:patterns.MakeRGB(0.8, 1.0, 0.6), Ambient:0.1,
							  Diffuse:0.9, Specular:0.9, Shininess:200, Reflective:0,
							  Transparency:0, RefractiveIndex:1},
							 {Pat:patterns.MakeRGB(1, 1, 1), Ambient:0.1,
							  Diffuse:0.9, Specular:0.9, Shininess:200, Reflective:0,
							  Transparency:0, RefractiveIndex:1}},
		 []primitives.Matrix{primitives.MakeIdentityMatrix(4), primitives.Scaling(0.5, 0.5, 0.5)},
		 components.PointLight{Intensity:patterns.MakeRGB(1, 1, 1), Position:primitives.MakePoint(-10, 10, -10)},
		 primitives.Ray{Origin:primitives.MakePoint(0, 0, -5), Direction:primitives.MakeVector(0, 1, 0)},
		 patterns.MakeRGB(0, 0, 0)},

		{[]shapes.Shape{shapes.MakeSphere(), shapes.MakeSphere()},
		 []patterns.Material{{Pat:patterns.MakeRGB(0.8, 1.0, 0.6), Ambient:1,
							  Diffuse:0.9, Specular:0.9, Shininess:200, Reflective:0,
							  Transparency:0, RefractiveIndex:1},
							 {Pat:patterns.MakeRGB(1, 1, 1), Ambient:1,
							  Diffuse:0.9, Specular:0.9, Shininess:200, Reflective:0,
							  Transparency:0, RefractiveIndex:1}},
		 []primitives.Matrix{primitives.MakeIdentityMatrix(4), primitives.Scaling(0.5, 0.5, 0.5)},
		 components.PointLight{Intensity:patterns.MakeRGB(1, 1, 1), Position:primitives.MakePoint(0, 0.25, 0)},
		 primitives.Ray{Origin:primitives.MakePoint(0, 0, 0.75), Direction:primitives.MakeVector(0, 0, -1)},
		 patterns.MakeRGB(1, 1, 1)},

		{[]shapes.Shape{shapes.MakeSphere(), shapes.MakeSphere(), shapes.MakePlane()},
		 []patterns.Material{{Pat:patterns.MakeRGB(0.8, 1.0, 0.6), Ambient:0.1,
							  Diffuse:0.7, Specular:0.2, Shininess:200, Reflective:0,
							  Transparency:0, RefractiveIndex:1},
							 {Pat:patterns.MakeRGB(1, 1, 1), Ambient:0.1,
							  Diffuse:0.9, Specular:0.9, Shininess:200, Reflective:0,
							  Transparency:0, RefractiveIndex:1},
							 {Pat:patterns.MakeRGB(1, 1, 1), Ambient:0.1,
							  Diffuse:0.9, Specular:0.9, Shininess:200, Reflective:0.5,
							  Transparency:0, RefractiveIndex:1}},
		 []primitives.Matrix{primitives.MakeIdentityMatrix(4),
							 primitives.Scaling(0.5, 0.5, 0.5),
							 primitives.Translation(0, -1, 0)},
		 components.PointLight{Intensity:patterns.MakeRGB(1, 1, 1), Position:primitives.MakePoint(-10, 10, -10)},
		 primitives.Ray{Origin:primitives.MakePoint(0, 0, -3),
					    Direction:primitives.MakeVector(0, -0.7071067811865476, 0.7071067811865476)},
		 patterns.MakeRGB(0.8767559872458571, 0.9243386368119461, 0.8291733376797682)},

		{[]shapes.Shape{shapes.MakeSphere(), shapes.MakeSphere(),
						shapes.MakePlane(), shapes.MakeSphere()},
		 []patterns.Material{{Pat:patterns.MakeRGB(0.8, 1.0, 0.6), Ambient:0.1,
							  Diffuse:0.7, Specular:0.2, Shininess:200, Reflective:0,
							  Transparency:0, RefractiveIndex:1},
							 {Pat:patterns.MakeRGB(1, 1, 1), Ambient:0.1,
							  Diffuse:0.9, Specular:0.9, Shininess:200, Reflective:0,
							  Transparency:0, RefractiveIndex:1},
							 {Pat:patterns.MakeRGB(1, 1, 1), Ambient:0.1,
							  Diffuse:0.9, Specular:0.9, Shininess:200, Reflective:0,
							  Transparency:0.5, RefractiveIndex:1.5},
							 {Pat:patterns.MakeRGB(1, 0, 0), Ambient:0.5,
							  Diffuse:0.9, Specular:0.9, Shininess:200, Reflective:0,
							  Transparency:0, RefractiveIndex:1.0}},
		 []primitives.Matrix{primitives.MakeIdentityMatrix(4),
							 primitives.Scaling(0.5, 0.5, 0.5),
							 primitives.Translation(0, -1, 0),
							 primitives.Translation(0, -3.5, -0.5)},
		 components.PointLight{Intensity:patterns.MakeRGB(1, 1, 1), Position:primitives.MakePoint(-10, 10, -10)},
		 primitives.Ray{Origin:primitives.MakePoint(0, 0, -3),
					    Direction:primitives.MakeVector(0, -0.7071067811865476, 0.7071067811865476)},
		 patterns.MakeRGB(1.1254657872747564, 0.6864253889815014, 0.6864253889815014)},
		
		{[]shapes.Shape{shapes.MakeSphere(), shapes.MakeSphere(),
						shapes.MakePlane(), shapes.MakeSphere()},
		 []patterns.Material{{Pat:patterns.MakeRGB(0.8, 1.0, 0.6), Ambient:0.1,
							  Diffuse:0.7, Specular:0.2, Shininess:200, Reflective:0,
							  Transparency:0, RefractiveIndex:1},
							 {Pat:patterns.MakeRGB(1, 1, 1), Ambient:0.1,
							  Diffuse:0.9, Specular:0.9, Shininess:200, Reflective:0,
							  Transparency:0, RefractiveIndex:1},
							 {Pat:patterns.MakeRGB(1, 1, 1), Ambient:0.1,
							  Diffuse:0.9, Specular:0.9, Shininess:200, Reflective:0.5,
							  Transparency:0.5, RefractiveIndex:1.5},
							 {Pat:patterns.MakeRGB(1, 0, 0), Ambient:0.5,
							  Diffuse:0.9, Specular:0.9, Shininess:200, Reflective:0,
							  Transparency:0, RefractiveIndex:1.0}},
		 []primitives.Matrix{primitives.MakeIdentityMatrix(4),
							 primitives.Scaling(0.5, 0.5, 0.5),
							 primitives.Translation(0, -1, 0),
							 primitives.Translation(0, -3.5, -0.5)},
		 components.PointLight{Intensity:patterns.MakeRGB(1, 1, 1), Position:primitives.MakePoint(-10, 10, -10)},
		 primitives.Ray{Origin:primitives.MakePoint(0, 0, -3),
						Direction:primitives.MakeVector(0, -0.7071067811865476, 0.7071067811865476)},
		 patterns.MakeRGB(1.1150027467686325, 0.6964342263843869, 0.6924306914232327)},
	}
	for _, table := range tables {
		world := &components.World{}
		for i := 0; i < len(table.shapes); i++ {
			table.shapes[i].SetMaterial(table.mats[i])
			table.shapes[i].SetTransform(table.transforms[i])
			world.AddObject(table.shapes[i])
		}
		world.AddLight(table.light)
		result := world.ColorAt(table.ray, 5)
		if !result.Equals(*table.result) {
			t.Errorf("Expected result color %v, got %v", table.result, result)
		}
	}
}

func TestInfiniteReflection(t *testing.T) {
	world := &components.World{}
	world.AddLight(components.PointLight{Intensity:patterns.MakeRGB(1, 1, 1), Position:primitives.MakePoint(0, 0, 0)})
	lower := shapes.MakePlane()
	lowerMat := lower.Material()
	lowerMat.Reflective = 1
	lower.SetMaterial(lowerMat)
	lower.SetTransform(primitives.Translation(0, -10, 0))
	world.AddObject(lower)
	upper := shapes.MakePlane()
	upperMat := upper.Material()
	upperMat.Reflective = 1
	upper.SetMaterial(upperMat)
	upper.SetTransform(primitives.Translation(0, 10, 0))
	world.AddObject(upper)
	ray := primitives.Ray{Origin:primitives.MakePoint(0, 0, 0), Direction:primitives.MakeVector(0, 1, 0)}
	col := world.ColorAt(ray, 5)
	t.Log(col)
}

func TestSchlick(t *testing.T) {
	tables := []struct {
		shape shapes.Shape
		mat patterns.Material
		transform primitives.Matrix
		ray primitives.Ray
		reflectance float64
	}{
		{shapes.MakeSphere(),
		 patterns.Material{Pat:patterns.MakeRGB(1, 1, 1), Ambient:0.1, Diffuse:0.9, Specular:0.9,
						   Reflective:0, Transparency:1.0, RefractiveIndex:1.5},
		 primitives.MakeIdentityMatrix(4),
		 primitives.Ray{Origin:primitives.MakePoint(0, 0, -0.7071067811865476),
						Direction:primitives.MakeVector(0, 1, 0)},
		 1.0},

		{shapes.MakeSphere(),
		 patterns.Material{Pat:patterns.MakeRGB(1, 1, 1), Ambient:0.1, Diffuse:0.9, Specular:0.9,
		 				   Reflective:0, Transparency:1.0, RefractiveIndex:1.5},
		 primitives.MakeIdentityMatrix(4),
		 primitives.Ray{Origin:primitives.MakePoint(0, 0, 0), Direction:primitives.MakeVector(0, 1, 0)},
		 0.04},

		{shapes.MakeSphere(),
		 patterns.Material{Pat:patterns.MakeRGB(1, 1, 1), Ambient:0.1, Diffuse:0.9, Specular:0.9,
						   Reflective:0, Transparency:1.0, RefractiveIndex:1.5},
		 primitives.MakeIdentityMatrix(4),
		 primitives.Ray{Origin:primitives.MakePoint(0, 0.99, -2), Direction:primitives.MakeVector(0, 0, 1)},
		 0.4888143830387389},
	}
	for _, table := range tables {
		world := components.World{}
		table.shape.SetMaterial(table.mat)
		table.shape.SetTransform(table.transform)
		world.AddObject(table.shape)
		xs := world.Intersect(table.ray)
		intersection, _ := xs.Hit()
		comps := components.PrepareComputations(intersection, table.ray, xs)
		reflectance := comps.Schlick()
		if math.Abs(reflectance - table.reflectance) > primitives.EPSILON {
			t.Errorf("Expected %v, got %v", table.reflectance, reflectance)
		}
	}
}

func BenchmarkNoBoundingBoxes(b *testing.B) {
	world := components.World{}
	light := components.PointLight{Intensity: patterns.MakeRGB(1, 1, 1),
		Position: primitives.MakePoint(-30, 60, -30)}
	world.AddLight(light)
	for x := 0.0; x < 16; x++ {
		for y := 0.0; y < 16; y++ {
			for z := 0.0; z < 16; z++ {
				sphere := shapes.MakeSphere()
				sphere.SetTransform(primitives.Translation(x * 4, y * 4, z * 4))
				world.AddObject(sphere)
			}
		}
	}
	ray := primitives.Ray{Origin:primitives.MakePoint(4, 4, -5), Direction:primitives.MakeVector(0, 0, 1)}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = world.ColorAt(ray, 5)
	}
}

func Benchmark8BoundingBoxes(b *testing.B) {
	world := components.World{}
	light := components.PointLight{Intensity: patterns.MakeRGB(1, 1, 1),
		Position: primitives.MakePoint(-30, 60, -30)}
	world.AddLight(light)
	var groups [8]*shapes.Group
	for index := 0; index < 8; index++ {
		groups[index] = shapes.MakeGroup()
		world.AddObject(groups[index])
	}
	for x := 0.0; x < 16; x++ {
		for y := 0.0; y < 16; y++ {
			for z := 0.0; z < 16; z++ {
				sphere := shapes.MakeSphere()
				sphere.SetTransform(primitives.Translation(x * 4, y * 4, z * 4))
				group_no := (int(x / 8) * 4) + (int(y / 8) * 2) + (int(z / 8))
				groups[group_no].AddShape(sphere)
			}
		}
	}
	ray := primitives.Ray{Origin:primitives.MakePoint(4, 4, -5), Direction:primitives.MakeVector(0, 0, 1)}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = world.ColorAt(ray, 5)
	}
}

func Benchmark64BoundingBoxes(b * testing.B) {
	world := components.World{}
	light := components.PointLight{Intensity: patterns.MakeRGB(1, 1, 1),
		Position: primitives.MakePoint(-30, 60, -30)}
	world.AddLight(light)
	var groups [64]*shapes.Group
	for index := 0; index < 64; index++ {
		groups[index] = shapes.MakeGroup()
		world.AddObject(groups[index])
	}
	for x := 0.0; x < 16; x++ {
		for y := 0.0; y < 16; y++ {
			for z := 0.0; z < 16; z++ {
				sphere := shapes.MakeSphere()
				sphere.SetTransform(primitives.Translation(x * 4, y * 4, z * 4))
				group_no := (int(x / 4) * 16) + (int(y / 4) * 4) + (int(z / 4))
				groups[group_no].AddShape(sphere)
			}
		}
	}
	ray := primitives.Ray{Origin:primitives.MakePoint(4, 4, -5), Direction:primitives.MakeVector(0, 0, 1)}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = world.ColorAt(ray, 5)
	}
}
