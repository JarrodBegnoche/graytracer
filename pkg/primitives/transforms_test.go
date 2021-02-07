package primitives_test

import (
	"math"
	"testing"
	"github.com/factorion/graytracer/pkg/primitives"
)

func TestLinearMovement(t *testing.T) {
	tables := []struct {
		origin, result primitives.PV
		x, y, z float64
		transformation func(x, y, z float64) primitives.Matrix
	}{
		{primitives.MakePoint(-3, 4, 5), primitives.MakePoint(2, 1, 7), 5, -3, 2, primitives.Translation},
		{primitives.MakeVector(-3, 4, 5), primitives.MakeVector(-3, 4, 5), 5, -3, 2, primitives.Translation},
		{primitives.MakePoint(-4, 6, 8), primitives.MakePoint(-8, 18, 32), 2, 3, 4, primitives.Scaling},
		{primitives.MakeVector(-4, 6, 8), primitives.MakeVector(-8, 18, 32), 2, 3, 4, primitives.Scaling},
	}
	for _, table := range tables {
		transform := table.transformation(table.x, table.y, table.z)
		result := table.origin.Transform(transform)
		if result != table.result {
			t.Errorf("Expected %v, got %v", table.result, result)
		}
	}
}

func TestRotation(t *testing.T) {
	tables := []struct {
		origin, result primitives.PV
		rad float64
		rotation func(rad float64) primitives.Matrix
		axis string
	}{
		{primitives.MakePoint(0, 1, 0), primitives.MakePoint(0, math.Sqrt(2) / 2, math.Sqrt(2) / 2),
		 math.Pi / 4, primitives.RotationX, "X"},
		
		{primitives.MakePoint(0, 1, 0), primitives.MakePoint(0, 0, 1),
		 math.Pi / 2, primitives.RotationX, "X"},
		
		{primitives.MakePoint(0, 0, 1), primitives.MakePoint(math.Sqrt(2) / 2, 0, math.Sqrt(2) / 2),
		 math.Pi / 4, primitives.RotationY, "Y"},
		
		{primitives.MakePoint(0, 0, 1), primitives.MakePoint(1, 0, 0),
		 math.Pi / 2, primitives.RotationY, "Y"},
		
		{primitives.MakePoint(0, 1, 0), primitives.MakePoint(-math.Sqrt(2) / 2, math.Sqrt(2) / 2, 0),
		 math.Pi / 4, primitives.RotationZ, "Z"},

		{primitives.MakePoint(0, 1, 0), primitives.MakePoint(-1, 0, 0),
		 math.Pi / 2, primitives.RotationZ, "Z"},
	}
	for _, table := range tables {
		rotation := table.rotation(table.rad)
		result := table.origin.Transform(rotation)
		if !result.Equals(table.result) {
			t.Errorf("Rotation%s: Expected %v, got %v", table.axis, table.result, result)
		}
	}
}

func TestShear(t *testing.T) {
	tables := []struct {
		origin, result primitives.PV
		xy, xz, yx, yz, zx, zy float64
	}{
		{primitives.MakePoint(2, 3, 4), primitives.MakePoint(5, 3, 4), 1, 0, 0, 0, 0, 0},
		{primitives.MakePoint(2, 3, 4), primitives.MakePoint(6, 3, 4), 0, 1, 0, 0, 0, 0},
		{primitives.MakePoint(2, 3, 4), primitives.MakePoint(2, 5, 4), 0, 0, 1, 0, 0, 0},
		{primitives.MakePoint(2, 3, 4), primitives.MakePoint(2, 7, 4), 0, 0, 0, 1, 0, 0},
		{primitives.MakePoint(2, 3, 4), primitives.MakePoint(2, 3, 6), 0, 0, 0, 0, 1, 0},
		{primitives.MakePoint(2, 3, 4), primitives.MakePoint(2, 3, 7), 0, 0, 0, 0, 0, 1},
	}
	for _, table := range tables {
		shear := primitives.Shearing(table.xy, table.xz, table.yx, table.yz, table.zx, table.zy)
		result := table.origin.Transform(shear)
		if !result.Equals(table.result) {
			t.Errorf("Expected %v, got %v", table.result, result)
		}
	}
}

func TestSequenceAndChain(t *testing.T) {
	tables := []struct {
		origin, result primitives.PV
		transforms []primitives.Matrix
	}{
		{primitives.MakePoint(1, 0, 1), primitives.MakePoint(15, 0, 7),
		 []primitives.Matrix{primitives.RotationX(math.Pi / 2), primitives.Scaling(5, 5, 5),
							 primitives.Translation(10, 5, 7)}},
	}
	for _, table := range tables {
		result := table.origin.Transform(table.transforms[0])
		for _, transform := range table.transforms[1:] {
			result = result.Transform(transform)
		}
		if !result.Equals(table.result) {
			t.Errorf("Sequence: Expected %v, got %v", table.result, result)
		}
	}
	for _, table := range tables {
		sequence := primitives.MakeIdentityMatrix(4)
		for i := len(table.transforms) - 1; i >= 0; i-- {
			sequence = sequence.Multiply(table.transforms[i])
		}
		result := table.origin.Transform(sequence)
		if !result.Equals(table.result) {
			t.Errorf("Chain: Expected %v, got %v", table.result, result)
		}
	}
}
