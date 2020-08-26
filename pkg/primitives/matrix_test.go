package primitives

import (
	"testing"
)

func TestMakeIdentitymatrix(t *testing.T) {
	tables := []struct {
		size uint8
	}{
		{2},
		{3},
		{4},
	}
	for _, table := range tables {
		matrix := MakeIdentityMatrix(table.size)
		for x := uint8(0); x < table.size; x++ {
			if matrix[x][x] != 1.0 {
				t.Errorf("Expected identity matrix, got %v", matrix)
			}
		}
	}
}

func TestMatrixEquals(t *testing.T) {
	tables := []struct {
		matrix1 matrix
		matrix2 matrix
		equals bool
	}{
		{matrix{{0, 1, 2, 3}, {4, 5, 6, 7}, {8, 9, 8, 7}, {6, 5, 4, 3}},
		 matrix{{0, 1, 2, 3}, {4, 5, 6, 7}, {8, 9, 8, 7}, {6, 5, 4, 3}}, true},

		{matrix{{0, 1, 2, 3}, {4, 5, 6, 7}, {8, 9, 8, 7}, {6, 5, 4, 3}},
		 matrix{{3, 2, 1, 0}, {4, 5, 6, 7}, {8, 9, 8, 7}, {6, 5, 4, 3}}, false},

		{matrix{{0, 1, 2, 3}, {4, 5, 6, 7}, {8, 9, 8, 7}, {6, 5, 4, 3}},
		 matrix{{0.000000001, 1, 2, 3}, {4, 5, 6, 7}, {8, 9, 8, 7}, {6, 5, 4, 3}}, true},

		{MakeMatrix(2), MakeMatrix(3), false},
	}
	for _, table := range tables {
		result := table.matrix1.Equals(table.matrix2)
		if result != table.equals {
			t.Errorf("Matrix %v and %v returned %v for Equals", table.matrix1, table.matrix2, result)
		}
	}
}

func TestSubmatrix(t * testing.T) {
	tables := []struct {
		matrix matrix
		row uint8
		column uint8
		submatrix matrix
	}{
		{matrix{{-6, 1, 1, 6}, {-8, 5, 8, 6}, {-1, 0, 8, 2}, {-7, 1, -1, 1}},
		 2, 1, matrix{{-6, 1, 6}, {-8, 8, 6}, {-7, -1, 1}}},
		
		{matrix{{1, 5, 0}, {-3, 2, 7}, {0, 6, -3}}, 0, 2, matrix{{-3, 2}, {0, 6}}},
	}
	for _, table := range tables {
		submatrix := table.matrix.Submatrix(table.row, table.column)
		if !submatrix.Equals(table.submatrix) {
			t.Errorf("Expected %v, got %v", table.submatrix, submatrix)
		}
	}
}

func TestMatrixMultiply(t *testing.T) {
	tables := []struct {
		matrix1 matrix
		matrix2 matrix
		product matrix
	}{
		{matrix{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 8, 7, 6}, {5, 4, 3, 2}},
		 matrix{{-2, 1, 2, 3}, {3, 2, 1, -1}, {4, 3, 6, 5}, {1, 2, 7, 8}},
		 matrix{{20, 22, 50, 48}, {44, 54, 114, 108}, {40, 58, 110, 102}, {16, 26, 46, 42}}},

		{MakeMatrix(2), MakeMatrix(3), nil},
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
		matrix1 matrix
		p, product pv
	}{
		{matrix{{1, 2, 3, 4}, {2, 4, 4, 2}, {8, 6, 4, 1}, {0, 0, 0, 1}},
		 MakePoint(1, 2, 3), MakePoint(18, 24, 33)},

		{MakeMatrix(3), MakeVector(4, 3, 2), MakeVector(0, 0, 0)},
	}
	for _, table := range tables {
		product := table.matrix1.MultiplyPV(table.p)
		if product != table.product {
			t.Errorf("Expect %v, got %v", table.product, product)
		}
	}
}

func TestMatrixTranspose(t *testing.T) {
	tables := []struct {
		matrix1 matrix
		transpose matrix
	}{
		{matrix{{0, 9, 3, 0}, {9, 8, 0, 8}, {1, 8, 5, 3}, {0, 0, 5, 8}},
		 matrix{{0, 9, 1, 0}, {9, 8, 8, 0}, {3, 0, 5, 5}, {0, 8, 3, 8}}},
		 
		{MakeIdentityMatrix(4), MakeIdentityMatrix(4)},
	}
	for _, table := range tables {
		transpose := table.matrix1.Transpose()
		if !transpose.Equals(table.transpose) {
			t.Errorf("Expected %v, got %v", table.transpose, transpose)
		}
	}
}

func TestMatrixDeterminant(t *testing.T) {
	tables := []struct {
		matrix1 matrix
		determinant float64
	}{
		{matrix{{1, 5}, {-3, 2}}, 17},
		{matrix{{1, 2, 6}, {-5, 8, -4}, {2, 6, 4}}, -196},
		{matrix{{-2, -8, 3, 5}, {-3, 1, 7, 3}, {1, 2, -9, 6}, {-6, 7, 7, -9}}, -4071},
	}
	for _, table := range tables {
		determinant := table.matrix1.Determinant()
		if determinant != table.determinant {
			t.Errorf("Expected %v, got %v", table.determinant, determinant)
		}
	}
}

func TestMatrixMinor(t *testing.T) {
	tables := []struct {
		matrix1 matrix
		row uint8
		column uint8
		minor float64
	}{
		{matrix{{3, 5, 0}, {2, -1, -7}, {6, -1, 5}}, 1, 0, 25},
		{MakeMatrix(2), 0, 0, 0},
	}
	for _, table := range tables {
		minor := table.matrix1.Minor(table.row, table.column)
		if minor != table.minor {
			t.Errorf("Expected %v, got %v", table.minor, minor)
		}
	}
}

func TestMatrixCofactor(t *testing.T) {
	tables := []struct {
		matrix1 matrix
		row uint8
		column uint8
		cofactor float64
	}{
		{matrix{{3, 5, 0}, {2, -1, -7}, {6, -1, 5}}, 0, 0, -12},
		{matrix{{3, 5, 0}, {2, -1, -7}, {6, -1, 5}}, 1, 0, -25},
	}
	for _, table := range tables {
		cofactor := table.matrix1.Cofactor(table.row, table.column)
		if cofactor != table.cofactor {
			t.Errorf("Expected %v, got %v", table.cofactor, cofactor)
		}
	}
}

func TestNonInvertibleMatrix(t *testing.T) {
	tables := []struct {
		matrix1 matrix
	}{
		{matrix{{-4, 2, -2, 3}, {9, 6, 2, 6}, {0, -5, 1, -5}, {0, 0, 0, 0}}},
	}
	for _, table := range tables {
		if _, ok := table.matrix1.Inverse(); ok == nil {
			t.Error("Failed, matrix is invertible")
		}
	}
}

func TestMatrixInverse(t *testing.T) {
	tables := []struct {
		matrix1 matrix
		inverse matrix
	}{
		{matrix{{-5, 2, 6, -8}, {1, -5, 1, 8}, {7, 7, -6, -7}, {1, -3, 7, 4}},
		 matrix{{0.21804511278195488, 0.45112781954887216, 0.24060150375939848, -0.045112781954887216},
		        {-0.8082706766917294, -1.4567669172932332, -0.44360902255639095, 0.5206766917293233},
				{-0.07894736842105263, -0.2236842105263158, -0.05263157894736842, 0.19736842105263158},
				{-0.5225563909774437, -0.8139097744360902, -0.3007518796992481, 0.30639097744360905}}},
		
		{matrix{{8, -5, 9, 2}, {7, 5, 6, 1}, {-6, 0, 9, 6}, {-3, 0, -9, -4}},
		 matrix{{-0.15384615384615385, -0.15384615384615385, -0.28205128205128205, -0.5384615384615384},
				{-0.07692307692307693, 0.12307692307692308, 0.02564102564102564, 0.03076923076923077},
				{0.358974358974359, 0.358974358974359, 0.4358974358974359, 0.9230769230769231},
				{-0.6923076923076923, -0.6923076923076923, -0.7692307692307693, -1.9230769230769231}}},
		
		{matrix{{9, 3, 0, 9}, {-5, -2, -6, -3}, {-4, 9, 6, 4}, {-7, 6, 6, 2}},
		 matrix{{-0.040740740740740744, -0.07777777777777778, 0.14444444444444443, -0.2222222222222222},
				{-0.07777777777777778, 0.03333333333333333, 0.36666666666666664, -0.3333333333333333},
				{-0.029012345679012345, -0.14629629629629629, -0.10925925925925926, 0.12962962962962962},
				{0.17777777777777778, 0.06666666666666667, -0.26666666666666666, 0.3333333333333333}}},
	}
	for _, table := range tables {
		inverse, ok := table.matrix1.Inverse()
		if ok != nil || !inverse.Equals(table.inverse) {
			t.Errorf("\nExpect %v, \ngot %v", table.inverse, inverse)
		}
	}
}

//Test
func TestMatrixProcess(t *testing.T) {
	tables := []struct {
		matrix1 matrix
		matrix2 matrix
	}{
		{matrix{{3, -9, 7, 3}, {3, -8, 2, -9}, {-4, 4, 4, 1}, {-6, 5, -1, 1}},
		 matrix{{8, 2, 2, 2}, {3, -1, 7, 0}, {7, 0, 5, 4}, {6, -2, 0, 5}}},
	}
	// A * B * B' = A
	for _, table := range tables {
		matrix3 := table.matrix1.Multiply(table.matrix2)
		inverse, ok := table.matrix2.Inverse()
		if ok != nil {
			t.Error("Matrix was not invertible")
		}
		matrix1 := matrix3.Multiply(inverse)
		if !matrix1.Equals(table.matrix1) {
			t.Errorf("Expected %v, got %v", table.matrix1, matrix1)
		}
	}
}