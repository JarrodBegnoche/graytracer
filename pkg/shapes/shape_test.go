package shapes_test

import (
	"math"
	"testing"
	"github.com/factorion/graytracer/pkg/primitives"
	"github.com/factorion/graytracer/pkg/patterns"
	"github.com/factorion/graytracer/pkg/shapes"
)

type full_shape struct {
	shape shapes.Shape
	transform primitives.Matrix
}

func TestSphereTransform(t *testing.T) {
	tables := []struct {
		fs full_shape
	}{
		{full_shape{shape:shapes.MakeSphere(), transform:primitives.Scaling(2, 2, 2)}},

		{full_shape{shape:shapes.MakeSphere(), transform:primitives.Translation(5, 0, 0)}},
	}
	for _, table := range tables {
		table.fs.shape.SetTransform(table.fs.transform)
		if !table.fs.shape.Transform().Equals(table.fs.transform) {
			t.Errorf("Expected %v, got %v", table.fs.transform, table.fs.shape.Transform())
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

func TestWorldToObjectPV(t * testing.T) {
	tables := []struct {
		outer, inner *shapes.Group
		outer_tr, inner_tr primitives.Matrix
		fs full_shape
		point, result primitives.PV
	}{
		{shapes.MakeGroup(), shapes.MakeGroup(),
		 primitives.RotationY(math.Pi / 2), primitives.Scaling(2, 2, 2),
		 full_shape{shape:shapes.MakeSphere(), transform:primitives.Translation(5, 0, 0)},
		 primitives.MakePoint(-2, 0, -10), primitives.MakePoint(0, 0, -1)},
	}
	for _, table := range tables {
		table.outer.SetTransform(table.outer_tr)
		table.inner.SetTransform(table.inner_tr)
		table.fs.shape.SetTransform(table.fs.transform)
		table.outer.AddShape(table.inner)
		table.inner.AddShape(table.fs.shape)
		point := table.fs.shape.WorldToObjectPV(table.point)
		if !point.Equals(table.result) {
			t.Errorf("Expected %v, got %v", table.result, point)
		}
	}
}

func TestObjectToWorldPV(t * testing.T) {
	tables := []struct {
		outer, inner *shapes.Group
		outer_tr, inner_tr primitives.Matrix
		fs full_shape
		vector, result primitives.PV
	}{
		{shapes.MakeGroup(), shapes.MakeGroup(),
		 primitives.RotationY(math.Pi / 2), primitives.Scaling(1, 2, 3),
		 full_shape{shape:shapes.MakeSphere(), transform:primitives.Translation(5, 0, 0)},
		 primitives.MakeVector(math.Sqrt(3) / 3, math.Sqrt(3) / 3, math.Sqrt(3) / 3),
		 primitives.MakeVector(0.19245008972987526, 0.28867513459481287, -0.5773502691896257)},
	}
	for _, table := range tables {
		table.outer.SetTransform(table.outer_tr)
		table.inner.SetTransform(table.inner_tr)
		table.fs.shape.SetTransform(table.fs.transform)
		table.outer.AddShape(table.inner)
		table.inner.AddShape(table.fs.shape)
		vector := table.fs.shape.ObjectToWorldPV(table.vector)
		vector.W = 0
		if !vector.Equals(table.result) {
			t.Errorf("Expected %v, got %v", table.result, vector)
		}
	}
}
