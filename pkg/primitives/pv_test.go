package primitives

import (
	"testing"
	"math"
)

func TestPVMake(t * testing.T) {
	tables := []struct {
		pv1, pv2 pv
	}{
		{MakePoint(1, 2, 3), pv{1, 2, 3, 1}},
		{MakeVector(3, 2, 1), pv{3, 2, 1, 0}},
	}
	for _, table := range tables {
		if table.pv1 != table.pv2 {
			t.Errorf("Expected %v, got %v", table.pv1, table.pv2)
		}
	}
}

func TestPVEquals(t *testing.T) {
	tables := []struct {
		pv1, pv2 pv
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
		pv1, pv2, result pv
		math func(pv, pv) pv
	}{
		{MakePoint(3, -2, 5), MakeVector(-2, 3, 1), MakePoint(1, 1, 6), pv.Add},
		{MakePoint(-4, 7, 2), MakeVector(3, 1, 1), MakePoint(-1, 8, 3), pv.Add},
		{MakePoint(3, 2, 1), MakeVector(5, 6, 7), MakePoint(-2, -4, -6), pv.Subtract},
		{MakeVector(3, -2, 5), MakeVector(-2, 3, 1), MakeVector(5, -5, 4), pv.Subtract},
	}
	for _, table := range tables {
		result := table.math(table.pv1, table.pv2)
		if !result.Equals(table.result) {
			t.Errorf("Expected %v, got %v", table.result, result)
		}
	}
}

func TestPVNegate(t *testing.T) {
	tables := []struct {
		v, n pv
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
		v, r pv
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
		v pv
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
		v, n pv
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
		v, u pv
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
		v, u, c pv
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
