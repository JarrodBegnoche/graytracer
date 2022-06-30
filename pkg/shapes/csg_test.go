package shapes_test

import (
	"testing"

	"github.com/factorion/graytracer/pkg/primitives"
	"github.com/factorion/graytracer/pkg/shapes"
)

func TestCSGIntersectionAllowed(t *testing.T) {
	tables := []struct {
		csg                     *shapes.CSG
		op                      shapes.Operation
		lhit, inl, inr, allowed bool
	}{
		{shapes.MakeCSG(shapes.UNION, shapes.MakeSphere(), shapes.MakeSphere()), shapes.UNION,
			true, true, true, false},
		{shapes.MakeCSG(shapes.UNION, shapes.MakeSphere(), shapes.MakeSphere()), shapes.UNION,
			true, true, false, true},
		{shapes.MakeCSG(shapes.UNION, shapes.MakeSphere(), shapes.MakeSphere()), shapes.UNION,
			true, false, true, false},
		{shapes.MakeCSG(shapes.UNION, shapes.MakeSphere(), shapes.MakeSphere()), shapes.UNION,
			true, false, false, true},
		{shapes.MakeCSG(shapes.UNION, shapes.MakeSphere(), shapes.MakeSphere()), shapes.UNION,
			false, true, true, false},
		{shapes.MakeCSG(shapes.UNION, shapes.MakeSphere(), shapes.MakeSphere()), shapes.UNION,
			false, true, false, false},
		{shapes.MakeCSG(shapes.UNION, shapes.MakeSphere(), shapes.MakeSphere()), shapes.UNION,
			false, false, true, true},
		{shapes.MakeCSG(shapes.UNION, shapes.MakeSphere(), shapes.MakeSphere()), shapes.UNION,
			false, false, false, true},

		{shapes.MakeCSG(shapes.INTERSECT, shapes.MakeSphere(), shapes.MakeSphere()), shapes.INTERSECT,
			true, true, true, true},
		{shapes.MakeCSG(shapes.INTERSECT, shapes.MakeSphere(), shapes.MakeSphere()), shapes.INTERSECT,
			true, true, false, false},
		{shapes.MakeCSG(shapes.INTERSECT, shapes.MakeSphere(), shapes.MakeSphere()), shapes.INTERSECT,
			true, false, true, true},
		{shapes.MakeCSG(shapes.INTERSECT, shapes.MakeSphere(), shapes.MakeSphere()), shapes.INTERSECT,
			true, false, false, false},
		{shapes.MakeCSG(shapes.INTERSECT, shapes.MakeSphere(), shapes.MakeSphere()), shapes.INTERSECT,
			false, true, true, true},
		{shapes.MakeCSG(shapes.INTERSECT, shapes.MakeSphere(), shapes.MakeSphere()), shapes.INTERSECT,
			false, true, false, true},
		{shapes.MakeCSG(shapes.INTERSECT, shapes.MakeSphere(), shapes.MakeSphere()), shapes.INTERSECT,
			false, false, true, false},
		{shapes.MakeCSG(shapes.INTERSECT, shapes.MakeSphere(), shapes.MakeSphere()), shapes.INTERSECT,
			false, false, false, false},

		{shapes.MakeCSG(shapes.DIFFERENCE, shapes.MakeSphere(), shapes.MakeSphere()), shapes.DIFFERENCE,
			true, true, true, false},
		{shapes.MakeCSG(shapes.DIFFERENCE, shapes.MakeSphere(), shapes.MakeSphere()), shapes.DIFFERENCE,
			true, true, false, true},
		{shapes.MakeCSG(shapes.DIFFERENCE, shapes.MakeSphere(), shapes.MakeSphere()), shapes.DIFFERENCE,
			true, false, true, false},
		{shapes.MakeCSG(shapes.DIFFERENCE, shapes.MakeSphere(), shapes.MakeSphere()), shapes.DIFFERENCE,
			true, false, false, true},
		{shapes.MakeCSG(shapes.DIFFERENCE, shapes.MakeSphere(), shapes.MakeSphere()), shapes.DIFFERENCE,
			false, true, true, true},
		{shapes.MakeCSG(shapes.DIFFERENCE, shapes.MakeSphere(), shapes.MakeSphere()), shapes.DIFFERENCE,
			false, true, false, true},
		{shapes.MakeCSG(shapes.DIFFERENCE, shapes.MakeSphere(), shapes.MakeSphere()), shapes.DIFFERENCE,
			false, false, true, false},
		{shapes.MakeCSG(shapes.DIFFERENCE, shapes.MakeSphere(), shapes.MakeSphere()), shapes.DIFFERENCE,
			false, false, false, false},
	}
	for _, table := range tables {
		allowed := table.csg.IntersectionAllowed(table.lhit, table.inl, table.inr)
		if allowed != table.allowed {
			t.Errorf("Incorrect allowed value for operation %d with lhit=%t, inl=%t, inr=%t", table.op, table.lhit, table.inl, table.inr)
		}
	}
}

func TestIntersetction(t *testing.T) {
	tables := []struct {
		op          shapes.Operation
		csg_shapes  []shapes.Shape
		transforms  []primitives.Matrix
		ray         primitives.Ray
		xs_count    int
		distances   []float64
		shape_index []int
	}{
		/*{
			shapes.UNION,
			[]shapes.Shape{shapes.MakeSphere(), shapes.MakeCube()},
			[]primitives.Matrix{primitives.MakeIdentityMatrix(4), primitives.MakeIdentityMatrix(4)},
			primitives.Ray{Origin: primitives.MakePoint(0, 2, -5), Direction: primitives.MakeVector(0, 0, 1)},
			0,
			[]float64{},
			[]int{},
		},*/

		{
			shapes.UNION,
			[]shapes.Shape{shapes.MakeSphere(), shapes.MakeSphere()},
			[]primitives.Matrix{primitives.MakeIdentityMatrix(4), primitives.Translation(0, 0, 0.5)},
			primitives.Ray{Origin: primitives.MakePoint(0, 0, -5), Direction: primitives.MakeVector(0, 0, 1)},
			2,
			[]float64{4, 6.5},
			[]int{0, 1},
		},
	}
	for _, table := range tables {
		table.csg_shapes[0].SetTransform(table.transforms[0])
		table.csg_shapes[1].SetTransform(table.transforms[1])
		csg := shapes.MakeCSG(table.op, table.csg_shapes[0], table.csg_shapes[1])
		xs := csg.Intersect(table.ray)
		if xs.Len() != table.xs_count {
			t.Errorf("Incorrect number of intersections: %d, expected %d", xs.Len(), table.xs_count)
		}
		for index, hit := range xs {
			if hit.Distance != table.distances[index] {
				t.Errorf("Incorrect hit distance %f, expected %f", hit.Distance, table.distances[index])
			}
			if hit.Obj != table.csg_shapes[table.shape_index[index]] {
				t.Errorf("Incorrect shape index hit, expected %d", table.shape_index[index])
			}
		}
	}
}
