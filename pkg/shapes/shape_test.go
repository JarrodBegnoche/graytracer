package shapes_test

import (
	"testing"
	"github.com/factorion/graytracer/pkg/primitives"
	"github.com/factorion/graytracer/pkg/patterns"
	"github.com/factorion/graytracer/pkg/shapes"
)

func TestSliceEquals(t *testing.T) {
	tables := []struct {
		a, b []float64
		equals bool
	}{
		{[]float64{0.0, 1.0}, []float64{0.0, 1.0}, true},
		{[]float64{0.0, 1.0}, []float64{1.0, 2.0}, false},
		{[]float64{0.0, 1.0, 2.0}, []float64{0.0, 1.0}, false},
		{[]float64{0.0, 1.0}, []float64{0.0, 1.0000000001}, true},
	}
	for _, table := range tables {
		equals := shapes.SliceEquals(table.a, table.b)
		if equals != table.equals {
			t.Errorf("Slice %v and %v returned %v as equals", table.a, table.b, equals)
		}
	}
}

func TestSphereTransform(t *testing.T) {
	tables := []struct {
		s *shapes.Sphere
		transform primitives.Matrix
	}{
		{shapes.MakeSphere(), primitives.Scaling(2, 2, 2)},
		{shapes.MakeSphere(), primitives.Translation(5, 0, 0)},
	}
	for _, table := range tables {
		table.s.SetTransform(table.transform)
		if !table.s.Transform().Equals(table.transform) {
			t.Errorf("Expected %v, got %v", table.transform, table.s.Transform())
		}
	}
}

func TestSphereMaterial(t *testing.T) {
	tables := []struct {
		s *shapes.Sphere
		mat patterns.Material
	}{
		{shapes.MakeSphere(), patterns.Material{Pat:patterns.MakeRGB(1, 0.9, 0.8),
										 Ambient:0.1, Diffuse:0.7, Specular:0.6, Shininess:150}},
	}
	for _, table := range tables {
		table.s.SetMaterial(table.mat)
		if table.s.Material() != table.mat {
			t.Errorf("Expected %v, got %v", table.mat, table.s.Material())
		}
	}
}
