package shapes_test

import (
	"math"
	"testing"
	"github.com/factorion/graytracer/pkg/primitives"
	"github.com/factorion/graytracer/pkg/shapes"
)

func TestSphereIntersection(t *testing.T) {
	tables := []struct {
		s *shapes.Sphere
		r primitives.Ray
		transform primitives.Matrix
		hits []float64
	}{
		{shapes.MakeSphere(),
		 primitives.Ray{Origin:primitives.MakePoint(0, 1, -5), Direction:primitives.MakeVector(0, 0, 1)},
		 primitives.MakeIdentityMatrix(4), []float64{5}},

		{shapes.MakeSphere(),
		 primitives.Ray{Origin:primitives.MakePoint(0, 2, -5), Direction:primitives.MakeVector(0, 0, 1)},
		 primitives.MakeIdentityMatrix(4), []float64{}},

		{shapes.MakeSphere(),
		 primitives.Ray{Origin:primitives.MakePoint(0, 0, 5), Direction:primitives.MakeVector(0, 0, 1)},
		 primitives.MakeIdentityMatrix(4), []float64{-6, -4}},

		{shapes.MakeSphere(),
		 primitives.Ray{Origin:primitives.MakePoint(0, 0, -5), Direction:primitives.MakeVector(0, 0, 1)},
		 primitives.Scaling(2, 2, 2), []float64{3, 7}},

		{shapes.MakeSphere(),
		 primitives.Ray{Origin:primitives.MakePoint(0, 0, -5), Direction:primitives.MakeVector(0, 0, 1)},
		 primitives.Translation(5, 0, 0), []float64{}},
	}
	for _, table := range tables {
		table.s.SetTransform(table.transform)
		hits := table.s.Intersect(table.r)
		if !shapes.SliceEquals(hits, table.hits) {
			t.Errorf("Expected hit %v, got %v", table.hits, hits)
		}
	}
}

func TestSphereNormal(t *testing.T) {
	tables := []struct {
		s *shapes.Sphere
		transform primitives.Matrix
		point, normal primitives.PV
	}{
		{shapes.MakeSphere(), primitives.Translation(0, 1, 0),
		 primitives.MakePoint(0, 1.7071067811865476, -0.7071067811865476),
		 primitives.MakeVector(0, 0.7071067811865476, -0.7071067811865476)},

		{shapes.MakeSphere(), primitives.Scaling(1.0, 0.5, 1.0).Multiply(primitives.RotationZ(math.Pi / 5.0)),
		 primitives.MakePoint(0, 0.7071067811865476, -0.7071067811865476),
		 primitives.MakeVector(0, 0.9701425001453319, -0.24253562503633294)},
	}
	for _, table := range tables {
		table.s.SetTransform(table.transform)
		normal := table.s.Normal(table.point)
		if !normal.Equals(table.normal) {
			t.Errorf("Expected %v, got %v", table.normal, normal)
		}
	}
}
