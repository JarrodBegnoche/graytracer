package shapes_test

import (
	"testing"
	"github.com/factorion/graytracer/pkg/primitives"
	"github.com/factorion/graytracer/pkg/shapes"
)

func TestMinMax(t *testing.T) {
	tables := []struct{
		floats []float64
		min, max float64
	}{
		{[]float64{1.0, 2.0, 3.0, 0.5}, 0.5, 3.0},
		{[]float64{70.0, 50.0, 30.0, 20.0, 10.0, 5.0, 2.0, 1.0}, 1.0, 70.0},
	}
	for _, table := range tables {
		min, max := shapes.MinMax(table.floats)
		if (min != table.min) {
			t.Errorf("Expected Minimum %v, got %v", table.min, min)
		}
		if (max != table.max) {
			t.Errorf("Expected Maximum %v, got %v", table.max, max)
		}
	}
}

func TestCombineBounds(t *testing.T) {
	tables := []struct {
		bounds []*shapes.Bounds
		min, max primitives.PV
	}{
		{[]*shapes.Bounds{{Min:primitives.MakePoint(-7, -1, -1), Max:primitives.MakePoint(1, 1, 1)},
						  {Min:primitives.MakePoint(-1, -1, -1), Max:primitives.MakePoint(8, 1, 1)},
						  {Min:primitives.MakePoint(-1, -16, -1), Max:primitives.MakePoint(1, 1, 1)},
						  {Min:primitives.MakePoint(-1, -1, -1), Max:primitives.MakePoint(1, 12, 1)},
						  {Min:primitives.MakePoint(-1, -1, -5), Max:primitives.MakePoint(1, 1, 1)},
						  {Min:primitives.MakePoint(-1, -1, -1), Max:primitives.MakePoint(1, 1, 9)}},
		 primitives.MakePoint(-7, -16, -5), primitives.MakePoint(8, 12, 9)},
	}
	for _, table := range tables {
		bounds := shapes.CombineBounds(table.bounds)
		if !bounds.Min.Equals(table.min) {
			t.Errorf("Expected Minimum %v, got %v", table.min, bounds.Min)
		}
		if !bounds.Max.Equals(table.max) {
			t.Errorf("Expected Maximum %v, got %v", table.max, bounds.Max)
		}
	}
}

func TestAddBounds(t *testing.T) {
	tables := []struct{
		bounds, other *shapes.Bounds
		min, max primitives.PV
	}{
		{&shapes.Bounds{Min:primitives.MakePoint(-3, -3, -3), Max:primitives.MakePoint(1, 1, 1)},
		 &shapes.Bounds{Min:primitives.MakePoint(-1, -1, -1), Max:primitives.MakePoint(5, 5, 5)},
		 primitives.MakePoint(-3, -3, -3), primitives.MakePoint(5, 5, 5)},
	}
	for _, table := range tables {
		table.bounds.AddBounds(table.other)
		if !table.bounds.Min.Equals(table.min) {
			t.Errorf("Expected Minimum %v, got %v", table.min, table.bounds.Min)
		}
		if !table.bounds.Max.Equals(table.max) {
			t.Errorf("Expected Maximum %v, got %v", table.max, table.bounds.Max)
		}
	}
}

func TestBoundsTransform(t *testing.T) {
	tables := []struct{
		bounds shapes.Bounds
		transform primitives.Matrix
		min, max primitives.PV
	}{
		{shapes.Bounds{Min:primitives.MakePoint(-1, -1, -1), Max:primitives.MakePoint(1, 1, 1)},
		 primitives.Translation(-1, 0, 0).Multiply(primitives.Scaling(2, 2, 2)),
		 primitives.MakePoint(-3, -2, -2), primitives.MakePoint(1, 2, 2)},
	}
	for _, table := range tables {
		bounds := table.bounds.Transform(table.transform)
		if !bounds.Min.Equals(table.min) {
			t.Errorf("Expected Minimum %v, got %v", table.min, bounds.Min)
		}
		if !bounds.Max.Equals(table.max) {
			t.Errorf("Expected Maximum %v, got %v", table.max, bounds.Max)
		}
	}
}

func TestBoundsIntersect(t *testing.T) {
	tables := []struct{
		bounds shapes.Bounds
		transform primitives.Matrix
		ray primitives.Ray
		hit bool
	}{
		{shapes.Bounds{Min:primitives.MakePoint(-1, -1, -1), Max:primitives.MakePoint(1, 1, 1)},
		 primitives.Translation(-1, 0, 0).Multiply(primitives.Scaling(2, 2, 2)),
		 primitives.Ray{Origin:primitives.MakePoint(1, 0, -3), Direction:primitives.MakeVector(-1, 0, 1)},
		 true},

		{shapes.Bounds{Min:primitives.MakePoint(-1, -1, -1), Max:primitives.MakePoint(1, 1, 1)},
		 primitives.Translation(-1, 0, 0).Multiply(primitives.Scaling(2, 2, 2)),
		 primitives.Ray{Origin:primitives.MakePoint(1, 0, -3), Direction:primitives.MakeVector(1, 0, 1)},
		 false},

		{shapes.Bounds{Min:primitives.MakePoint(8, -2, -2), Max:primitives.MakePoint(12, 2, 2)},
		 primitives.MakeIdentityMatrix(4),
		 primitives.Ray{Origin:primitives.MakePoint(10, 0, -10), Direction:primitives.MakeVector(0, 0, 1)},
		 true},
		
		{shapes.Bounds{Min:primitives.MakePoint(-1, -1, -1), Max:primitives.MakePoint(1, 1, 1)},
		 primitives.Scaling(2, 2, 2).Multiply(primitives.Translation(5, 0, 0)),
		 primitives.Ray{Origin:primitives.MakePoint(10, 0, -10), Direction:primitives.MakeVector(0, 0, 1)},
		 true},
	}
	for _, table := range tables {
		bounds := table.bounds.Transform(table.transform)
		hit := bounds.Intersect(table.ray)
		if (hit != table.hit) {
			t.Errorf("Expected %v, got %v", table.hit, hit)
		}
	}
}
