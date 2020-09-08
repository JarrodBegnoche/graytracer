package primitives

import (
	"testing"
	"math"
)

func TestMakePV(t * testing.T) {
	tables := []struct {
		pv1, pv2 PV
	}{
		{MakePoint(1, 2, 3), PV{1, 2, 3, 1}},
		{MakeVector(3, 2, 1), PV{3, 2, 1, 0}},
	}
	for _, table := range tables {
		if table.pv1 != table.pv2 {
			t.Errorf("Expected %v, got %v", table.pv1, table.pv2)
		}
	}
}

func TestPVEquals(t *testing.T) {
	tables := []struct {
		pv1, pv2 PV
		equals bool
	}{
		{MakePoint(1, 2, 3), MakePoint(1, 2, 3), true},
		{MakePoint(1, 2, 3), MakeVector(1, 2, 3), false},
		{MakePoint(1, 1, 1), MakePoint(0, 1, 1), false},
		{MakePoint(1, 1, 1), MakePoint(1, 0, 1), false},
		{MakePoint(1, 1, 1), MakePoint(1, 1, 0), false},
		{MakePoint(1, 2, 3.123456789), MakePoint(1, 2, 3.123456788), true},
	}
	for _, table := range tables {
		equals := table.pv1.Equals(table.pv2)
		if equals != table.equals {
			t.Errorf("PV %v and %v returned %v for Equals", table.pv1, table.pv2, equals)
		}
	}
}

func TestPVMath(t *testing.T) {
	tables := []struct {
		pv1, pv2, result PV
		math func(PV, PV) PV
	}{
		{MakePoint(3, -2, 5), MakeVector(-2, 3, 1), MakePoint(1, 1, 6), PV.Add},
		{MakePoint(-4, 7, 2), MakeVector(3, 1, 1), MakePoint(-1, 8, 3), PV.Add},
		{MakePoint(3, 2, 1), MakeVector(5, 6, 7), MakePoint(-2, -4, -6), PV.Subtract},
		{MakeVector(3, -2, 5), MakeVector(-2, 3, 1), MakeVector(5, -5, 4), PV.Subtract},
	}
	for _, table := range tables {
		result := table.math(table.pv1, table.pv2)
		if !result.Equals(table.result) {
			t.Errorf("Expected %v, got %v", table.result, result)
		}
	}
}

func TestPVTransform(t *testing.T) {
	tables := []struct {
		matrix1 Matrix
		p, product PV
	}{
		{Matrix{{1, 2, 3, 4}, {2, 4, 4, 2}, {8, 6, 4, 1}, {0, 0, 0, 1}},
		 MakePoint(1, 2, 3), MakePoint(18, 24, 33)},

		{MakeMatrix(3), MakeVector(4, 3, 2), MakeVector(0, 0, 0)},
	}
	for _, table := range tables {
		product := table.p.Transform(table.matrix1)
		if product != table.product {
			t.Errorf("Expect %v, got %v", table.product, product)
		}
	}
}

func TestPVNegate(t *testing.T) {
	tables := []struct {
		v, n PV
	}{
		{MakeVector(1, -2, 3), MakeVector(-1, 2, -3)},
	}
	for _, table := range tables {
		negative := table.v.Negate()
		if !negative.Equals(table.n) {
			t.Errorf("Expected %v, got %v", table.n, negative)
		}
	}
}

func TestPVScalar(t *testing.T) {
	tables := []struct {
		v, r PV
		s float64
	}{
		{MakePoint(1, -2, 3), MakePoint(3.5, -7, 10.5), 3.5},
		{MakeVector(1, -2, 3), MakeVector(0.5, -1, 1.5), 0.5},
	}
	for _, table := range tables {
		scalar := table.v.Scalar(table.s)
		if !scalar.Equals(table.r) {
			t.Errorf("Expected %v, got %v", table.r, scalar)
		}
	}
}

func TestPVMagnitude(t *testing.T) {
	tables := []struct {
		v PV
		m float64
	}{
		{MakeVector(1, 0, 0), 1},
		{MakeVector(0, 1, 0), 1},
		{MakeVector(0, 0, 1), 1},
		{MakeVector(1, 2, 3), math.Sqrt(14)},
		{MakeVector(-1, -2, -3), math.Sqrt(14)},
	}
	for _, table := range tables {
		magnitude := table.v.Magnitude()
		if magnitude != table.m {
			t.Errorf("Expected %v, got %v", table.m, magnitude)
		}
	}
}

func TestPVNormalize(t *testing.T) {
	tables := []struct {
		v, n PV
	}{
		{MakeVector(4, 0, 0), MakeVector(1, 0, 0)},
		{MakeVector(1, 2, 3), MakeVector(0.2672612419124244, 0.5345224838248488, 0.8017837257372732)},
	}
	for _, table := range tables {
		normal := table.v.Normalize()
		if !normal.Equals(table.n) {
			t.Errorf("Expected %v, got %v", table.n, normal)
		}
	}
}

func TestPVDotProduct(t *testing.T) {
	tables := []struct {
		v, u PV
		d float64
	}{
		{MakeVector(1, 2, 3), MakeVector(2, 3, 4), 20},
	}
	for _, table := range tables {
		dot := table.v.DotProduct(table.u)
		if dot != table.d {
			t.Errorf("Expected %v, got %v", table.d, dot)
		}
	}
}

func TestCrossProduct(t *testing.T) {
	tables := []struct {
		v, u, c PV
	}{
		{MakeVector(1, 0, 0), MakeVector(0, 1, 0), MakeVector(0, 0, 1)},
		{MakeVector(0, 1, 0), MakeVector(0, 0, 1), MakeVector(1, 0, 0)},
		{MakeVector(0, 0, 1), MakeVector(1, 0, 0), MakeVector(0, 1, 0)},
	}
	for _, table := range tables {
		cross := table.v.CrossProduct(table.u)
		if !cross.Equals(table.c) {
			t.Errorf("Expected %v, got %v", table.c, cross)
		}
	}
}

func TestReflect(t *testing.T) {
	tables := []struct {
		start, normal, reflect PV
	}{
		{MakeVector(1, -1, 0), MakeVector(0, 1, 0), MakeVector(1, 1, 0)},
		{MakeVector(0, -1, 0), MakeVector(0.7071067811865476, 0.7071067811865476, 0), MakeVector(1, 0, 0)},
	}
	for _, table := range tables {
		reflect := table.start.Reflect(table.normal)
		if !reflect.Equals(table.reflect) {
			t.Errorf("Expect %v, got %v", table.reflect, reflect)
		}
	}
}
