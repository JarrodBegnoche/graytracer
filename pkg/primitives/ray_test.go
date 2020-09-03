package primitives

import (
	"testing"
)

func TestEquals(t *testing.T) {
	tables := []struct {
		ray1, ray2 Ray
		equals bool
	}{
		{Ray{MakePoint(1, 2, 3), MakeVector(0, 0, 1)}, Ray{MakePoint(1, 2, 3), MakeVector(0, 0, 1)}, true},
		{Ray{MakePoint(1, 2, 3), MakeVector(0, 0, 1)}, Ray{MakePoint(3, 2, 1), MakeVector(0, 0, 1)}, false},
		{Ray{MakePoint(1, 2, 3), MakeVector(0, 0, 1)},
		 Ray{MakePoint(1, 2, 3), MakeVector(0, 0, 1.00000000001)}, true},
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
		ray Ray
		time float64
		destination PV
	}{
		{Ray{MakePoint(2, 3, 4), MakeVector(1, 0, 0)}, 0, MakePoint(2, 3, 4)},
		{Ray{MakePoint(2, 3, 4), MakeVector(1, 0, 0)}, 1, MakePoint(3, 3, 4)},
		{Ray{MakePoint(2, 3, 4), MakeVector(1, 0, 0)}, -1, MakePoint(1, 3, 4)},
		{Ray{MakePoint(2, 3, 4), MakeVector(1, 0, 0)}, 2.5, MakePoint(4.5, 3, 4)},
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
		start, end Ray
		transform Matrix
	}{
		{Ray{MakePoint(1, 2, 3), MakeVector(0, 1, 0)}, Ray{MakePoint(4, 6, 8), MakeVector(0, 1, 0)},
		 Translation(3, 4, 5)},

		{Ray{MakePoint(1, 2, 3), MakeVector(0, 1, 0)}, Ray{MakePoint(2, 6, 12), MakeVector(0, 3, 0)},
		 Scaling(2, 3, 4)},

		{Ray{MakePoint(1, 2, 3), MakeVector(0, 1, 0)}, Ray{}, MakeMatrix(3)},
	}
	for _, table := range tables {
		result := table.start.Transform(table.transform)
		if !result.Equals(table.end) {
			t.Errorf("Expected %v, got %v", table.end, result)
		}
	}
}