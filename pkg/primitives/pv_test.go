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
		{PV{x:3, y:-2, z:5, w:1.0}, PV{x:-2, y:3, z:1, w:0.0}, PV{x:1, y:1, z:6, w:1.0}},
		{PV{x:-4, y:7, z:2, w:1.0}, PV{x:3, y:1, z:1, w:0.0}, PV{x:-1, y:8, z:3, w:1.0}},
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
		{PV{x:3, y:2, z:1, w:1.0}, PV{x:5, y:6, z:7, w:0.0}, PV{x:-2, y:-4, z:-6, w:1.0}},
		{PV{x:3, y:-2, z:5, w:0.0}, PV{x:-2, y:3, z:1, w:0.0}, PV{x:5, y:-5, z:4, w:0.0}},
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
		{PV{x:1, y:-2, z:3, w:0.0}, PV{x:-1, y:2, z:-3, w:0.0}},
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
		{PV{x:1, y:-2, z:3, w:1.0}, 3.5, PV{x:3.5, y:-7, z:10.5, w:3.5}},
		{PV{x:1, y:-2, z:3, w:1.0}, 0.5, PV{x:0.5, y:-1, z:1.5, w:0.5}},
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
		{PV{x:1, y:0, z:0, w:0}, 1},
		{PV{x:0, y:1, z:0, w:0}, 1},
		{PV{x:0, y:0, z:1, w:0}, 1},
		{PV{x:1, y:2, z:3, w:0}, math.Sqrt(14)},
		{PV{x:-1, y:-2, z:-3, w:0}, math.Sqrt(14)},
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
		{PV{x:4, y:0, z:0, w:0}, PV{x:1, y:0, z:0, w:0}},
		{PV{x:1, y:2, z:3, w:0}, PV{x:0.26726, y:0.53452, z:0.80178, w:0}},
	}
	for _, table := range tables {
		normal := table.v.Normalize()
		if (math.Abs(table.n.X() - normal.X()) > 0.00001) || (math.Abs(table.n.Y() - normal.Y()) > 0.00001) ||
		   (math.Abs(table.n.Z() - normal.Z()) > 0.00001) {
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
		{PV{x:1, y:2, z:3, w:0}, PV{x:2, y:3, z:4, w:0}, 20},
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
		{PV{x:1, y:0, z:0, w:0}, PV{x:0, y:1, z:0, w:0}, PV{x:0, y:0, z:1, w:0}},
		{PV{x:0, y:1, z:0, w:0}, PV{x:0, y:0, z:1, w:0}, PV{x:1, y:0, z:0, w:0}},
		{PV{x:0, y:0, z:1, w:0}, PV{x:1, y:0, z:0, w:0}, PV{x:0, y:1, z:0, w:0}},
	}
	for _, table := range tables {
		cross := table.v.CrossProduct(table.u)
		if cross != table.c {
			t.Errorf("Expected %v, got %v", table.c, cross)
		}
	}
}