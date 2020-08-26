package primitives

import (
	"testing"
)

func TestPosition(t *testing.T) {
	tables := []struct {
		ray Ray
		time float64
		destination pv
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