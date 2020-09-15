package shapes

import (
	"testing"
	"github.com/factorion/graytracer/pkg/primitives"
)

func TestPlaneIntersection(t *testing.T) {
	tables := []struct {
		p *Plane
		r primitives.Ray
		transform primitives.Matrix
		hits []float64
	}{
		{MakePlane(),
		 primitives.Ray{Origin:primitives.MakePoint(0, 10, 0), Direction:primitives.MakeVector(0, 0, 1)},
		 primitives.MakeIdentityMatrix(4), []float64{}},

		{MakePlane(),
		 primitives.Ray{Origin:primitives.MakePoint(0, 0, 0), Direction:primitives.MakeVector(0, 0, 1)},
		 primitives.MakeIdentityMatrix(4), []float64{}},

		{MakePlane(),
		 primitives.Ray{Origin:primitives.MakePoint(0, 1, 0), Direction:primitives.MakeVector(0, -1, 0)},
		 primitives.MakeIdentityMatrix(4), []float64{1}},

		{MakePlane(),
		 primitives.Ray{Origin:primitives.MakePoint(0, -1, 0), Direction:primitives.MakeVector(0, 1, 0)},
		 primitives.MakeIdentityMatrix(4), []float64{1}},
	}
	for _, table := range tables {
		table.p.SetTransform(table.transform)
		hits := table.p.Intersect(table.r)
		if !SliceEquals(hits, table.hits) {
			t.Errorf("Expected hit %v, got %v", table.hits, hits)
		}
	}
}

func TestPlaneNormal(t *testing.T) {
	tables := []struct {
		p *Plane
		point, normal primitives.PV
	}{
		{MakePlane(),
		 primitives.MakePoint(0, 0, 0),
		 primitives.MakeVector(0, 1, 0)},

		{MakePlane(),
		 primitives.MakePoint(10, 0, -10),
		 primitives.MakeVector(0, 1, 0)},

		{MakePlane(),
		 primitives.MakePoint(-5, 0, 150),
		 primitives.MakeVector(0, 1, 0)},
	}
	for _, table := range tables {
		normal := table.p.Normal(table.point)
		if !normal.Equals(table.normal) {
			t.Errorf("Expected %v, got %v", table.normal, normal)
		}
	}
}
