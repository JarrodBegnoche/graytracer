package shapes_test

import (
	//"math"
	"testing"
	"github.com/factorion/graytracer/pkg/primitives"
	"github.com/factorion/graytracer/pkg/shapes"
)

func TestCubeGetBounds(t *testing.T) {
	tables := []struct {
		cube *shapes.Cube
		transform primitives.Matrix
		min, max primitives.PV
	}{
		{shapes.MakeCube(),
		 primitives.MakeIdentityMatrix(4),
		 primitives.MakePoint(-1, -1, -1), primitives.MakePoint(1, 1, 1)},

		{shapes.MakeCube(),
		 primitives.Scaling(2, 3, 4),
		 primitives.MakePoint(-2, -3, -4), primitives.MakePoint(2, 3, 4)},
	}
	for _, table := range tables {
		table.cube.SetTransform(table.transform)
		bounds := table.cube.GetBounds()
		if !bounds.Min.Equals(table.min) {
			t.Errorf("Expected Minimum %v, got %v", table.min, bounds.Min)
		}
		if !bounds.Max.Equals(table.max) {
			t.Errorf("Expected Maximum %v, got %v", table.max, bounds.Max)
		}
	}
}

func TestCubeIntersection(t *testing.T) {
	tables := []struct {
		c *shapes.Cube
		r primitives.Ray
		transform primitives.Matrix
		hits []float64
	}{
		{shapes.MakeCube(),
		 primitives.Ray{Origin:primitives.MakePoint(2, 0, 1), Direction:primitives.MakeVector(0, 0, -1)},
		 primitives.MakeIdentityMatrix(4), []float64{}},

		{shapes.MakeCube(),
		 primitives.Ray{Origin:primitives.MakePoint(5, 0.5, 0), Direction:primitives.MakeVector(-1, 0, 0)},
		 primitives.MakeIdentityMatrix(4), []float64{4, 6}},
		
		{shapes.MakeCube(),
		 primitives.Ray{Origin:primitives.MakePoint(-5, 0.5, 0), Direction:primitives.MakeVector(1, 0, 0)},
		 primitives.MakeIdentityMatrix(4), []float64{4, 6}},

		{shapes.MakeCube(),
		 primitives.Ray{Origin:primitives.MakePoint(0.5, 5, 0), Direction:primitives.MakeVector(0, -1, 0)},
		 primitives.MakeIdentityMatrix(4), []float64{4, 6}},

		{shapes.MakeCube(),
		 primitives.Ray{Origin:primitives.MakePoint(0.5, -5, 0), Direction:primitives.MakeVector(0, 1, 0)},
		 primitives.MakeIdentityMatrix(4), []float64{4, 6}},

		{shapes.MakeCube(),
		 primitives.Ray{Origin:primitives.MakePoint(0.5, 0, 5), Direction:primitives.MakeVector(0, 0, -1)},
		 primitives.MakeIdentityMatrix(4), []float64{4, 6}},

		{shapes.MakeCube(),
		 primitives.Ray{Origin:primitives.MakePoint(0.5, 0, -5), Direction:primitives.MakeVector(0, 0, 1)},
		 primitives.MakeIdentityMatrix(4), []float64{4, 6}},

		{shapes.MakeCube(),
		 primitives.Ray{Origin:primitives.MakePoint(0, 0.5, 0), Direction:primitives.MakeVector(0, 0, 1)},
		 primitives.MakeIdentityMatrix(4), []float64{-1, 1}},
	}
	for _, table := range tables {
		table.c.SetTransform(table.transform)
		hits := table.c.Intersect(table.r)
		if !shapes.IntersectEquals(hits, table.hits) {
			t.Errorf("Expected hit %v, got %v", table.hits, hits)
		}
	}
}

func TestCubeNormal(t *testing.T) {
	tables := []struct {
		c *shapes.Cube
		transform primitives.Matrix
		point, normal primitives.PV
	}{
		{shapes.MakeCube(), primitives.MakeIdentityMatrix(4),
		 primitives.MakePoint(1, 0.5, -0.8), primitives.MakeVector(1, 0, 0)},

		{shapes.MakeCube(), primitives.MakeIdentityMatrix(4),
		 primitives.MakePoint(-1, -0.2, 0.9), primitives.MakeVector(-1, 0, 0)},

		{shapes.MakeCube(), primitives.MakeIdentityMatrix(4),
		 primitives.MakePoint(-0.4, 1, -0.1), primitives.MakeVector(0, 1, 0)},

		{shapes.MakeCube(), primitives.MakeIdentityMatrix(4),
		 primitives.MakePoint(0.3, -1, -0.7), primitives.MakeVector(0, -1, 0)},
		
		{shapes.MakeCube(), primitives.MakeIdentityMatrix(4),
		 primitives.MakePoint(-0.6, 0.3, 1), primitives.MakeVector(0, 0, 1)},
   
		{shapes.MakeCube(), primitives.MakeIdentityMatrix(4),
		 primitives.MakePoint(0.4, 0.4, -1), primitives.MakeVector(0, 0, -1)},
   
		{shapes.MakeCube(), primitives.MakeIdentityMatrix(4),
		 primitives.MakePoint(1, 1, 1), primitives.MakeVector(1, 0, 0)},
   
		{shapes.MakeCube(), primitives.MakeIdentityMatrix(4),
		 primitives.MakePoint(-1, -1, -1), primitives.MakeVector(-1, 0, 0)},
	}
	for _, table := range tables {
		table.c.SetTransform(table.transform)
		normal := table.c.Normal(table.point)
		if !normal.Equals(table.normal) {
			t.Errorf("Expected %v, got %v", table.normal, normal)
		}
	}
}

func BenchmarkCubeIntersection(b *testing.B) {
	cube := shapes.MakeCube()
	cube.SetTransform(primitives.Scaling(0.5, 0.5, 0.5))
	ray := primitives.Ray{Origin:primitives.MakePoint(0, 0, -2), Direction:primitives.MakeVector(0, 0, 1)}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cube.Intersect(ray)
	}
}

func BenchmarkCubeNormal(b *testing.B) {
	cube := shapes.MakeCube()
	cube.SetTransform(primitives.Scaling(0.5, 0.5, 0.5))
	point := primitives.MakePoint(0, 0, 0.5)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cube.Normal(point)
	}
}
