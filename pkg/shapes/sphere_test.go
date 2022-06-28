package shapes_test

import (
	"math"
	"testing"

	"github.com/factorion/graytracer/pkg/primitives"
	"github.com/factorion/graytracer/pkg/shapes"
)

func TestSphereGetBounds(t *testing.T) {
	tables := []struct {
		sphere    *shapes.Sphere
		transform primitives.Matrix
		min, max  primitives.PV
	}{
		{shapes.MakeSphere(),
			primitives.MakeIdentityMatrix(4),
			primitives.MakePoint(-1, -1, -1), primitives.MakePoint(1, 1, 1)},

		{shapes.MakeSphere(),
			primitives.Scaling(2, 3, 4),
			primitives.MakePoint(-2, -3, -4), primitives.MakePoint(2, 3, 4)},
	}
	for _, table := range tables {
		table.sphere.SetTransform(table.transform)
		bounds := table.sphere.GetBounds()
		if !bounds.Min.Equals(table.min) {
			t.Errorf("Expected Minimum %v, got %v", table.min, bounds.Min)
		}
		if !bounds.Max.Equals(table.max) {
			t.Errorf("Expected Maximum %v, got %v", table.max, bounds.Max)
		}
	}
}

func TestSphereIntersection(t *testing.T) {
	tables := []struct {
		s         *shapes.Sphere
		r         primitives.Ray
		transform primitives.Matrix
		hits      []float64
	}{
		{shapes.MakeSphere(),
			primitives.Ray{Origin: primitives.MakePoint(0, 1, -5), Direction: primitives.MakeVector(0, 0, 1)},
			primitives.MakeIdentityMatrix(4), []float64{5}},

		{shapes.MakeSphere(),
			primitives.Ray{Origin: primitives.MakePoint(0, 2, -5), Direction: primitives.MakeVector(0, 0, 1)},
			primitives.MakeIdentityMatrix(4), []float64{}},

		{shapes.MakeSphere(),
			primitives.Ray{Origin: primitives.MakePoint(0, 0, 5), Direction: primitives.MakeVector(0, 0, 1)},
			primitives.MakeIdentityMatrix(4), []float64{-6, -4}},

		{shapes.MakeSphere(),
			primitives.Ray{Origin: primitives.MakePoint(0, 0, -5), Direction: primitives.MakeVector(0, 0, 1)},
			primitives.Scaling(2, 2, 2), []float64{3, 7}},

		{shapes.MakeSphere(),
			primitives.Ray{Origin: primitives.MakePoint(0, 0, -5), Direction: primitives.MakeVector(0, 0, 1)},
			primitives.Translation(5, 0, 0), []float64{}},
	}
	for _, table := range tables {
		table.s.SetTransform(table.transform)
		hits := table.s.Intersect(table.r)
		if !shapes.IntersectEquals(hits, table.hits) {
			t.Errorf("Expected hit %v, got %v", table.hits, hits)
		}
	}
}

func BenchmarkSphereIntersection(b *testing.B) {
	sphere := shapes.MakeSphere()
	sphere.SetTransform(primitives.Scaling(0.5, 0.5, 0.5))
	ray := primitives.Ray{Origin: primitives.MakePoint(0, 0, -2), Direction: primitives.MakeVector(0, 0, 1)}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sphere.Intersect(ray)
	}
}

func TestSphereNormal(t *testing.T) {
	tables := []struct {
		s             *shapes.Sphere
		transform     primitives.Matrix
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
		normal := table.s.Normal(table.point, 0.0, 0.0)
		if !normal.Equals(table.normal) {
			t.Errorf("Expected %v, got %v", table.normal, normal)
		}
	}
}

func BenchmarkSphereNormal(b *testing.B) {
	sphere := shapes.MakeSphere()
	sphere.SetTransform(primitives.Scaling(0.5, 0.5, 0.5))
	point := primitives.MakePoint(0, 0, 0.5)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sphere.Normal(point, 0.0, 0.0)
	}
}
