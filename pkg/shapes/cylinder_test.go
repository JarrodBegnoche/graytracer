package shapes_test

import (
	"testing"
	"github.com/factorion/graytracer/pkg/primitives"
	"github.com/factorion/graytracer/pkg/shapes"
)

func TestCylinderIntersection(t *testing.T) {
	tables := []struct {
		s *shapes.Cylinder
		r primitives.Ray
		transform primitives.Matrix
		hits []float64
	}{
		// Open intersections
		{shapes.MakeCylinder(false),
		 primitives.Ray{Origin:primitives.MakePoint(1, 0, 0), Direction:primitives.MakeVector(0, 1, 0)},
		 primitives.MakeIdentityMatrix(4), []float64{}},
		
		{shapes.MakeCylinder(false),
		 primitives.Ray{Origin:primitives.MakePoint(0, 0, -5), Direction:primitives.MakeVector(1, 1, 1)},
		 primitives.MakeIdentityMatrix(4), []float64{}},

		{shapes.MakeCylinder(false),
		 primitives.Ray{Origin:primitives.MakePoint(1, 0.5, -5), Direction:primitives.MakeVector(0, 0, 1)},
		 primitives.MakeIdentityMatrix(4), []float64{5, 5}},

		{shapes.MakeCylinder(false),
		 primitives.Ray{Origin:primitives.MakePoint(0, 0, -5), Direction:primitives.MakeVector(0, 0, 1)},
		 primitives.Translation(0, -0.5, 0), []float64{4, 6}},

		{shapes.MakeCylinder(false),
		 primitives.Ray{Origin:primitives.MakePoint(0.5, 0, -5), Direction:primitives.MakeVector(0.1, 1, 1)},
		 primitives.Translation(0, 4.5, 0), []float64{4.801980198019795, 5}},

		{shapes.MakeCylinder(false),
		 primitives.Ray{Origin:primitives.MakePoint(0, 1.5, 0), Direction:primitives.MakeVector(0.1, 1, 0)},
		 primitives.Scaling(0, 2, 0), []float64{}},

		{shapes.MakeCylinder(false),
		 primitives.Ray{Origin:primitives.MakePoint(0, 3, -5), Direction:primitives.MakeVector(0, 0, 1)},
		 primitives.Scaling(0, 2, 0), []float64{}},

		{shapes.MakeCylinder(false),
		 primitives.Ray{Origin:primitives.MakePoint(0, -1, -5), Direction:primitives.MakeVector(0, 0, 1)},
		 primitives.Scaling(0, 2, 0), []float64{}},

		{shapes.MakeCylinder(false),
		 primitives.Ray{Origin:primitives.MakePoint(0, 2, -5), Direction:primitives.MakeVector(0, 0, 1)},
		 primitives.Scaling(0, 2, 0), []float64{}},

		{shapes.MakeCylinder(false),
		 primitives.Ray{Origin:primitives.MakePoint(0, 0, -5), Direction:primitives.MakeVector(0, 0, 1)},
		 primitives.Scaling(0, 2, 0), []float64{}},

		{shapes.MakeCylinder(false),
		 primitives.Ray{Origin:primitives.MakePoint(0, 1.5, -2), Direction:primitives.MakeVector(0, 0, 1)},
		 primitives.Scaling(0, 2, 0), []float64{1, 3}},

		{shapes.MakeCylinder(true),
		 primitives.Ray{Origin:primitives.MakePoint(0, 0.5, -2), Direction:primitives.MakeVector(0, 0, 1)},
		 primitives.Scaling(0.5, 1, 0.5), []float64{1.5, 2.5}},

		// Closed cylinder intersections
		{shapes.MakeCylinder(true),
		 primitives.Ray{Origin:primitives.MakePoint(0, 3, 0), Direction:primitives.MakeVector(0, -1, 0)},
		 primitives.Scaling(0, 2, 0), []float64{3, 1}},
		
		{shapes.MakeCylinder(true),
		 primitives.Ray{Origin:primitives.MakePoint(0, 3, -2), Direction:primitives.MakeVector(0, -1, 2)},
		 primitives.Translation(0, 1, 0), []float64{1.5, 1}},
		
		{shapes.MakeCylinder(true),
		 primitives.Ray{Origin:primitives.MakePoint(0, 4, -2), Direction:primitives.MakeVector(0, -1, 1)},
		 primitives.Translation(0, 1, 0), []float64{3, 2}},
		
		{shapes.MakeCylinder(true),
		 primitives.Ray{Origin:primitives.MakePoint(0, 0, -2), Direction:primitives.MakeVector(0, 1, 2)},
		 primitives.Translation(0, 1, 0), []float64{1.5, 1}},
		
		{shapes.MakeCylinder(true),
		 primitives.Ray{Origin:primitives.MakePoint(0, -1, -2), Direction:primitives.MakeVector(0, 1, 1)},
		 primitives.Translation(0, 1, 0), []float64{2, 3}},
	}
	for _, table := range tables {
		table.s.SetTransform(table.transform)
		hits := table.s.Intersect(table.r)
		if !shapes.IntersectEquals(hits, table.hits) {
			t.Errorf("Expected hit %v, got %v", table.hits, hits)
		}
	}
}

func TestCylinderNormal(t *testing.T) {
	tables := []struct {
		c *shapes.Cylinder
		transform primitives.Matrix
		point, normal primitives.PV
	}{
		// Side normals
		{shapes.MakeCylinder(false), primitives.MakeIdentityMatrix(4),
		 primitives.MakePoint(1, 0, 0),
		 primitives.MakeVector(1, 0, 0)},

		{shapes.MakeCylinder(false), primitives.Scaling(0, 5, 0),
		 primitives.MakePoint(0, 5, -1),
		 primitives.MakeVector(0, 0, -1)},

		{shapes.MakeCylinder(false), primitives.Translation(0, -2.5, 0),
		 primitives.MakePoint(0, -2, 1),
		 primitives.MakeVector(0, 0, 1)},

		{shapes.MakeCylinder(false), primitives.Scaling(0, 2, 0),
		 primitives.MakePoint(-1, 1, 0),
		 primitives.MakeVector(-1, 0, 0)},

		// End cap normals
		{shapes.MakeCylinder(true), primitives.Translation(0, 1, 0),
		 primitives.MakePoint(0, 1, 0),
		 primitives.MakeVector(0, -1, 0)},

		{shapes.MakeCylinder(true), primitives.Translation(0, 1, 0),
		 primitives.MakePoint(0.5, 1, 0),
		 primitives.MakeVector(0, -1, 0)},

		{shapes.MakeCylinder(true), primitives.Translation(0, 1, 0),
		 primitives.MakePoint(0, 1, 0.5),
		 primitives.MakeVector(0, -1, 0)},
		
		{shapes.MakeCylinder(true), primitives.Scaling(0, 2, 0),
		 primitives.MakePoint(0, 2, 0),
		 primitives.MakeVector(0, 1, 0)},
 
		{shapes.MakeCylinder(true), primitives.Scaling(0, 2, 0),
		 primitives.MakePoint(0.5, 2, 0),
		 primitives.MakeVector(0, 1, 0)},
 
		{shapes.MakeCylinder(true), primitives.Scaling(0, 2, 0),
		 primitives.MakePoint(0, 2, 0.5),
		 primitives.MakeVector(0, 1, 0)},
	}
	for _, table := range tables {
		table.c.SetTransform(table.transform)
		normal := table.c.Normal(table.point)
		if !normal.Equals(table.normal) {
			t.Errorf("Expected %v, got %v", table.normal, normal)
		}
	}
}
