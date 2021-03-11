package shapes_test

import (
	"sort"
	"testing"
	"github.com/factorion/graytracer/pkg/primitives"
	"github.com/factorion/graytracer/pkg/shapes"
)


func TestGroupIntersection(t *testing.T) {
	tables := []struct {
		group *shapes.Group
		transform primitives.Matrix
		shapes []full_shape
		ray primitives.Ray
		hits []float64
	}{
		{shapes.MakeGroup(), primitives.MakeIdentityMatrix(4),
		 []full_shape{},
		 primitives.Ray{Origin:primitives.MakePoint(0, 0, 0), Direction:primitives.MakeVector(0, 0, 1)},
		 []float64{}},

		{shapes.MakeGroup(), primitives.MakeIdentityMatrix(4),
		 []full_shape{{shape:shapes.MakeSphere(), transform:primitives.MakeIdentityMatrix(4)},
				 	  {shape:shapes.MakeSphere(), transform:primitives.Translation(0, 0, -3)},
					  {shape:shapes.MakeSphere(), transform:primitives.Translation(5, 0, 0)}},
		 primitives.Ray{Origin:primitives.MakePoint(0, 0, -5), Direction:primitives.MakeVector(0, 0, 1)},
		 []float64{1, 3, 4, 6}},

		{shapes.MakeGroup(), primitives.Scaling(2, 2, 2),
		 []full_shape{{shape:shapes.MakeSphere(), transform:primitives.Translation(5, 0, 0)}},
		 primitives.Ray{Origin:primitives.MakePoint(10, 0, -10), Direction:primitives.MakeVector(0, 0, 1)},
		 []float64{8, 12}},
	}
	for _, table := range tables {
		table.group.SetTransform(table.transform)
		for _, fs := range table.shapes {
			fs.shape.SetTransform(fs.transform)
			table.group.AddShape(fs.shape)
		}
		hits := table.group.Intersect(table.ray)
		sort.Sort(hits)
		if !shapes.IntersectEquals(hits, table.hits) {
			t.Errorf("Expected hit %v, got %v", table.hits, hits)
		}
	}
}
