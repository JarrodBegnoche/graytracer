package primitives

import (
	"testing"
)

func TestMakeIdentityMatrix4(t *testing.T) {
	matrix := MakeIdentityMatrix4()
	for x := uint8(0); x < 4; x++ {
		if matrix[x][x] != 1.0 {
			t.Errorf("Expected identity matrix, got %v", matrix)
		}
	}
}

func TestMatrix4Equals(t *testing.T) {
	tables := []struct {
		matrix1 Matrix4
		matrix2 Matrix4
		equals bool
	}{
		{Matrix4{{0, 1, 2, 3}, {4, 5, 6, 7}, {8, 9, 8, 7}, {6, 5, 4, 3}},
		 Matrix4{{0, 1, 2, 3}, {4, 5, 6, 7}, {8, 9, 8, 7}, {6, 5, 4, 3}}, true},

		{Matrix4{{0, 1, 2, 3}, {4, 5, 6, 7}, {8, 9, 8, 7}, {6, 5, 4, 3}},
		 Matrix4{{3, 2, 1, 0}, {4, 5, 6, 7}, {8, 9, 8, 7}, {6, 5, 4, 3}}, false},

		{Matrix4{{0, 1, 2, 3}, {4, 5, 6, 7}, {8, 9, 8, 7}, {6, 5, 4, 3}},
		 Matrix4{{0.000000001, 1, 2, 3}, {4, 5, 6, 7}, {8, 9, 8, 7}, {6, 5, 4, 3}}, true},
	}
	for _, table := range tables {
		result := table.matrix1.Equals(table.matrix2)
		if result != table.equals {
			t.Errorf("Matrix %v and %v returned %v for Equals", table.matrix1, table.matrix2, result)
		}
	}
}

func TestMatrixMultiply4(t *testing.T) {
	tables := []struct {
		matrix1 Matrix4
		matrix2 Matrix4
		product Matrix4
	}{
		{Matrix4{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 8, 7, 6}, {5, 4, 3, 2}},
		 Matrix4{{-2, 1, 2, 3}, {3, 2, 1, -1}, {4, 3, 6, 5}, {1, 2, 7, 8}},
		 Matrix4{{20, 22, 50, 48}, {44, 54, 114, 108}, {40, 58, 110, 102}, {16, 26, 46, 42}}},
	}
	for _, table := range tables {
		product := table.matrix1.Multiply4(table.matrix2)
		if !product.Equals(table.product) {
			t.Errorf("Expected %v, got %v", table.product, product)
		}
	}
}

func TestMatrixMultiplyPV(t *testing.T) {
	tables := []struct {
		matrix Matrix4
		pv PV
		product PV
	}{
		{Matrix4{{1, 2, 3, 4}, {2, 4, 4, 2}, {8, 6, 4, 1}, {0, 0, 0, 1}},
		 PV{1, 2, 3, 1}, PV{18, 24, 33, 1}},
	}
	for _, table := range tables {
		product := table.matrix.Multiply4PV(table.pv)
		if product != table.product {
			t.Errorf("Expect %v, got %v", table.product, product)
		}
	}
}

func TestMatrixTranspose(t *testing.T) {
	tables := []struct {
		matrix Matrix4
		transpose Matrix4
	}{
		{Matrix4{{0, 9, 3, 0}, {9, 8, 0, 8}, {1, 8, 5, 3}, {0, 0, 5, 8}},
		 Matrix4{{0, 9, 1, 0}, {9, 8, 8, 0}, {3, 0, 5, 5}, {0, 8, 3, 8}}},
		 
		{Matrix4{{1, 0, 0, 0}, {0, 1, 0, 0}, {0, 0, 1, 0}, {0, 0, 0, 1}},
		 Matrix4{{1, 0, 0, 0}, {0, 1, 0, 0}, {0, 0, 1, 0}, {0, 0, 0, 1}}},
	}
	for _, table := range tables {
		transpose := table.matrix.Transpose()
		if !transpose.Equals(table.transpose) {
			t.Errorf("Expected %v, got %v", table.transpose, transpose)
		}
	}
}
