package shapes_test

import (
	"testing"
	"github.com/factorion/graytracer/pkg/primitives"
	"github.com/factorion/graytracer/pkg/shapes"
)

func TestMakeTriangle(t *testing.T) {
	tables := []struct{
		point1, point2, point3, edge1, edge2, normal primitives.PV
	}{
		{primitives.MakePoint(0, 1, 0), primitives.MakePoint(-1, 0, 0), primitives.MakePoint(1, 0, 0),
		 primitives.MakeVector(-1, -1, 0), primitives.MakeVector(1, -1, 0), primitives.MakeVector(0, 0, -1)},
	}
	for _, table := range tables {
		triangle := shapes.MakeTriangle(table.point1, table.point2, table.point3)
		if !triangle.Edge1.Equals(table.edge1) {
			t.Errorf("Expected Edge1 %v, got %v", table.edge1, triangle.Edge1)
		}
		if !triangle.Edge2.Equals(table.edge2) {
			t.Errorf("Expected Edge2 %v, got %v", table.edge1, triangle.Edge2)
		}
		if !triangle.Norm.Equals(table.normal) {
			t.Errorf("Expected Normal %v, got %v", table.normal, triangle.Norm)
		}
	}
}

func TestTriangleGetBounds(t *testing.T) {
	tables := []struct{
		t *shapes.Triangle
		transform primitives.Matrix
		min, max primitives.PV
	}{
		{shapes.MakeTriangle(primitives.MakePoint(0, 1, 0), primitives.MakePoint(-1, 0, 0),
							 primitives.MakePoint(1, 0, 0)),
		 primitives.MakeIdentityMatrix(4),
		 primitives.MakePoint(-1, 0, 0), primitives.MakePoint(1, 1, 0)},
	}
	for _, table := range tables {
		table.t.SetTransform(table.transform)
		bounds := table.t.GetBounds()
		if !bounds.Min.Equals(table.min) {
			t.Errorf("Expected Minimum %v, got %v", table.min, bounds.Min)
		}
		if !bounds.Max.Equals(table.max) {
			t.Errorf("Expected Maximum %v, got %v", table.max, bounds.Max)
		}
	}
}

func TestTriangleIntersect(t *testing.T) {
	tables := []struct{
		t *shapes.Triangle
		transform primitives.Matrix
		r primitives.Ray
		hits []float64
	}{
		{shapes.MakeTriangle(primitives.MakePoint(0, 1, 0), primitives.MakePoint(-1, 0, 0),
							 primitives.MakePoint(1, 0, 0)),
		 primitives.MakeIdentityMatrix(4),
		 primitives.Ray{Origin:primitives.MakePoint(0, -1, -2), Direction:primitives.MakeVector(0, 1, 0)},
		 []float64{}},
		
		{shapes.MakeTriangle(primitives.MakePoint(0, 1, 0), primitives.MakePoint(-1, 0, 0),
							 primitives.MakePoint(1, 0, 0)),
		 primitives.MakeIdentityMatrix(4),
		 primitives.Ray{Origin:primitives.MakePoint(1, 1, -2), Direction:primitives.MakeVector(0, 0, 1)},
		 []float64{}},
		
		{shapes.MakeTriangle(primitives.MakePoint(0, 1, 0), primitives.MakePoint(-1, 0, 0),
							 primitives.MakePoint(1, 0, 0)),
		 primitives.MakeIdentityMatrix(4),
		 primitives.Ray{Origin:primitives.MakePoint(-1, 1, -2), Direction:primitives.MakeVector(0, 0, 1)},
		 []float64{}},
		
		{shapes.MakeTriangle(primitives.MakePoint(0, 1, 0), primitives.MakePoint(-1, 0, 0),
							 primitives.MakePoint(1, 0, 0)),
		 primitives.MakeIdentityMatrix(4),
		 primitives.Ray{Origin:primitives.MakePoint(0, -1, -2), Direction:primitives.MakeVector(0, 0, 1)},
		 []float64{}},
		
		{shapes.MakeTriangle(primitives.MakePoint(0, 1, 0), primitives.MakePoint(-1, 0, 0),
							 primitives.MakePoint(1, 0, 0)),
		 primitives.MakeIdentityMatrix(4),
		 primitives.Ray{Origin:primitives.MakePoint(0, 0.5, -2), Direction:primitives.MakeVector(0, 0, 1)},
		 []float64{2}},
	}
	for _, table := range tables {
		table.t.SetTransform(table.transform)
		hits := table.t.Intersect(table.r)
		if !shapes.IntersectEquals(hits, table.hits) {
			t.Errorf("Expected hit %v, got %v", table.hits, hits)
		}
	}
}

func BenchmarkTriangleIntersection(b *testing.B) {
	triangle := shapes.MakeTriangle(primitives.MakePoint(0, 1, 0), primitives.MakePoint(-1, 0, 0),
									primitives.MakePoint(1, 0, 0))
	ray := primitives.Ray{Origin:primitives.MakePoint(0, 0.5, -2), Direction:primitives.MakeVector(0, 0, 1)}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		triangle.Intersect(ray)
	}
}

func TestTriangleNormal(t *testing.T) {
	tables := []struct{
		t *shapes.Triangle
		transform primitives.Matrix
		point, normal primitives.PV
	}{
		{shapes.MakeTriangle(primitives.MakePoint(0, 1, 0), primitives.MakePoint(-1, 0, 0),
							 primitives.MakePoint(1, 0, 0)),
		 primitives.MakeIdentityMatrix(4), primitives.MakePoint(0, 0.5, 0), primitives.MakeVector(0, 0, -1)},
		
		{shapes.MakeTriangle(primitives.MakePoint(0, 1, 0), primitives.MakePoint(-1, 0, 0),
							 primitives.MakePoint(1, 0, 0)),
		 primitives.MakeIdentityMatrix(4), primitives.MakePoint(-0.5, 0.75, 0), primitives.MakeVector(0, 0, -1)},
		
		{shapes.MakeTriangle(primitives.MakePoint(0, 1, 0), primitives.MakePoint(-1, 0, 0),
							 primitives.MakePoint(1, 0, 0)),
		 primitives.MakeIdentityMatrix(4), primitives.MakePoint(0.5, 0.25, 0), primitives.MakeVector(0, 0, -1)},
	}
	for _, table := range tables {
		table.t.SetTransform(table.transform)
		normal := table.t.Normal(table.point)
		if !normal.Equals(table.normal) {
			t.Errorf("Expected normal %v, got %v", table.normal, normal)
		}
	}
}

func BenchmarkTriangleNormal(b *testing.B) {
	triangle := shapes.MakeTriangle(primitives.MakePoint(0, 1, 0), primitives.MakePoint(-1, 0, 0),
									primitives.MakePoint(1, 0, 0))
	point := primitives.MakePoint(0, 0.5, 0)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		triangle.Normal(point)
	}
}
