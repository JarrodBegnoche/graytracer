package primitives_test

import (
	"testing"
	"github.com/factorion/graytracer/pkg/primitives"
)

func TestEquals(t *testing.T) {
	tables := []struct {
		ray1, ray2 primitives.Ray
		equals bool
	}{
		{primitives.Ray{primitives.MakePoint(1, 2, 3), primitives.MakeVector(0, 0, 1)},
		 primitives.Ray{primitives.MakePoint(1, 2, 3), primitives.MakeVector(0, 0, 1)}, true},
		
		{primitives.Ray{primitives.MakePoint(1, 2, 3), primitives.MakeVector(0, 0, 1)},
		 primitives.Ray{primitives.MakePoint(3, 2, 1), primitives.MakeVector(0, 0, 1)}, false},
		
		{primitives.Ray{primitives.MakePoint(1, 2, 3), primitives.MakeVector(0, 0, 1)},
		 primitives.Ray{primitives.MakePoint(1, 2, 3), primitives.MakeVector(0, 0, 1.00000000001)}, true},
	}
	for _, table := range tables {
		equals := table.ray1.Equals(table.ray2)
		if equals != table.equals {
			t.Errorf("PV %v and %v returned %v for Equals", table.ray1, table.ray2, equals)
		}
	}
}
func TestPosition(t *testing.T) {
	tables := []struct {
		ray primitives.Ray
		time float64
		destination primitives.PV
	}{
		{primitives.Ray{primitives.MakePoint(2, 3, 4), primitives.MakeVector(1, 0, 0)}, 0, primitives.MakePoint(2, 3, 4)},
		{primitives.Ray{primitives.MakePoint(2, 3, 4), primitives.MakeVector(1, 0, 0)}, 1, primitives.MakePoint(3, 3, 4)},
		{primitives.Ray{primitives.MakePoint(2, 3, 4), primitives.MakeVector(1, 0, 0)}, -1, primitives.MakePoint(1, 3, 4)},
		{primitives.Ray{primitives.MakePoint(2, 3, 4), primitives.MakeVector(1, 0, 0)}, 2.5, primitives.MakePoint(4.5, 3, 4)},
	}
	for _, table := range tables {
		destination := table.ray.Position(table.time)
		if !destination.Equals(table.destination) {
			t.Errorf("Expected %v, got %v", table.destination, destination)
		}
	}
}

func TestRayTransform(t *testing.T) {
	tables := []struct {
		start, end primitives.Ray
		transform primitives.Matrix
	}{
		{primitives.Ray{primitives.MakePoint(1, 2, 3), primitives.MakeVector(0, 1, 0)},
		 primitives.Ray{primitives.MakePoint(4, 6, 8), primitives.MakeVector(0, 1, 0)},
		 primitives.Translation(3, 4, 5)},

		{primitives.Ray{primitives.MakePoint(1, 2, 3), primitives.MakeVector(0, 1, 0)},
		 primitives.Ray{primitives.MakePoint(2, 6, 12), primitives.MakeVector(0, 3, 0)},
		 primitives.Scaling(2, 3, 4)},

		{primitives.Ray{primitives.MakePoint(1, 2, 3), primitives.MakeVector(0, 1, 0)},
		 primitives.Ray{},
		 primitives.MakeMatrix(3)},
	}
	for _, table := range tables {
		result := table.start.Transform(table.transform)
		if !result.Equals(table.end) {
			t.Errorf("Expected %v, got %v", table.end, result)
		}
	}
}

func BenchmarkRayTransform(b *testing.B) {
	ray := primitives.Ray{primitives.MakePoint(1, 2, 3), primitives.MakeVector(0, 1, 0)}
	transform := primitives.Translation(3, 4, 5)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ray.Transform(transform)
	}
}
