package primitives

import (
	"math"
)

// Translation Move a point, not a vector
func Translation(x, y, z float64) Matrix {
	matrix := MakeIdentityMatrix(4)
	matrix[0][3] = x
	matrix[1][3] = y
	matrix[2][3] = z
	return matrix
}

// Scaling Scale the point/vector
func Scaling(x, y, z float64) Matrix {
	matrix := MakeIdentityMatrix(4)
	matrix[0][0] = x
	matrix[1][1] = y
	matrix[2][2] = z
	return matrix
}

// RotationX Rotate around the X-Axis
func RotationX(rad float64) Matrix {
	matrix := MakeIdentityMatrix(4)
	matrix[1][1] = math.Cos(rad)
	matrix[2][1] = math.Sin(rad)
	matrix[1][2] = -matrix[2][1]
	matrix[2][2] = matrix[1][1]
	return matrix
}

// RotationY Rotate around the Y-Axis
func RotationY(rad float64) Matrix {
	matrix := MakeIdentityMatrix(4)
	matrix[0][0] = math.Cos(rad)
	matrix[0][2] = math.Sin(rad)
	matrix[2][0] = -matrix[0][2]
	matrix[2][2] = matrix[0][0]
	return matrix
}

// RotationZ Rotate around the Z-Axis
func RotationZ(rad float64) Matrix {
	matrix := MakeIdentityMatrix(4)
	matrix[0][0] = math.Cos(rad)
	matrix[1][0] = math.Sin(rad)
	matrix[0][1] = -matrix[1][0]
	matrix[1][1] = matrix[0][0]
	return matrix
}

// Shearing Shear an axis along another axis
func Shearing(xy, xz, yx, yz, zx, zy float64) Matrix {
	matrix := MakeIdentityMatrix(4)
	matrix[0][1] = xy
	matrix[0][2] = xz
	matrix[1][0] = yx
	matrix[1][2] = yz
	matrix[2][0] = zx
	matrix[2][1] = zy
	return matrix
}
