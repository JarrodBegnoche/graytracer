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
func (m Matrix) Get(row, column uint8) (float64, error) {
	return m.Values[(row * m.Size) + column], nil
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
				a, _ := m.Get(row, val)
				b, _ := o.Get(val, column)
				sum += a * b
			}
			matrix.Set(row, column, sum)
		}
	}
	return matrix
}