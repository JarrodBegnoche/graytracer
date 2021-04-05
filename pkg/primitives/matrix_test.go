package primitives_test

import (
	"testing"
	"github.com/factorion/graytracer/pkg/primitives"
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
		matrix := primitives.MakeIdentityMatrix(table.size)
		for x := uint8(0); x < table.size; x++ {
			if matrix[x][x] != 1.0 {
				t.Errorf("Expected identity matrix, got %v", matrix)
			}
		}
	}
}

func TestMatrixEquals(t *testing.T) {
	tables := []struct {
		matrix1, matrix2 primitives.Matrix
		equals bool
	}{
		{primitives.Matrix{{0, 1, 2, 3}, {4, 5, 6, 7}, {8, 9, 8, 7}, {6, 5, 4, 3}},
		primitives.Matrix{{0, 1, 2, 3}, {4, 5, 6, 7}, {8, 9, 8, 7}, {6, 5, 4, 3}}, true},

		{primitives.Matrix{{0, 1, 2, 3}, {4, 5, 6, 7}, {8, 9, 8, 7}, {6, 5, 4, 3}},
		primitives.Matrix{{3, 2, 1, 0}, {4, 5, 6, 7}, {8, 9, 8, 7}, {6, 5, 4, 3}}, false},

		{primitives.Matrix{{0, 1, 2, 3}, {4, 5, 6, 7}, {8, 9, 8, 7}, {6, 5, 4, 3}},
		primitives.Matrix{{0.000000001, 1, 2, 3}, {4, 5, 6, 7}, {8, 9, 8, 7}, {6, 5, 4, 3}}, true},

		{primitives.MakeMatrix(2), primitives.MakeMatrix(3), false},
	}
	for _, table := range tables {
		result := table.matrix1.Equals(table.matrix2)
		if result != table.equals {
			t.Errorf("Matrix %v and %v returned %v for Equals", table.matrix1, table.matrix2, result)
		}
	}
}

func TestSubmatrix(t *testing.T) {
	tables := []struct {
		matrix primitives.Matrix
		row uint8
		column uint8
		submatrix primitives.Matrix
	}{
		{primitives.Matrix{{-6, 1, 1, 6}, {-8, 5, 8, 6}, {-1, 0, 8, 2}, {-7, 1, -1, 1}},
		 2, 1, primitives.Matrix{{-6, 1, 6}, {-8, 8, 6}, {-7, -1, 1}}},
		
		{primitives.Matrix{{1, 5, 0}, {-3, 2, 7}, {0, 6, -3}}, 0, 2, primitives.Matrix{{-3, 2}, {0, 6}}},
	}
	for _, table := range tables {
		submatrix := table.matrix.Submatrix(table.row, table.column)
		if !submatrix.Equals(table.submatrix) {
			t.Errorf("Expected %v, got %v", table.submatrix, submatrix)
		}
	}
}

func BenchmarkSubmatrix4x4(b *testing.B) {
	matrix := primitives.Matrix{{-6, 1, 1, 6}, {-8, 5, 8, 6}, {-1, 0, 8, 2}, {-7, 1, -1, 1}}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		matrix.Submatrix(2, 1)
	}
}

func TestMatrixMultiply(t *testing.T) {
	tables := []struct {
		matrix1, matrix2, product primitives.Matrix
	}{
		{primitives.Matrix{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 8, 7, 6}, {5, 4, 3, 2}},
		 primitives.Matrix{{-2, 1, 2, 3}, {3, 2, 1, -1}, {4, 3, 6, 5}, {1, 2, 7, 8}},
		 primitives.Matrix{{20, 22, 50, 48}, {44, 54, 114, 108}, {40, 58, 110, 102}, {16, 26, 46, 42}}},

		{primitives.MakeMatrix(2), primitives.MakeMatrix(3), nil},
	}
	for _, table := range tables {
		product := table.matrix1.Multiply(table.matrix2)
		if !product.Equals(table.product) {
			t.Errorf("Expected %v, got %v", table.product, product)
		}
	}
}

func BenchmarkMatrixMultiply(b *testing.B) {
	matrix1 := primitives.Matrix{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 8, 7, 6}, {5, 4, 3, 2}}
	matrix2 := primitives.Matrix{{-2, 1, 2, 3}, {3, 2, 1, -1}, {4, 3, 6, 5}, {1, 2, 7, 8}}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		matrix1.Multiply(matrix2)
	}
}

func TestMatrixTranspose(t *testing.T) {
	tables := []struct {
		matrix1, transpose primitives.Matrix
	}{
		{primitives.Matrix{{0, 9, 3, 0}, {9, 8, 0, 8}, {1, 8, 5, 3}, {0, 0, 5, 8}},
		 primitives.Matrix{{0, 9, 1, 0}, {9, 8, 8, 0}, {3, 0, 5, 5}, {0, 8, 3, 8}}},
		 
		{primitives.MakeIdentityMatrix(4), primitives.MakeIdentityMatrix(4)},
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
		matrix1 primitives.Matrix
		determinant float64
	}{
		{primitives.Matrix{{1, 5}, {-3, 2}}, 17},
		{primitives.Matrix{{1, 2, 6}, {-5, 8, -4}, {2, 6, 4}}, -196},
		{primitives.Matrix{{-2, -8, 3, 5}, {-3, 1, 7, 3}, {1, 2, -9, 6}, {-6, 7, 7, -9}}, -4071},
	}
	for _, table := range tables {
		determinant := table.matrix1.Determinant()
		if determinant != table.determinant {
			t.Errorf("Expected %v, got %v", table.determinant, determinant)
		}
	}
}

func BenchmarkMatrixDeterminant4x4(b *testing.B) {
	matrix := primitives.Matrix{{-2, -8, 3, 5}, {-3, 1, 7, 3}, {1, 2, -9, 6}, {-6, 7, 7, -9}}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		matrix.Determinant()
	}
}

func TestMatrixMinor(t *testing.T) {
	tables := []struct {
		matrix1 primitives.Matrix
		row, column uint8
		minor float64
	}{
		{primitives.Matrix{{3, 5, 0}, {2, -1, -7}, {6, -1, 5}}, 1, 0, 25},
		{primitives.MakeMatrix(2), 0, 0, 0},
	}
	for _, table := range tables {
		minor := table.matrix1.Minor(table.row, table.column)
		if minor != table.minor {
			t.Errorf("Expected %v, got %v", table.minor, minor)
		}
	}
}

func BenchmarkMatrixMinor3x3(b *testing.B) {
	matrix := primitives.Matrix{{3, 5, 0}, {2, -1, -7}, {6, -1, 5}}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		matrix.Minor(1, 0)
	}
}

func TestMatrixCofactor(t *testing.T) {
	tables := []struct {
		matrix1 primitives.Matrix
		row, column uint8
		cofactor float64
	}{
		{primitives.Matrix{{3, 5, 0}, {2, -1, -7}, {6, -1, 5}}, 0, 0, -12},
		{primitives.Matrix{{3, 5, 0}, {2, -1, -7}, {6, -1, 5}}, 1, 0, -25},
	}
	for _, table := range tables {
		cofactor := table.matrix1.Cofactor(table.row, table.column)
		if cofactor != table.cofactor {
			t.Errorf("Expected %v, got %v", table.cofactor, cofactor)
		}
	}
}

func BenchmarkMatrixCofactor3x3(b *testing.B) {
	matrix := primitives.Matrix{{3, 5, 0}, {2, -1, -7}, {6, -1, 5}}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		matrix.Cofactor(0, 0)
	}
}

func TestNonInvertibleMatrix(t *testing.T) {
	tables := []struct {
		matrix1 primitives.Matrix
	}{
		{primitives.Matrix{{-4, 2, -2, 3}, {9, 6, 2, 6}, {0, -5, 1, -5}, {0, 0, 0, 0}}},
	}
	for _, table := range tables {
		if _, ok := table.matrix1.Inverse(); ok == nil {
			t.Error("Failed, matrix is invertible")
		}
	}
}

func TestMatrixInverse(t *testing.T) {
	tables := []struct {
		matrix1, inverse primitives.Matrix
	}{
		{primitives.Matrix{{-5, 2, 6, -8}, {1, -5, 1, 8}, {7, 7, -6, -7}, {1, -3, 7, 4}},
		 primitives.Matrix{{0.21804511278195488, 0.45112781954887216, 0.24060150375939848, -0.045112781954887216},
		        {-0.8082706766917294, -1.4567669172932332, -0.44360902255639095, 0.5206766917293233},
				{-0.07894736842105263, -0.2236842105263158, -0.05263157894736842, 0.19736842105263158},
				{-0.5225563909774437, -0.8139097744360902, -0.3007518796992481, 0.30639097744360905}}},
		
		{primitives.Matrix{{8, -5, 9, 2}, {7, 5, 6, 1}, {-6, 0, 9, 6}, {-3, 0, -9, -4}},
		 primitives.Matrix{{-0.15384615384615385, -0.15384615384615385, -0.28205128205128205, -0.5384615384615384},
				{-0.07692307692307693, 0.12307692307692308, 0.02564102564102564, 0.03076923076923077},
				{0.358974358974359, 0.358974358974359, 0.4358974358974359, 0.9230769230769231},
				{-0.6923076923076923, -0.6923076923076923, -0.7692307692307693, -1.9230769230769231}}},
		
		{primitives.Matrix{{9, 3, 0, 9}, {-5, -2, -6, -3}, {-4, 9, 6, 4}, {-7, 6, 6, 2}},
		 primitives.Matrix{{-0.040740740740740744, -0.07777777777777778, 0.14444444444444443, -0.2222222222222222},
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

func BenchmarkMatrixInverse4x4(b *testing.B) {
	matrix := primitives.Matrix{{-5, 2, 6, -8}, {1, -5, 1, 8}, {7, 7, -6, -7}, {1, -3, 7, 4}}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		matrix.Inverse()
	}
}

func TestMatrixProcess(t *testing.T) {
	tables := []struct {
		matrix1, matrix2 primitives.Matrix
	}{
		{primitives.Matrix{{3, -9, 7, 3}, {3, -8, 2, -9}, {-4, 4, 4, 1}, {-6, 5, -1, 1}},
		 primitives.Matrix{{8, 2, 2, 2}, {3, -1, 7, 0}, {7, 0, 5, 4}, {6, -2, 0, 5}}},
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
