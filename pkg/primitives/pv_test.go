package primitives

import (
	"testing"
	"math"
)

func TestMakePoint(t *testing.T) {
	tables := []struct {
		x float64
		y float64
		z float64
		point PV
	}{
		{4.0, -4.0, 3.0, PV{4.0, -4.0, 3.0, 1.0}},
	}
	for _, table := range tables {
		point := MakePoint(table.x, table.y, table.z)
		if point != table.point {
			t.Errorf("Expected %v, got %v", table.point, point)
		}
	}
}

func TestMakeVector(t *testing.T) {
	tables := []struct {
		x float64
		y float64
		z float64
		vector PV
	}{
		{4.0, -4.0, 3.0, PV{4.0, -4.0, 3.0, 0.0}},
	}
	for _, table := range tables {
		vector := MakeVector(table.x, table.y, table.z)
		if vector != table.vector {
			t.Errorf("Expected %v, got %v", table.vector, vector)
		}
	}
}

func TestPVAdd(t *testing.T) {
	tables := []struct {
		p PV
		q PV
		s PV
	}{
		{PV{X:3, Y:-2, Z:5, W:1.0}, PV{X:-2, Y:3, Z:1, W:0.0}, PV{X:1, Y:1, Z:6, W:1.0}},
		{PV{X:-4, Y:7, Z:2, W:1.0}, PV{X:3, Y:1, Z:1, W:0.0}, PV{X:-1, Y:8, Z:3, W:1.0}},
	}
	for _, table := range tables {
		sum := table.p.Add(table.q)
		if sum != table.s {
			t.Errorf("Expected %v, got %v", table.s, sum)
		}
	}
}

func TestPVSubtract(t *testing.T) {
	tables := []struct {
		p PV
		q PV
		s PV
	}{
		{PV{X:3, Y:2, Z:1, W:1.0}, PV{X:5, Y:6, Z:7, W:0.0}, PV{X:-2, Y:-4, Z:-6, W:1.0}},
		{PV{X:3, Y:-2, Z:5, W:0.0}, PV{X:-2, Y:3, Z:1, W:0.0}, PV{X:5, Y:-5, Z:4, W:0.0}},
	}
	for _, table := range tables {
		diff := table.p.Subtract(table.q)
		if diff != table.s {
			t.Errorf("Expected %v, got %v", table.s, diff)
		}
	}
}

func TestPVNegate(t *testing.T) {
	tables := []struct {
		v PV
		n PV
	}{
		{PV{X:1, Y:-2, Z:3, W:0.0}, PV{X:-1, Y:2, Z:-3, W:0.0}},
	}
	for _, table := range tables {
		negative := table.v.Negate()
		if negative != table.n {
			t.Errorf("Expected %v, got %v", table.n, negative)
		}
	}
}

func TestPVScalar(t *testing.T) {
	tables := []struct {
		v PV
		s float64
		r PV
	}{
		{PV{X:1, Y:-2, Z:3, W:1.0}, 3.5, PV{X:3.5, Y:-7, Z:10.5, W:3.5}},
		{PV{X:1, Y:-2, Z:3, W:1.0}, 0.5, PV{X:0.5, Y:-1, Z:1.5, W:0.5}},
	}
	for _, table := range tables {
		scalar := table.v.Scalar(table.s)
		if scalar != table.r {
			t.Errorf("Expected %v, got %v", table.r, scalar)
		}
	}
}

func TestPVMagnitude(t *testing.T) {
	tables := []struct {
		v PV
		m float64
	}{
		{PV{X:1, Y:0, Z:0, W:0}, 1},
		{PV{X:0, Y:1, Z:0, W:0}, 1},
		{PV{X:0, Y:0, Z:1, W:0}, 1},
		{PV{X:1, Y:2, Z:3, W:0}, math.Sqrt(14)},
		{PV{X:-1, Y:-2, Z:-3, W:0}, math.Sqrt(14)},
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
		v PV
		n PV
	}{
		{PV{X:4, Y:0, Z:0, W:0}, PV{X:1, Y:0, Z:0, W:0}},
		{PV{X:1, Y:2, Z:3, W:0}, PV{X:0.26726, Y:0.53452, Z:0.80178, W:0}},
	}
	for _, table := range tables {
		normal := table.v.Normalize()
		if (math.Abs(table.n.X - normal.X) > 0.00001) || (math.Abs(table.n.Y - normal.Y) > 0.00001) || (math.Abs(table.n.Z - normal.Z) > 0.00001) {
			t.Errorf("Expected %v, got %v", table.n, normal)
		}
	}
}

func TestPVDotProduct(t *testing.T) {
	tables := []struct {
		v PV
		u PV
		d float64
	}{
		{PV{X:1, Y:2, Z:3, W:0}, PV{X:2, Y:3, Z:4, W:0}, 20},
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
		v PV
		u PV
		c PV
	}{
		{PV{X:1, Y:0, Z:0, W:0}, PV{X:0, Y:1, Z:0, W:0}, PV{X:0, Y:0, Z:1, W:0}},
		{PV{X:0, Y:1, Z:0, W:0}, PV{X:0, Y:0, Z:1, W:0}, PV{X:1, Y:0, Z:0, W:0}},
		{PV{X:0, Y:0, Z:1, W:0}, PV{X:1, Y:0, Z:0, W:0}, PV{X:0, Y:1, Z:0, W:0}},
	}
	for _, table := range tables {
		cross := table.v.CrossProduct(table.u)
		if cross != table.c {
			t.Errorf("Expected %v, got %v", table.c, cross)
		}
	}
}