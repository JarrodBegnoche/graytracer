package primitives_test

import (
	"math"
	"testing"
	"github.com/factorion/graytracer/pkg/primitives"
)

func TestMakePV(t * testing.T) {
	tables := []struct {
		pv1, pv2 primitives.PV
	}{
		{primitives.MakePoint(1, 2, 3), primitives.PV{1, 2, 3, 1}},
		{primitives.MakeVector(3, 2, 1), primitives.PV{3, 2, 1, 0}},
	}
	for _, table := range tables {
		if table.pv1 != table.pv2 {
			t.Errorf("Expected %v, got %v", table.pv1, table.pv2)
		}
	}
}

func TestPVEquals(t *testing.T) {
	tables := []struct {
		pv1, pv2 primitives.PV
		equals bool
	}{
		{primitives.MakePoint(1, 2, 3), primitives.MakePoint(1, 2, 3), true},
		{primitives.MakePoint(1, 2, 3), primitives.MakeVector(1, 2, 3), false},
		{primitives.MakePoint(1, 1, 1), primitives.MakePoint(0, 1, 1), false},
		{primitives.MakePoint(1, 1, 1), primitives.MakePoint(1, 0, 1), false},
		{primitives.MakePoint(1, 1, 1), primitives.MakePoint(1, 1, 0), false},
		{primitives.MakePoint(1, 2, 3.123456789), primitives.MakePoint(1, 2, 3.123456788), true},
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
		pv1, pv2, result primitives.PV
		math func(primitives.PV, primitives.PV) primitives.PV
	}{
		{primitives.MakePoint(3, -2, 5), primitives.MakeVector(-2, 3, 1), primitives.MakePoint(1, 1, 6),
		 primitives.PV.Add},
		{primitives.MakePoint(-4, 7, 2), primitives.MakeVector(3, 1, 1), primitives.MakePoint(-1, 8, 3),
		 primitives.PV.Add},
		{primitives.MakePoint(3, 2, 1), primitives.MakeVector(5, 6, 7), primitives.MakePoint(-2, -4, -6),
		 primitives.PV.Subtract},
		{primitives.MakeVector(3, -2, 5), primitives.MakeVector(-2, 3, 1), primitives.MakeVector(5, -5, 4),
		 primitives.PV.Subtract},
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
		matrix1 primitives.Matrix
		p, product primitives.PV
	}{
		{primitives.Matrix{{1, 2, 3, 4}, {2, 4, 4, 2}, {8, 6, 4, 1}, {0, 0, 0, 1}},
		 primitives.MakePoint(1, 2, 3), primitives.MakePoint(18, 24, 33)},

		{primitives.MakeMatrix(3), primitives.MakeVector(4, 3, 2), primitives.MakeVector(0, 0, 0)},
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
		v, n primitives.PV
	}{
		{primitives.MakeVector(1, -2, 3), primitives.MakeVector(-1, 2, -3)},
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
		v, r primitives.PV
		s float64
	}{
		{primitives.MakePoint(1, -2, 3), primitives.MakePoint(3.5, -7, 10.5), 3.5},
		{primitives.MakeVector(1, -2, 3), primitives.MakeVector(0.5, -1, 1.5), 0.5},
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
		v primitives.PV
		m float64
	}{
		{primitives.MakeVector(1, 0, 0), 1},
		{primitives.MakeVector(0, 1, 0), 1},
		{primitives.MakeVector(0, 0, 1), 1},
		{primitives.MakeVector(1, 2, 3), math.Sqrt(14)},
		{primitives.MakeVector(-1, -2, -3), math.Sqrt(14)},
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
		v, n primitives.PV
	}{
		{primitives.MakeVector(4, 0, 0), primitives.MakeVector(1, 0, 0)},
		{primitives.MakeVector(1, 2, 3), primitives.MakeVector(0.2672612419124244, 0.5345224838248488, 0.8017837257372732)},
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
		v, u primitives.PV
		d float64
	}{
		{primitives.MakeVector(1, 2, 3), primitives.MakeVector(2, 3, 4), 20},
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
		v, u, c primitives.PV
	}{
		{primitives.MakeVector(1, 0, 0), primitives.MakeVector(0, 1, 0), primitives.MakeVector(0, 0, 1)},
		{primitives.MakeVector(0, 1, 0), primitives.MakeVector(0, 0, 1), primitives.MakeVector(1, 0, 0)},
		{primitives.MakeVector(0, 0, 1), primitives.MakeVector(1, 0, 0), primitives.MakeVector(0, 1, 0)},
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
		start, normal, reflect primitives.PV
	}{
		{primitives.MakeVector(1, -1, 0), primitives.MakeVector(0, 1, 0), primitives.MakeVector(1, 1, 0)},
		{primitives.MakeVector(0, -1, 0), primitives.MakeVector(0.7071067811865476, 0.7071067811865476, 0), primitives.MakeVector(1, 0, 0)},
	}
	for _, table := range tables {
		reflect := table.start.Reflect(table.normal)
		if !reflect.Equals(table.reflect) {
			t.Errorf("Expect %v, got %v", table.reflect, reflect)
		}
	}
}
