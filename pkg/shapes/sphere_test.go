package shapes

import (
	"testing"
	"github.com/factorion/graytracer/pkg/primitives"
)

func TestSphereIntersection(t *testing.T) {
	tables := []struct {
		s Sphere
		r primitives.Ray
		hits []float64
	}{
		{MakeSphere(0, 0, 0, 1.0),
		 primitives.Ray{Origin:primitives.MakePoint(0, 1, -5), Direction:primitives.MakeVector(0, 0, 1)},
		 []float64{5}},

		{MakeSphere(0, 0, 0, 1.0),
		 primitives.Ray{Origin:primitives.MakePoint(0, 2, -5), Direction:primitives.MakeVector(0, 0, 1)},
		 []float64{}},

		{MakeSphere(0, 0, 0, 1.0),
		 primitives.Ray{Origin:primitives.MakePoint(0, 0, 5), Direction:primitives.MakeVector(0, 0, 1)},
		 []float64{-6, -4}},

		{MakeTransformedSphere(0, 0, 0, 1.0, primitives.Scaling(2, 2, 2)),
		 primitives.Ray{Origin:primitives.MakePoint(0, 0, -5), Direction:primitives.MakeVector(0, 0, 1)},
		 []float64{3, 7}},

		{MakeTransformedSphere(0, 0, 0, 1.0, primitives.Translation(5, 0, 0)),
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
