package components

import (
	"sort"
	"testing"
	"github.com/factorion/graytracer/pkg/primitives"
	"github.com/factorion/graytracer/pkg/shapes"
)

func TestIntersection(t *testing.T) {
	tables := []struct {
		inters Intersections
		distance float64
		hit bool
	}{
		{[]Intersection{Intersection{1, shapes.MakeSphere()}, Intersection{2, shapes.MakeSphere()}}, 1, true},

		{[]Intersection{Intersection{-1, shapes.MakeSphere()}, Intersection{1, shapes.MakeSphere()}}, 1, true},

		{[]Intersection{Intersection{-12, shapes.MakeSphere()}, Intersection{-11, shapes.MakeSphere()}}, -1, false},

		{[]Intersection{Intersection{5, shapes.MakeSphere()}, Intersection{7, shapes.MakeSphere()},
		                Intersection{-3, shapes.MakeSphere()}, Intersection{2, shapes.MakeSphere()}}, 2, true},
	}
	for _, table := range tables {
		sort.Sort(table.inters)
		intersection, hit := table.inters.Hit()
		if hit != table.hit {
			t.Errorf("Expected hit %v, got %v", table.hit, hit)
		}
		if hit && (intersection.Distance != table.distance) {
			t.Errorf("Expected distance %v, got %v", table.distance, intersection.Distance)
		}
	}
}

func TestPrepareComputations(t *testing.T) {
	tables := []struct {
		i Intersection
		ray primitives.Ray
		point, eyev, normalv primitives.PV
		inside bool
	}{
		{Intersection{Distance:4, Obj:shapes.MakeSphere()},
		 primitives.Ray{Origin:primitives.MakePoint(0, 0, -5), Direction:primitives.MakeVector(0, 0, 1)},
		 primitives.MakePoint(0, 0, -1),
		 primitives.MakeVector(0, 0, -1),
		 primitives.MakeVector(0, 0, -1),
		 false},

		{Intersection{Distance:1, Obj:shapes.MakeSphere()},
		 primitives.Ray{Origin:primitives.MakePoint(0, 0, 0), Direction:primitives.MakeVector(0, 0, 1)},
		 primitives.MakePoint(0, 0, 1),
		 primitives.MakeVector(0, 0, -1),
		 primitives.MakeVector(0, 0, -1),
		 true},
	}
	for _, table := range tables {
		comp := table.i.PrepareComputations(table.ray)
		if !comp.Point.Equals(table.point) {
			t.Errorf("Wrong intersection point, expect %v, got %v", table.point, comp.Point)
		}
		if !comp.EyeVector.Equals(table.eyev) {
			t.Errorf("Wrong eye vector, expected %v, got %v", table.eyev, comp.EyeVector)
		}
		if !comp.NormalVector.Equals(table.normalv) {
			t.Errorf("Wrong normal vector, expected %v, got %v", table.normalv, comp.NormalVector)
		}
		if comp.Inside != table.inside {
			t.Errorf("Wrong inside value: %v", comp.Inside)
		}
	}
}