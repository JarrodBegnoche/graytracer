package shapes_test

import (
	//"math"
	"testing"
	"github.com/factorion/graytracer/pkg/primitives"
	"github.com/factorion/graytracer/pkg/shapes"
)

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
		if !shapes.SliceEquals(hits, table.hits) {
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
