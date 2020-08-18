package primitives

import (
	"testing"
)

func TestMakeEmptyMatrix(t *testing.T) {
	tables := []struct {
		size uint8
	}{
		{2},
		{3},
	}
	for _, table := range tables {
		matrix := MakeEmptyMatrix(table.size)
		for x := uint8(0); x < (table.size * table.size); x++ {
			if matrix.Values[x] != 0 {
				t.Errorf("Expected empty matrix, got %v", matrix.Values)
			}
		}
	}
}

func TestMakeIdentityMatrix(t *testing.T) {
	tables := []struct {
		size uint8
	}{
		{2},
		{3},
	}
	for _, table := range tables {
		matrix := MakeIdentityMatrix(table.size)
		for x := uint8(0); x < table.size; x++ {
			if matrix.Values[(x * table.size) + x] != 1.0 {
				t.Errorf("Expected empty matrix, got %v", matrix.Values)
			}
		}
	}
}

func TestMakeMatrix(t *testing.T) {
	tables := []struct {
		size uint8
		values []float64
		matrix Matrix
	}{
		{2, []float64{-3, 5, 1, -2}, Matrix{2, []float64{-3, 5, 1, -2}}},
	}
	for _, table := range tables {
		matrix := MakeMatrix(table.size, table.values)
		if !matrix.Equals(table.matrix) {
			t.Errorf("Created matrix %v did not match expected matrix %v", matrix, table.matrix)
		}
	}
}

func TestMatrixEquals(t *testing.T) {
	tables := []struct {
		m1 Matrix
		m2 Matrix
		equals bool
	}{
		{Matrix{2, []float64{3, 2, 1, 0}}, Matrix{3, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}}, false},
		{Matrix{2, []float64{3, 2, 1, 0}}, Matrix{2, []float64{3, 2, 1, 1}}, false},
		{Matrix{2, []float64{1, 0, 0, 1}}, Matrix{2, []float64{1, 0, 0, 1}}, true},
	}
	for _, table := range tables {
		result := table.m1.Equals(table.m2)
		if result != table.equals {
			t.Errorf("Matrix %v and %v returned %v for Equals", table.m1, table.m2, result)
		}
	}
}

func TestMatrixGet(t *testing.T) {
	tables := []struct {
		m Matrix
		x uint8
		y uint8
		value float64
	}{
		{Matrix{2, []float64{3, 5, 1, 2}}, 1, 1, 2},
		{Matrix{3, []float64{3, 5, 0, 1, -2, -7, 0, 1, 1}}, 0, 0, 3},
	}
	for _, table := range tables {
		value, _ := table.m.Get(table.x, table.y)
		if value != table.value {
			t.Errorf("Got %v, expected %v", value, table.value)
		}
	}
}

func TestMatrixSet(t *testing.T) {
	tables := []struct {
		m Matrix
		x uint8
		y uint8
		value float64
		pos uint8
	}{
		{Matrix{2, []float64{3, 5, 1, 2}}, 1, 1, 2, 3},
		{Matrix{3, []float64{3, 5, 0, 1, -2, -7, 0, 1, 1}}, 0, 0, 3, 0},
	}
	for _, table := range tables {
		table.m.Set(table.x, table.y, table.value)
		if  table.m.Values[table.pos] != table.value {
			t.Errorf("Got %v, expected %v", table.m.Values[table.pos], table.value)
		}
	}
}