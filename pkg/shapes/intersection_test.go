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

func TestIntersection(t *testing.T) {
	tables := []struct {
		inters []Intersection
		distance float64
		hit bool
	}{
		{[]Intersection{Intersection{1, MakeSphere()}, Intersection{2, MakeSphere()}}, 1, true},

		{[]Intersection{Intersection{-1, MakeSphere()}, Intersection{1, MakeSphere()}}, 1, true},

		{[]Intersection{Intersection{-12, MakeSphere()}, Intersection{-11, MakeSphere()}}, -1, false},

		{[]Intersection{Intersection{5, MakeSphere()}, Intersection{7, MakeSphere()},
		                Intersection{-3, MakeSphere()}, Intersection{2, MakeSphere()}}, 2, true},
	}
	for _, table := range tables {
		intersection, hit := Hit(table.inters)
		if hit != table.hit {
			t.Errorf("Expected hit %v, got %v", table.hit, hit)
		}
		if hit && (intersection.Distance != table.distance) {
			t.Errorf("Expected distance%v, got %v", table.distance, intersection.Distance)
		}
	}
}