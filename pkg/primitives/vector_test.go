package primitives

import (
	"testing"
	"math"
)

func TestVectorAdd(t *testing.T) {
	tables := []struct {
		p Vector
		q Vector
		s Vector
	}{
		{Vector{X:3, Y:-2, Z:5}, Vector{X:-2, Y:3, Z:1}, Vector{X:1, Y:1, Z:6}},
		{Vector{X:-4, Y:7, Z:2}, Vector{X:3, Y:1, Z:1}, Vector{X:-1, Y:8, Z:3}},
	}
	for _, table := range tables {
		sum := table.p.Add(table.q)
		if sum != table.s {
			t.Errorf("Expected %v, got %v", table.s, sum)
		}
	}
}

func TestVectorSubtract(t *testing.T) {
	tables := []struct {
		p Vector
		q Vector
		s Vector
	}{
		{Vector{X:3, Y:2, Z:1}, Vector{X:5, Y:6, Z:7}, Vector{X:-2, Y:-4, Z:-6}},
		{Vector{X:3, Y:-2, Z:5}, Vector{X:-2, Y:3, Z:1}, Vector{X:5, Y:-5, Z:4}},
	}
	for _, table := range tables {
		diff := table.p.Subtract(table.q)
		if diff != table.s {
			t.Errorf("Expected %v, got %v", table.s, diff)
		}
	}
}

func TestVectorNegate(t *testing.T) {
	tables := []struct {
		v Vector
		n Vector
	}{
		{Vector{X:1, Y:-2, Z:3}, Vector{X:-1, Y:2, Z:-3}},
	}
	for _, table := range tables {
		negative := table.v.Negate()
		if negative != table.n {
			t.Errorf("Expected %v, got %v", table.n, negative)
		}
	}
}

func TestVectorScalar(t *testing.T) {
	tables := []struct {
		v Vector
		s float64
		r Vector
	}{
		{Vector{X:1, Y:-2, Z:3}, 3.5, Vector{X:3.5, Y:-7, Z:10.5}},
		{Vector{X:1, Y:-2, Z:3}, 0.5, Vector{X:0.5, Y:-1, Z:1.5}},
	}
	for _, table := range tables {
		scalar := table.v.Scalar(table.s)
		if scalar != table.r {
			t.Errorf("Expected %v, got %v", table.r, scalar)
		}
	}
}

func TestVectorMagnitude(t *testing.T) {
	tables := []struct {
		v Vector
		m float64
	}{
		{Vector{X:1, Y:0, Z:0}, 1},
		{Vector{X:0, Y:1, Z:0}, 1},
		{Vector{X:0, Y:0, Z:1}, 1},
		{Vector{X:1, Y:2, Z:3}, math.Sqrt(14)},
		{Vector{X:-1, Y:-2, Z:-3}, math.Sqrt(14)},
	}
	for _, table := range tables {
		magnitude := table.v.Magnitude()
		if magnitude != table.m {
			t.Errorf("Expected %v, got %v", table.m, magnitude)
		}
	}
}

func TestVectorNormalize(t *testing.T) {
	tables := []struct {
		v Vector
		n Vector
	}{
		{Vector{X:4, Y:0, Z:0}, Vector{X:1, Y:0, Z:0}},
		{Vector{X:1, Y:2, Z:3}, Vector{X:0.26726, Y:0.53452, Z:0.80178}},
	}
	for _, table := range tables {
		normal := table.v.Normalize()
		if (math.Abs(table.n.X - normal.X) > 0.00001) || (math.Abs(table.n.Y - normal.Y) > 0.00001) || (math.Abs(table.n.Z - normal.Z) > 0.00001) {
			t.Errorf("Expected %v, got %v", table.n, normal)
		}
	}
}

func TestVectorDotProduct(t *testing.T) {
	tables := []struct {
		v Vector
		u Vector
		d float64
	}{
		{Vector{X:1, Y:2, Z:3}, Vector{X:2, Y:3, Z:4}, 20},
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
		v Vector
		u Vector
		c Vector
	}{
		{Vector{X:1, Y:0, Z:0}, Vector{X:0, Y:1, Z:0}, Vector{X:0, Y:0, Z:1}},
		{Vector{X:0, Y:1, Z:0}, Vector{X:0, Y:0, Z:1}, Vector{X:1, Y:0, Z:0}},
		{Vector{X:0, Y:0, Z:1}, Vector{X:1, Y:0, Z:0}, Vector{X:0, Y:1, Z:0}},
	}
	for _, table := range tables {
		cross := table.v.CrossProduct(table.u)
		if cross != table.c {
			t.Errorf("Expected %v, got %v", table.c, cross)
		}
	}
}