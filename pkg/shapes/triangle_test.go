package shapes_test

import (
	"testing"

	"github.com/factorion/graytracer/pkg/primitives"
	"github.com/factorion/graytracer/pkg/shapes"
)

func TestMakeTriangle(t *testing.T) {
	tables := []struct {
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
			t.Errorf("Expected Edge2 %v, got %v", table.edge2, triangle.Edge2)
		}
		if !triangle.Normal1.Equals(table.normal) {
			t.Errorf("Expected Normal %v, got %v", table.normal, triangle.Normal1)
		}
	}
}

func TestTriangleGetBounds(t *testing.T) {
	tables := []struct {
		t         *shapes.Triangle
		transform primitives.Matrix
		min, max  primitives.PV
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
	tables := []struct {
		t         *shapes.Triangle
		transform primitives.Matrix
		r         primitives.Ray
		hits      []float64
	}{
		{shapes.MakeTriangle(primitives.MakePoint(0, 1, 0), primitives.MakePoint(-1, 0, 0),
			primitives.MakePoint(1, 0, 0)),
			primitives.MakeIdentityMatrix(4),
			primitives.Ray{Origin: primitives.MakePoint(0, -1, -2), Direction: primitives.MakeVector(0, 1, 0)},
			[]float64{}},

		{shapes.MakeTriangle(primitives.MakePoint(0, 1, 0), primitives.MakePoint(-1, 0, 0),
			primitives.MakePoint(1, 0, 0)),
			primitives.MakeIdentityMatrix(4),
			primitives.Ray{Origin: primitives.MakePoint(1, 1, -2), Direction: primitives.MakeVector(0, 0, 1)},
			[]float64{}},

		{shapes.MakeTriangle(primitives.MakePoint(0, 1, 0), primitives.MakePoint(-1, 0, 0),
			primitives.MakePoint(1, 0, 0)),
			primitives.MakeIdentityMatrix(4),
			primitives.Ray{Origin: primitives.MakePoint(-1, 1, -2), Direction: primitives.MakeVector(0, 0, 1)},
			[]float64{}},

		{shapes.MakeTriangle(primitives.MakePoint(0, 1, 0), primitives.MakePoint(-1, 0, 0),
			primitives.MakePoint(1, 0, 0)),
			primitives.MakeIdentityMatrix(4),
			primitives.Ray{Origin: primitives.MakePoint(0, -1, -2), Direction: primitives.MakeVector(0, 0, 1)},
			[]float64{}},

		{shapes.MakeTriangle(primitives.MakePoint(0, 1, 0), primitives.MakePoint(-1, 0, 0),
			primitives.MakePoint(1, 0, 0)),
			primitives.MakeIdentityMatrix(4),
			primitives.Ray{Origin: primitives.MakePoint(0, 0.5, -2), Direction: primitives.MakeVector(0, 0, 1)},
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
	ray := primitives.Ray{Origin: primitives.MakePoint(0, 0.5, -2), Direction: primitives.MakeVector(0, 0, 1)}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		triangle.Intersect(ray)
	}
}

func TestTriangleNormal(t *testing.T) {
	tables := []struct {
		t             *shapes.Triangle
		transform     primitives.Matrix
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
		normal := table.t.Normal(table.point, 0.0, 0.0)
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
		triangle.Normal(point, 0.0, 0.0)
	}
}

func TestMakeSmoothTriangle(t *testing.T) {
	tables := []struct {
		point1, point2, point3, normal1, normal2, normal3, edge1, edge2 primitives.PV
	}{
		{primitives.MakePoint(0, 1, 0), primitives.MakePoint(-1, 0, 0), primitives.MakePoint(1, 0, 0),
			primitives.MakePoint(0, 1, 0), primitives.MakePoint(-1, 0, 0), primitives.MakePoint(1, 0, 0),
			primitives.MakeVector(-1, -1, 0), primitives.MakeVector(1, -1, 0)},
	}
	for _, table := range tables {
		smooth_triangle := shapes.MakeSmoothTriangle(table.point1, table.point2, table.point3,
			table.normal1, table.normal2, table.normal3)
		if !smooth_triangle.Point1.Equals(table.point1) {
			t.Errorf("Expected Point1 %v, got %v", table.point1, smooth_triangle.Point1)
		}
		if !smooth_triangle.Point2.Equals(table.point2) {
			t.Errorf("Expected Point2 %v, got %v", table.point2, smooth_triangle.Point2)
		}
		if !smooth_triangle.Point3.Equals(table.point3) {
			t.Errorf("Expected Point3 %v, got %v", table.point3, smooth_triangle.Point3)
		}
		if !smooth_triangle.Normal1.Equals(table.normal1) {
			t.Errorf("Expected Normal1 %v, got %v", table.normal1, smooth_triangle.Normal1)
		}
		if !smooth_triangle.Normal2.Equals(table.normal2) {
			t.Errorf("Expected Normal2 %v, got %v", table.normal2, smooth_triangle.Normal2)
		}
		if !smooth_triangle.Normal3.Equals(table.normal3) {
			t.Errorf("Expected Normal3 %v, got %v", table.normal3, smooth_triangle.Normal3)
		}
		if !smooth_triangle.Edge1.Equals(table.edge1) {
			t.Errorf("Expected Edge1 %v, got %v", table.edge1, smooth_triangle.Edge1)
		}
		if !smooth_triangle.Edge2.Equals(table.edge2) {
			t.Errorf("Expected Edge2 %v, got %v", table.edge2, smooth_triangle.Edge2)
		}
	}
}

func TestSmoothTriangleIntersect(t *testing.T) {
	tables := []struct {
		t         *shapes.Triangle
		transform primitives.Matrix
		r         primitives.Ray
		u, v      float64
	}{
		{shapes.MakeSmoothTriangle(primitives.MakePoint(0, 1, 0), primitives.MakePoint(-1, 0, 0), primitives.MakePoint(1, 0, 0),
			primitives.MakePoint(0, 1, 0), primitives.MakePoint(-1, 0, 0), primitives.MakePoint(1, 0, 0)),
			primitives.MakeIdentityMatrix(4),
			primitives.Ray{Origin: primitives.MakePoint(-0.2, 0.3, -2), Direction: primitives.MakeVector(0, 0, 1)},
			0.44999999999999996, 0.24999999999999997},
	}
	for _, table := range tables {
		table.t.SetTransform(table.transform)
		hits := table.t.Intersect(table.r)
		if hits.Len() != 1 {
			t.Errorf("Expected 1 hit, got %v", hits.Len())
		}
		if hits[0].U != table.u {
			t.Errorf("Expected U %v, got %v", table.u, hits[0].U)
		}
		if hits[0].V != table.v {
			t.Errorf("Expected V %v, got %v", table.v, hits[0].V)
		}
	}
}

func TestSmoothTriangleNormal(t *testing.T) {
	tables := []struct {
		t      *shapes.Triangle
		u, v   float64
		normal primitives.PV
	}{
		{shapes.MakeSmoothTriangle(primitives.MakePoint(0, 1, 0), primitives.MakePoint(-1, 0, 0), primitives.MakePoint(1, 0, 0),
			primitives.MakePoint(0, 1, 0), primitives.MakePoint(-1, 0, 0), primitives.MakePoint(1, 0, 0)),
			0.44999999999999996, 0.24999999999999997,
			primitives.MakeVector(-0.554700196225229, 0.8320502943378437, 0)},
	}
	for _, table := range tables {
		i := shapes.Intersection{1, table.t, table.u, table.v}
		normal := table.t.Normal(primitives.MakePoint(0, 0, 0), i.U, i.V)
		if !normal.Equals(table.normal) {
			t.Errorf("Expected normal %v, got %v", table.normal, normal)
		}
	}
}