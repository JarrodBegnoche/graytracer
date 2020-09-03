package shapes

import (
	"testing"
	"github.com/factorion/graytracer/pkg/primitives"
)

func SliceEquals(a, b []float64) bool {
	if len(a) != len(b) {
        return false
    }
    for i, v := range a {
        if v != b[i] {
            return false
        }
    }
    return true
}

func TestSphereIntersection(t *testing.T) {
	tables := []struct {
		s Sphere
		r primitives.Ray
		hits []float64
	}{
		{Sphere{center:primitives.MakePoint(0, 0, 0), radius:1.0, transform:primitives.MakeIdentityMatrix(4)},
		 primitives.Ray{Origin:primitives.MakePoint(0, 1, -5), Direction:primitives.MakeVector(0, 0, 1)},
		 []float64{5}},

		{Sphere{center:primitives.MakePoint(0, 0, 0), radius:1.0, transform:primitives.MakeIdentityMatrix(4)},
		 primitives.Ray{Origin:primitives.MakePoint(0, 2, -5), Direction:primitives.MakeVector(0, 0, 1)},
		 []float64{}},

		{Sphere{center:primitives.MakePoint(0, 0, 0), radius:1.0, transform:primitives.MakeIdentityMatrix(4)},
		 primitives.Ray{Origin:primitives.MakePoint(0, 0, 5), Direction:primitives.MakeVector(0, 0, 1)},
		 []float64{-6, -4}},

		{Sphere{center:primitives.MakePoint(0, 0, 0), radius:1.0, transform:primitives.Scaling(2, 2, 2)},
		 primitives.Ray{Origin:primitives.MakePoint(0, 0, -5), Direction:primitives.MakeVector(0, 0, 1)},
		 []float64{3, 7}},

		 {Sphere{center:primitives.MakePoint(0, 0, 0), radius:1.0, transform:primitives.Translation(5, 0, 0)},
		  primitives.Ray{Origin:primitives.MakePoint(0, 0, -5), Direction:primitives.MakeVector(0, 0, 1)},
		  []float64{}},
	}
	for _, table := range tables {
		hits := table.s.Intersect(table.r)
		if !SliceEquals(hits, table.hits) {
			t.Errorf("Expected hit %v, got %v", table.hits, hits)
		}
	}
}
