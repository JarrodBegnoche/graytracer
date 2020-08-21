package primitives

import (
	"math"
)

// Matrix2 A 2x2 matrix
type Matrix2 [2][2]float64

// Matrix3 A 3x3 matrix
type Matrix3 [3][3]float64

// Matrix4 A 4x4 matrix
type Matrix4 [4][4]float64

// MakeIdentityMatrix4 Make an identity matrix4x4
func MakeIdentityMatrix4() Matrix4 {
	matrix := Matrix4{}
	for x := uint8(0); x < 4; x++ {
		matrix[x][x] = 1
	}
	return matrix
}

// Equals Compares two matrices with an amount for approximation
func (m Matrix4) Equals(o Matrix4) bool {
	EPSILON := 0.00000001
	for row := uint8(0); row < 4; row++ {
		for column := uint8(0); column < 4; column++ {
			if math.Abs(m[row][column] - o[row][column]) > EPSILON {
				return false
			}
		}
	}
	return true
}

// Multiply4 Matrix multiplication function
func (m Matrix4) Multiply4(o Matrix4) Matrix4 {
	matrix := Matrix4{}
	for row := uint8(0); row < 4; row++ {
		for column := uint8(0); column < 4; column++ {
			matrix[row][column] = (m[row][0] * o[0][column]) +
								  (m[row][1] * o[1][column]) +
								  (m[row][2] * o[2][column]) +
								  (m[row][3] * o[3][column])
		}
	}
	return matrix
}

// Multiply4PV Multiply a matrix by a point/vector
func (m Matrix4) Multiply4PV(pv PV) PV {
	return PV{x:(m[0][0] * pv.X()) + (m[0][1] * pv.Y()) + (m[0][2] * pv.Z()) + (m[0][3] * pv.W()),
			  y:(m[1][0] * pv.X()) + (m[1][1] * pv.Y()) + (m[1][2] * pv.Z()) + (m[1][3] * pv.W()),
			  z:(m[2][0] * pv.X()) + (m[2][1] * pv.Y()) + (m[2][2] * pv.Z()) + (m[2][3] * pv.W()),
			  w:(m[3][0] * pv.X()) + (m[3][1] * pv.Y()) + (m[3][2] * pv.Z()) + (m[3][3] * pv.W())}
}

// Transpose Flip the rows with the columns of a matrix
func (m Matrix4) Transpose() Matrix4 {
	matrix := Matrix4{}
	for row := uint8(0); row < 4; row++ {
		for column := uint8(0); column < 4; column++ {
			matrix[row][column] = m[column][row]
		}
	}
	return matrix
}