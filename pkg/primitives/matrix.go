package primitives

import (
	"math"
)

// Matrix Represents a square matrix, usually 4x4
type Matrix struct {
	Size uint8
	Values []float64
}

// MakeEmptyMatrix Create a matrix of a specific size
func MakeEmptyMatrix(size uint8) Matrix {
	matrix := Matrix{Size: size, Values:make([]float64, size * size)}
	return matrix
}

// MakeIdentityMatrix Make an identity matrix
func MakeIdentityMatrix(size uint8) Matrix {
	matrix := Matrix{Size: size, Values:make([]float64, size * size)}
	for x := uint8(0); x < size; x++ {
		matrix.Set(x, x, 1)
	}
	return matrix
}

// MakeMatrix Create a matrix of a specific size and 
func MakeMatrix(size uint8, values []float64) Matrix {
	matrix := Matrix{Size: size, Values:make([]float64, size * size)}
	for x := uint8(0); x < (size * size); x++ {
		matrix.Values[x] = values[x]
	}
	return matrix
}

// Equals Compares two matrices with an amount for approximation
func (m Matrix) Equals(m2 Matrix) bool {
	EPSILON := 0.00000001
	if m.Size != m2.Size {
		return false
	}
	for x := uint8(0); x < (m.Size * m.Size); x++ {
		if math.Abs(m.Values[x] - m2.Values[x]) > EPSILON {
			return false
		}
	}
	return true
}

// Get Return the value at a given position in the matrix
func (m Matrix) Get(row, column uint8) float64 {
	return m.Values[(row * m.Size) + column]
}

// Set Set the value at a give position in the matrix
func (m Matrix) Set(row, column uint8, value float64) {
	m.Values[(row * m.Size) + column] = value
}

// Multiply Matrix multiplication function
func (m Matrix) Multiply(o Matrix) Matrix {
	matrix := MakeEmptyMatrix(m.Size)
	for row := uint8(0); row < m.Size; row++ {
		for column := uint8(0); column < m.Size; column++ {
			sum := float64(0)
			for val := uint8(0); val < m.Size; val++ {
				sum += m.Get(row, val) * o.Get(val, column)
			}
			matrix.Set(row, column, sum)
		}
	}
	return matrix
}

// MultiplyPV Multiply a matrix by a point/vector
func (m Matrix) MultiplyPV(pv PV) PV {
	return PV{X:(m.Get(0, 0) * pv.X) + (m.Get(0, 1) * pv.Y) + (m.Get(0, 2) * pv.Z) + (m.Get(0, 3) * pv.W),
			  Y:(m.Get(1, 0) * pv.X) + (m.Get(1, 1) * pv.Y) + (m.Get(1, 2) * pv.Z) + (m.Get(1, 3) * pv.W),
			  Z:(m.Get(2, 0) * pv.X) + (m.Get(2, 1) * pv.Y) + (m.Get(2, 2) * pv.Z) + (m.Get(2, 3) * pv.W),
			  W:(m.Get(3, 0) * pv.X) + (m.Get(3, 1) * pv.Y) + (m.Get(3, 2) * pv.Z) + (m.Get(3, 3) * pv.W)}
}

// Transpose Flip the rows with the columns of a matrix
func (m Matrix) Transpose() Matrix {
	matrix := MakeMatrix(m.Size, m.Values)
	for row := uint8(1); row < m.Size; row++ {
		for column := uint8(0); column < row; column++ {
			spot1 := (row * m.Size) + column
			spot2 := (column * m.Size) + row
			matrix.Values[spot1], matrix.Values[spot2] = matrix.Values[spot2], matrix.Values[spot1]
		}
	}
	return matrix
}