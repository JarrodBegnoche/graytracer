package shapes_test

import (
	"testing"
	"github.com/factorion/graytracer/pkg/primitives"
	"github.com/factorion/graytracer/pkg/shapes"
)

func TestConeGetBounds(t *testing.T) {
	tables := []struct {
		cone *shapes.Cone
		transform primitives.Matrix
		min, max primitives.PV
	}{
		{shapes.MakeCone(false),
		 primitives.MakeIdentityMatrix(4),
		 primitives.MakePoint(-1, -1, -1), primitives.MakePoint(1, 0, 1)},

		{shapes.MakeCone(false),
		 primitives.Scaling(2, 2, 2),
		 primitives.MakePoint(-2, -2, -2), primitives.MakePoint(2, 0, 2)},
	}
	for _, table := range tables {
		table.cone.SetTransform(table.transform)
		bounds := table.cone.GetBounds()
		if !bounds.Min.Equals(table.min) {
			t.Errorf("Expected Minimum %v, got %v", table.min, bounds.Min)
		}
		if !bounds.Max.Equals(table.max) {
			t.Errorf("Expected Maximum %v, got %v", table.max, bounds.Max)
		}
	}
}

func TestConeIntersection(t *testing.T) {
	tables := []struct {
		s *shapes.Cone
		r primitives.Ray
		transform primitives.Matrix
		hits []float64
	}{
		// Open intersections
		{shapes.MakeCone(false),
		 primitives.Ray{Origin:primitives.MakePoint(0, -0.5, -5), Direction:primitives.MakeVector(0, 0, 1)},
		 primitives.MakeIdentityMatrix(4), []float64{4.5, 5.5}},
		
		{shapes.MakeCone(false),
		 primitives.Ray{Origin:primitives.MakePoint(0, 0, -1), Direction:primitives.MakeVector(0, -1, 1)},
		 primitives.MakeIdentityMatrix(4), []float64{0.25}},

		{shapes.MakeCone(false),
		 primitives.Ray{Origin:primitives.MakePoint(1, 1, -2), Direction:primitives.MakeVector(-0.5, -1, 1)},
		 primitives.MakeIdentityMatrix(4), []float64{1.5278640450004204}},

		{shapes.MakeCone(true),
		 primitives.Ray{Origin:primitives.MakePoint(0, -1, -0.25), Direction:primitives.MakeVector(0, 1, 0)},
		 primitives.Translation(0, 1, 0), []float64{1.75, 1}},
	}
	for _, table := range tables {
		table.s.SetTransform(table.transform)
		hits := table.s.Intersect(table.r)
		if !shapes.IntersectEquals(hits, table.hits) {
			t.Errorf("Expected hit %v, got %v", table.hits, hits)
		}
	}
}

func TestConeNormal(t *testing.T) {
	tables := []struct {
		c *shapes.Cone
		transform primitives.Matrix
		point, normal primitives.PV
	}{
		{shapes.MakeCone(true), primitives.Translation(0, 1, 0),
		 primitives.MakePoint(0, 0, 0),
		 primitives.MakeVector(0, -1, 0)},

		{shapes.MakeCone(false), primitives.Translation(0, 1, 0),
		 primitives.MakePoint(0.5, 0.5, 0),
		 primitives.MakeVector(0.7071067811865475, 0.7071067811865475, 0)},
	}
	for _, table := range tables {
		table.c.SetTransform(table.transform)
		normal := table.c.Normal(table.point)
		if !normal.Equals(table.normal) {
			t.Errorf("Expected %v, got %v", table.normal, normal)
		}
	}
}

func BenchmarkConeIntersection(b *testing.B) {
	cone := shapes.MakeCone(true)
	cone.SetTransform(primitives.Translation(0, 1, 0))
	ray := primitives.Ray{Origin:primitives.MakePoint(0, 0.5, -2), Direction:primitives.MakeVector(0, 0, 1)}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cone.Intersect(ray)
	}
}

func BenchmarkConeNormal(b *testing.B) {
	cone := shapes.MakeCone(true)
	cone.SetTransform(primitives.Translation(0, 1, 0))
	point := primitives.MakePoint(0, 0.5, 0.5)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cone.Normal(point)
	}
}
