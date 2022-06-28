package shapes_test

import (
	"math"
	"testing"

	"github.com/factorion/graytracer/pkg/primitives"
	"github.com/factorion/graytracer/pkg/shapes"
)

func TestPlaneGetBounds(t *testing.T) {
	tables := []struct {
		plane     *shapes.Plane
		transform primitives.Matrix
		min, max  primitives.PV
	}{
		{shapes.MakePlane(),
			primitives.MakeIdentityMatrix(4),
			primitives.MakePoint(math.Inf(-1), -primitives.EPSILON, math.Inf(-1)),
			primitives.MakePoint(math.Inf(1), primitives.EPSILON, math.Inf(1))},

		{shapes.MakePlane(),
			primitives.RotationX(math.Pi / 2),
			primitives.MakePoint(math.Inf(-1), math.Inf(-1), -primitives.EPSILON),
			primitives.MakePoint(math.Inf(1), math.Inf(1), primitives.EPSILON)},
	}
	for _, table := range tables {
		table.plane.SetTransform(table.transform)
		bounds := table.plane.GetBounds()
		if !bounds.Min.Equals(table.min) {
			t.Errorf("Expected Minimum %v, got %v", table.min, bounds.Min)
		}
		if !bounds.Max.Equals(table.max) {
			t.Errorf("Expected Maximum %v, got %v", table.max, bounds.Max)
		}
	}
}

func TestPlaneIntersection(t *testing.T) {
	tables := []struct {
		p         *shapes.Plane
		r         primitives.Ray
		transform primitives.Matrix
		hits      []float64
	}{
		{shapes.MakePlane(),
			primitives.Ray{Origin: primitives.MakePoint(0, 10, 0), Direction: primitives.MakeVector(0, 0, 1)},
			primitives.MakeIdentityMatrix(4), []float64{}},

		{shapes.MakePlane(),
			primitives.Ray{Origin: primitives.MakePoint(0, 0, 0), Direction: primitives.MakeVector(0, 0, 1)},
			primitives.MakeIdentityMatrix(4), []float64{}},

		{shapes.MakePlane(),
			primitives.Ray{Origin: primitives.MakePoint(0, 1, 0), Direction: primitives.MakeVector(0, -1, 0)},
			primitives.MakeIdentityMatrix(4), []float64{1}},

		{shapes.MakePlane(),
			primitives.Ray{Origin: primitives.MakePoint(0, -1, 0), Direction: primitives.MakeVector(0, 1, 0)},
			primitives.MakeIdentityMatrix(4), []float64{1}},
	}
	for _, table := range tables {
		table.p.SetTransform(table.transform)
		hits := table.p.Intersect(table.r)
		if !shapes.IntersectEquals(hits, table.hits) {
			t.Errorf("Expected hit %v, got %v", table.hits, hits)
		}
	}
}

func BenchmarkPlaneIntersection(b *testing.B) {
	plane := shapes.MakePlane()
	plane.SetTransform(primitives.RotationX(math.Pi / 2))
	ray := primitives.Ray{Origin: primitives.MakePoint(0, 0, -2), Direction: primitives.MakeVector(0, 0, 1)}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		plane.Intersect(ray)
	}
}

func TestPlaneNormal(t *testing.T) {
	tables := []struct {
		p             *shapes.Plane
		point, normal primitives.PV
	}{
		{shapes.MakePlane(),
			primitives.MakePoint(0, 0, 0),
			primitives.MakeVector(0, 1, 0)},

		{shapes.MakePlane(),
			primitives.MakePoint(10, 0, -10),
			primitives.MakeVector(0, 1, 0)},

		{shapes.MakePlane(),
			primitives.MakePoint(-5, 0, 150),
			primitives.MakeVector(0, 1, 0)},
	}
	for _, table := range tables {
		normal := table.p.Normal(table.point, 0.0, 0.0)
		if !normal.Equals(table.normal) {
			t.Errorf("Expected %v, got %v", table.normal, normal)
		}
	}
}

func BenchmarkPlaneNormal(b *testing.B) {
	plane := shapes.MakePlane()
	plane.SetTransform(primitives.RotationX(math.Pi / 2))
	point := primitives.MakePoint(1, 1, 0)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		plane.Normal(point, 0.0, 0.0)
	}
}
