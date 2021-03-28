package shapes_test

import (
	"testing"
	"github.com/factorion/graytracer/pkg/shapes"
)

func TestIntersectEquals(t *testing.T) {
	tables := []struct {
		xs shapes.Intersections
		dists []float64
		equals bool
	}{
		{shapes.Intersections{shapes.Intersection{Distance:0.0, Obj:nil},
							  shapes.Intersection{Distance:1.0, Obj:nil}}, []float64{0.0, 1.0}, true},
		
    	{shapes.Intersections{shapes.Intersection{Distance:0.0, Obj:nil},
							  shapes.Intersection{Distance:1.0, Obj:nil}}, []float64{1.0, 2.0}, false},
		
		{shapes.Intersections{shapes.Intersection{Distance:0.0, Obj:nil},
							  shapes.Intersection{Distance:1.0, Obj:nil},
							  shapes.Intersection{Distance:2.0, Obj:nil}}, []float64{0.0, 1.0}, false},
		
		{shapes.Intersections{shapes.Intersection{Distance:0.0, Obj:nil},
							  shapes.Intersection{Distance:1.0, Obj:nil}}, []float64{0.0, 1.0000000001}, true},
	}
	for _, table := range tables {
		equals := shapes.IntersectEquals(table.xs, table.dists)
		if equals != table.equals {
			t.Errorf("Intersections %v and %v returned %v as equals", table.xs, table.dists, equals)
		}
	}
}
