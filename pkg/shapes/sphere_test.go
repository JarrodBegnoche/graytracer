package shapes

import (
	"testing"
	"github.com/factorion/graytracer/pkg/primitives"
)

func TestSphereIntersection(t *testing.T) {
	tables := []struct {
		s Sphere
		r primitives.Ray
		hit bool
		t1, t2 float64
	}{
		{Sphere{primitives.MakePoint(0, 0, 0), 1.0},
		 primitives.Ray{Origin:primitives.MakePoint(0, 1, -5), Direction:primitives.MakeVector(0, 0, 1)},
		 true, 5, 5},

		{Sphere{primitives.MakePoint(0, 0, 0), 1.0},
		 primitives.Ray{Origin:primitives.MakePoint(0, 2, -5), Direction:primitives.MakeVector(0, 0, 1)},
		 false, 0, 0},

		{Sphere{primitives.MakePoint(0, 0, 0), 1.0},
		 primitives.Ray{Origin:primitives.MakePoint(0, 0, 5), Direction:primitives.MakeVector(0, 0, 1)},
		 true, -6, -4},
	}
	for _, table := range tables {
		hit, t1, t2 := table.s.Intersect(table.r)
		if hit != table.hit {
			t.Errorf("Expected hit %v, got %v", table.hit, hit)
		}
		if t1 != table.t1 {
			t.Errorf("Expected first hit %v, got %v", table.t1, t1)
		}
		if t2 != table.t2 {
			t.Errorf("Expected first hit %v, got %v", table.t2, t2)
		}
	}
}
