package components_test

import (
	"math"
	"sort"
	"testing"
	"github.com/factorion/graytracer/pkg/components"
	"github.com/factorion/graytracer/pkg/primitives"
	"github.com/factorion/graytracer/pkg/shapes"
)

func TestIntersection(t *testing.T) {
	tables := []struct {
		inters shapes.Intersections
		distance float64
		hit bool
	}{
		{[]shapes.Intersection{{Distance:1, Obj:shapes.MakeSphere()},
							   {Distance:2, Obj:shapes.MakeSphere()}}, 1, true},

		{[]shapes.Intersection{{Distance:-1, Obj:shapes.MakeSphere()},
							   {Distance:1, Obj:shapes.MakeSphere()}}, 1, true},

		{[]shapes.Intersection{{Distance:-12, Obj:shapes.MakeSphere()},
							   {Distance:-11, Obj:shapes.MakeSphere()}}, -1, false},

		{[]shapes.Intersection{{Distance:5, Obj:shapes.MakeSphere()},
							   {Distance:7, Obj:shapes.MakeSphere()},
							   {Distance:-3, Obj:shapes.MakeSphere()},
							   {Distance:2, Obj:shapes.MakeSphere()}}, 2, true},
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
		i shapes.Intersection
		ray primitives.Ray
		point, eyev, normalv primitives.PV
		inside bool
	}{
		{shapes.Intersection{Distance:4, Obj:shapes.MakeSphere()},
		 primitives.Ray{Origin:primitives.MakePoint(0, 0, -5), Direction:primitives.MakeVector(0, 0, 1)},
		 primitives.MakePoint(0, 0, -1),
		 primitives.MakeVector(0, 0, -1),
		 primitives.MakeVector(0, 0, -1),
		 false},

		{shapes.Intersection{Distance:1, Obj:shapes.MakeSphere()},
		 primitives.Ray{Origin:primitives.MakePoint(0, 0, 0), Direction:primitives.MakeVector(0, 0, 1)},
		 primitives.MakePoint(0, 0, 1),
		 primitives.MakeVector(0, 0, -1),
		 primitives.MakeVector(0, 0, -1),
		 true},
	}
	for _, table := range tables {
		comp := components.PrepareComputations(table.i, table.ray, shapes.Intersections{})
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

func TestComputeReflection(t *testing.T) {
	tables := []struct {
		i shapes.Intersection
		ray primitives.Ray
		reflectv primitives.PV
	}{
		{shapes.Intersection{Distance:math.Sqrt(2), Obj:shapes.MakePlane()},
		 primitives.Ray{Origin:primitives.MakePoint(0, 0, -1),
						Direction:primitives.MakeVector(0, -0.7071067811865476, 0.7071067811865476)},
		 primitives.MakeVector(0, 0.7071067811865476, 0.7071067811865476)},
	}
	for _, table := range tables {
		comp := components.PrepareComputations(table.i, table.ray, shapes.Intersections{})
		if !comp.ReflectVector.Equals(table.reflectv) {
			t.Errorf("Expected %v, got %v", table.reflectv, comp.ReflectVector)
		}
	}
}

func TestComputeRefractionIndices(t *testing.T) {
	tables := []struct {
		shapes []shapes.Shape
		transforms []primitives.Matrix
		refractions []float64
		ray primitives.Ray
		distance []float64
		sIndex []int
		index1, index2 []float64
	}{
		{
			[]shapes.Shape{shapes.MakeSphere(), shapes.MakeSphere(), shapes.MakeSphere()},
			[]primitives.Matrix{primitives.Scaling(2, 2, 2), primitives.Translation(0, 0, -0.25),
								primitives.Translation(0, 0, 0.25)},
			[]float64{1.5, 2.0, 2.5},
			primitives.Ray{Origin:primitives.MakePoint(0, 0, -4), Direction:primitives.MakePoint(0, 0, 1)},
			[]float64{2, 2.75, 3.25, 4.75, 5.25, 6},
			[]int{0, 1, 2, 1, 2, 0},
			[]float64{1.0, 1.5, 2.0, 2.5, 2.5, 1.5},
			[]float64{1.5, 2.0, 2.5, 2.5, 1.5, 1.0},
		},
	}
	for _, table := range tables {
		for i, s := range table.shapes {
			s.SetTransform(table.transforms[i])
			mat := s.Material()
			mat.RefractiveIndex = table.refractions[i]
			s.SetMaterial(mat)
		}
		intersections := shapes.Intersections{}
		for i, d := range table.distance {
			intersections = append(intersections, shapes.Intersection{Distance:d,
																      Obj:table.shapes[table.sIndex[i]]})
		}
		for i, inter := range intersections {
			comp := components.PrepareComputations(inter, table.ray, intersections)
			if (comp.Index1 != table.index1[i]) || (comp.Index2 != table.index2[i]) {
				t.Errorf("Index %v: Expected indices %v, %v; got %v, %v", i, table.index1[i], table.index2[i],
						 comp.Index1, comp.Index2)
			}
		}
	}
}
