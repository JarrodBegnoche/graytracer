package primitives

import (
	"errors"
	"math"
)

// Matrix2 A 2x2 matrix
type matrix [][]float64

// MakeMatrix Makes an emptry square matrix of the given size
func MakeMatrix(size uint8) matrix {
	m := make([][]float64, size)
	for x := uint8(0); x < size; x++ {
		m[x] = make([]float64, size)
	}
	return m
}

// MakeIdentityMatrix Make an identity matrix of the given size
func MakeIdentityMatrix(size uint8) matrix {
	m := make([][]float64, size)
	for x := uint8(0); x < size; x++ {
		m[x] = make([]float64, size)
		m[x][x] = 1
	}
	return m
}

// Equals Compares two matrices with an amount for approximation
func (m matrix) Equals(o matrix) bool {
	size := uint8(len(m))
	if size != uint8(len(o)) {
		return false
	}
	EPSILON := 0.00000001
	for row := uint8(0); row < size; row++ {
		for column := uint8(0); column < size; column++ {
			if math.Abs(m[row][column] - o[row][column]) > EPSILON {
				return false
			}
		}
	}
	return true
}

// Submatrix Return a smaller matrix that exists within
func (m matrix) Submatrix(row uint8, column uint8) matrix {
	size := uint8(len(m))
	submatrix := MakeMatrix(size - 1)
	for x := uint8(0); x < size; x++ {
		subX := x
		if subX == row {
			continue
		} else if subX > row {
			subX--
		}
		for y := uint8(0); y < size; y++ {
			subY := y
			if subY == column {
				continue
			} else if subY > column {
				subY--
			}
			submatrix[subX][subY] = m[x][y]
		}
	}
	return submatrix
}

// Multiply Matrix multiplication function
func (m matrix) Multiply(o matrix) matrix {
	size := uint8(len(m))
	if size != uint8(len(o)) {
		return nil
	}
	matrix := MakeMatrix(size)
	for row := uint8(0); row < size; row++ {
		for column := uint8(0); column < size; column++ {
			sum := float64(0)
			for val := uint8(0); val < size; val++ {
				sum += m[row][val] * o[val][column]
			}
			matrix[row][column] = sum
		}
	}
	return matrix
}

// MultiplyPV Multiply a matrix by a point/vector
func (m matrix) MultiplyPV(p pv) pv {
	if len(m) != 4 {
		return pv{}
	}
	return pv{x:(m[0][0] * p.X()) + (m[0][1] * p.Y()) + (m[0][2] * p.Z()) + (m[0][3] * p.W()),
			  y:(m[1][0] * p.X()) + (m[1][1] * p.Y()) + (m[1][2] * p.Z()) + (m[1][3] * p.W()),
			  z:(m[2][0] * p.X()) + (m[2][1] * p.Y()) + (m[2][2] * p.Z()) + (m[2][3] * p.W()),
			  w:(m[3][0] * p.X()) + (m[3][1] * p.Y()) + (m[3][2] * p.Z()) + (m[3][3] * p.W())}
}

// Transpose Flip the rows with the columns of a matrix
func (m matrix) Transpose() matrix {
	size := uint8(len(m))
	matrix := MakeMatrix(size)
	for row := uint8(0); row < size; row++ {
		for column := uint8(0); column < size; column++ {
			matrix[row][column] = m[column][row]
		}
	}
	return matrix
}

// Determinant Calculates the determinant of a 2x2 matrix
func (m matrix) Determinant() float64 {
	size := uint8(len(m))
	determinant := float64(0)
	if size == 2 {
		determinant = (m[0][0] * m[1][1]) - (m[0][1] * m[1][0])
	} else if size > 2 {
		for column := uint8(0); column < size; column++ {
			determinant += m[0][column] * m.Cofactor(0, column)
		}
	}
	return determinant
}

// Minor Calculate the determinant of the submatrix
func (m matrix) Minor(row, column uint8) float64 {
	return (m.Submatrix(row, column)).Determinant()
}

// Cofactor Calculate minor and negate if necessary
func (m matrix) Cofactor(row, column uint8) float64 {
	cofactor := m.Minor(row, column)
	if (row + column) % 2 == 1 {
		cofactor = -cofactor
	}
	return cofactor
}

// Inverse Invert the function if possible
func (m matrix) Inverse() (matrix, error) {
	determinant := m.Determinant()
	if determinant == 0 {
		return nil, errors.New("Not invertible")
	}
	size := uint8(len(m))
	inverse := MakeMatrix(size)
	for row := uint8(0); row < size; row++ {
		for column := uint8(0); column < size; column++ {
			inverse[column][row] = m.Cofactor(row, column) / determinant
		}
	}
	return inverse, nil
}
