package primitives

import (
	"math"
	"testing"
)

func TestLinearMovement(t *testing.T) {
	tables := []struct {
		origin, result PV
		x, y, z float64
		transformation func(x, y, z float64) Matrix
	}{
		{MakePoint(-3, 4, 5), MakePoint(2, 1, 7), 5, -3, 2, Translation},
		{MakeVector(-3, 4, 5), MakeVector(-3, 4, 5), 5, -3, 2, Translation},
		{MakePoint(-4, 6, 8), MakePoint(-8, 18, 32), 2, 3, 4, Scaling},
		{MakeVector(-4, 6, 8), MakeVector(-8, 18, 32), 2, 3, 4, Scaling},
	}
	for _, table := range tables {
		transform := table.transformation(table.x, table.y, table.z)
		result := transform.MultiplyPV(table.origin)
		if result != table.result {
			t.Errorf("Expected %v, got %v", table.result, result)
		}
	}
}

func TestRotation(t *testing.T) {
	tables := []struct {
		origin, result PV
		rad float64
		rotation func(rad float64) Matrix
		axis string
	}{
		{MakePoint(0, 1, 0), MakePoint(0, math.Sqrt(2) / 2, math.Sqrt(2) / 2), math.Pi / 4, RotationX, "X"},
		{MakePoint(0, 1, 0), MakePoint(0, 0, 1), math.Pi / 2, RotationX, "X"},
		{MakePoint(0, 0, 1), MakePoint(math.Sqrt(2) / 2, 0, math.Sqrt(2) / 2), math.Pi / 4, RotationY, "Y"},
		{MakePoint(0, 0, 1), MakePoint(1, 0, 0), math.Pi / 2, RotationY, "Y"},
		{MakePoint(0, 1, 0), MakePoint(-math.Sqrt(2) / 2, math.Sqrt(2) / 2, 0), math.Pi / 4, RotationZ, "Z"},
		{MakePoint(0, 1, 0), MakePoint(-1, 0, 0), math.Pi / 2, RotationZ, "Z"},
	}
	for _, table := range tables {
		rotation := table.rotation(table.rad)
		result := rotation.MultiplyPV(table.origin)
		if !result.Equals(table.result) {
			t.Errorf("Rotation%s: Expected %v, got %v", table.axis, table.result, result)
		}
	}
}

func TestShear(t *testing.T) {
	tables := []struct {
		origin, result PV
		xy, xz, yx, yz, zx, zy float64
	}{
		{MakePoint(2, 3, 4), MakePoint(5, 3, 4), 1, 0, 0, 0, 0, 0},
		{MakePoint(2, 3, 4), MakePoint(6, 3, 4), 0, 1, 0, 0, 0, 0},
		{MakePoint(2, 3, 4), MakePoint(2, 5, 4), 0, 0, 1, 0, 0, 0},
		{MakePoint(2, 3, 4), MakePoint(2, 7, 4), 0, 0, 0, 1, 0, 0},
		{MakePoint(2, 3, 4), MakePoint(2, 3, 6), 0, 0, 0, 0, 1, 0},
		{MakePoint(2, 3, 4), MakePoint(2, 3, 7), 0, 0, 0, 0, 0, 1},
	}
	for _, table := range tables {
		shear := Shearing(table.xy, table.xz, table.yx, table.yz, table.zx, table.zy)
		result := shear.MultiplyPV(table.origin)
		if !result.Equals(table.result) {
			t.Errorf("Expected %v, got %v", table.result, result)
		}
	}
}

func TestSequenceAndChain(t *testing.T) {
	tables := []struct {
		origin, result PV
		transforms []Matrix
	}{
		{MakePoint(1, 0, 1), MakePoint(15, 0, 7), []Matrix{RotationX(math.Pi / 2), Scaling(5, 5, 5), Translation(10, 5, 7)}},
	}
	for _, table := range tables {
		result := table.transforms[0].MultiplyPV(table.origin)
		for _, transform := range table.transforms[1:] {
			result = transform.MultiplyPV(result)
		}
		if !result.Equals(table.result) {
			t.Errorf("Sequence: Expected %v, got %v", table.result, result)
		}
	}
	for _, table := range tables {
		sequence := MakeIdentityMatrix(4)
		for i := len(table.transforms) - 1; i >= 0; i-- {
			sequence = sequence.Multiply(table.transforms[i])
		}
		result := sequence.MultiplyPV(table.origin)
		if !result.Equals(table.result) {
			t.Errorf("Chain: Expected %v, got %v", table.result, result)
		}
	}
}
