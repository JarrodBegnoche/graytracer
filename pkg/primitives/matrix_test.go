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
		matrix1 Matrix
		matrix2 Matrix
		equals bool
	}{
		{Matrix{2, []float64{3, 2, 1, 0}}, Matrix{3, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}}, false},
		{Matrix{2, []float64{3, 2, 1, 0}}, Matrix{2, []float64{3, 2, 1, 1}}, false},
		{Matrix{2, []float64{1, 0, 0, 1}}, Matrix{2, []float64{1, 0, 0, 1}}, true},
	}
	for _, table := range tables {
		result := table.matrix1.Equals(table.matrix2)
		if result != table.equals {
			t.Errorf("Matrix %v and %v returned %v for Equals", table.matrix1, table.matrix2, result)
		}
	}
}

func TestMatrixGet(t *testing.T) {
	tables := []struct {
		matrix Matrix
		row uint8
		column uint8
		value float64
	}{
		{Matrix{2, []float64{3, 5, 1, 2}}, 1, 1, 2},
		{Matrix{2, []float64{3, 5, 1, 2}}, 1, 0, 1},
		{Matrix{3, []float64{3, 5, 0, 1, -2, -7, 0, 1, 1}}, 0, 0, 3},
	}
	for _, table := range tables {
		value := table.matrix.Get(table.row, table.column)
		if value != table.value {
			t.Errorf("Got %v, expected %v", value, table.value)
		}
	}
}

func TestMatrixSet(t *testing.T) {
	tables := []struct {
		matrix Matrix
		row uint8
		column uint8
		value float64
		pos uint8
	}{
		{Matrix{2, []float64{3, 5, 1, 2}}, 1, 1, 2, 3},
		{Matrix{2, []float64{3, 5, 1, 2}}, 1, 0, 2, 2},
		{Matrix{3, []float64{3, 5, 0, 1, -2, -7, 0, 1, 1}}, 0, 0, 3, 0},
	}
	for _, table := range tables {
		table.matrix.Set(table.row, table.column, table.value)
		if  table.matrix.Values[table.pos] != table.value {
			t.Errorf("Got %v, expected %v", table.matrix.Values[table.pos], table.value)
		}
	}
}

func TestMatrixMultiply(t *testing.T) {
	tables := []struct {
		matrix1 Matrix
		matrix2 Matrix
		product Matrix
	}{
		{Matrix{4, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 8, 7, 6, 5, 4, 3, 2}},
		 Matrix{4, []float64{-2, 1, 2, 3, 3, 2, 1, -1, 4, 3, 6, 5, 1, 2, 7, 8}},
		 Matrix{4, []float64{20, 22, 50, 48, 44, 54, 114, 108, 40, 58, 110, 102, 16, 26, 46, 42}}},
	}
	for _, table := range tables {
		product := table.matrix1.Multiply(table.matrix2)
		if !product.Equals(table.product) {
			t.Errorf("Expected %v, got %v", table.product, product)
		}
	}
}

func TestMatrixMultiplyPV(t *testing.T) {
	tables := []struct {
		matrix Matrix
		pv PV
		product PV
	}{
		{Matrix{4, []float64{1,2,3,4,2,4,4,2,8,6,4,1,0,0,0,1}},PV{1,2,3,1},PV{18,24,33,1}},
	}
	for _, table := range tables {
		product := table.matrix.MultiplyPV(table.pv)
		if product != table.product {
			t.Errorf("Expect %v, got %v", table.product, product)
		}
	}
}

func TestMatrixTranspose(t *testing.T) {
	tables := []struct {
		matrix Matrix
		transpose Matrix
	}{
		{Matrix{4, []float64{0,9,3,0,9,8,0,8,1,8,5,3,0,0,5,8}},
		 Matrix{4, []float64{0,9,1,0,9,8,8,0,3,0,5,5,0,8,3,8}}},
		 
		{Matrix{4, []float64{1,0,0,0,0,1,0,0,0,0,1,0,0,0,0,1}},
		 Matrix{4, []float64{1,0,0,0,0,1,0,0,0,0,1,0,0,0,0,1}}},
	}
	for _, table := range tables {
		transpose := table.matrix.Transpose()
		if !transpose.Equals(table.transpose) {
			t.Errorf("Expected %v, got %v", table.transpose, transpose)
		}
	}
}