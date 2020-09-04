package shapes

import (
	"testing"
)

func TestSliceEquals(t *testing.T) {
	tables := []struct {
		a, b []float64
		equals bool
	}{
		{[]float64{0.0, 1.0}, []float64{0.0, 1.0}, true},
		{[]float64{0.0, 1.0}, []float64{1.0, 2.0}, false},
		{[]float64{0.0, 1.0, 2.0}, []float64{0.0, 1.0}, false},
		{[]float64{0.0, 1.0}, []float64{0.0, 1.0000000001}, true},
	}
	for _, table := range tables {
		equals := SliceEquals(table.a, table.b)
		if equals != table.equals {
			t.Errorf("Slice %v and %v returned %v as equals", table.a, table.b, equals)
		}
	}
}

func TestIntersections(t *testing.T) {
	tables := []struct {
		inters Intersections
		hit float64
	}{
		{Intersections{1:Intersection{1, MakeSphere()}, 2:Intersection{2, MakeSphere()}}, 1},

		{Intersections{-1:Intersection{-1, MakeSphere()}, 1:Intersection{1, MakeSphere()}}, 1},

		{Intersections{-12:Intersection{-12, MakeSphere()}, -11:Intersection{-11, MakeSphere()}}, -1},

		{Intersections{5:Intersection{5, MakeSphere()}, 7:Intersection{7, MakeSphere()},
		               -3:Intersection{-3, MakeSphere()}, 2:Intersection{2, MakeSphere()}}, 2},
	}
	for _, table := range tables {
		hit := Hit(table.inters)
		if hit != table.hit {
			t.Errorf("Expected %v, got %v", table.hit, hit)
		}
	}
}