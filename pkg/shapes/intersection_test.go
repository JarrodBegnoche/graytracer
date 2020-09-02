package shapes

import (
	"testing"
)

func TestIntersections(t *testing.T) {
	tables := []struct {
		inters Intersections
		hit float64
	}{
		{Intersections{1:Intersection{1, Sphere{}}, 2:Intersection{2, Sphere{}}}, 1},

		{Intersections{-1:Intersection{-1, Sphere{}}, 1:Intersection{1, Sphere{}}}, 1},

		{Intersections{-12:Intersection{-12, Sphere{}}, -11:Intersection{-11, Sphere{}}}, -1},

		{Intersections{5:Intersection{5, Sphere{}}, 7:Intersection{7, Sphere{}},
		               -3:Intersection{-3, Sphere{}}, 2:Intersection{2, Sphere{}}}, 2},
	}
	for _, table := range tables {
		hit := Hit(table.inters)
		if hit != table.hit {
			t.Errorf("Expected %v, got %v", table.hit, hit)
		}
	}
}