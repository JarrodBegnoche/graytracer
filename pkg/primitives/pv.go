package primitives

import (
	"math"
)

// PV represents 3D coordinates and a w variable for distinction between point and vector
type PV struct {X, Y, Z, W float64}

// MakeVector Create a vector PV type
func MakeVector(x, y, z float64) PV {
	return PV{X:x, Y:y, Z:z, W:0.0}
}

// MakePoint Create a point PV type
func MakePoint(x, y, z float64) PV {
	return PV{X:x, Y:y, Z:z, W:1.0}
}

// Equals Compares two PVs with an amount for approximation
func (p PV) Equals(o PV) bool {
	EPSILON := 0.00000001
	if math.Abs(p.X - o.X) > EPSILON {
		return false
	}
	if math.Abs(p.Y - o.Y) > EPSILON {
		return false
	}
	if math.Abs(p.Z - o.Z) > EPSILON {
		return false
	}
	if math.Abs(p.W - o.W) > EPSILON {
		return false
	}
	return true
}

// Add adds one PV to another and returns the result
func (p PV) Add(o PV) PV {
	return PV{p.X + o.X, p.Y + o.Y, p.Z + o.Z, p.W + o.W}
}

// Subtract subtracts one PV from another and returns the result
func (p PV) Subtract(o PV) PV {
	return PV{p.X - o.X, p.Y - o.Y, p.Z - o.Z, p.W - o.W}
}

// Transform Transform the PV by a matrix
func (p PV) Transform(m Matrix) PV {
	if len(m) != 4 {
		return PV{}
	}
	return PV{X:(m[0][0] * p.X) + (m[0][1] * p.Y) + (m[0][2] * p.Z) + (m[0][3] * p.W),
			  Y:(m[1][0] * p.X) + (m[1][1] * p.Y) + (m[1][2] * p.Z) + (m[1][3] * p.W),
			  Z:(m[2][0] * p.X) + (m[2][1] * p.Y) + (m[2][2] * p.Z) + (m[2][3] * p.W),
			  W:(m[3][0] * p.X) + (m[3][1] * p.Y) + (m[3][2] * p.Z) + (m[3][3] * p.W)}
}

// Negate Negate the PV to return its opposite
func (p PV) Negate() PV {
	return PV{-p.X, -p.Y, -p.Z, p.W}
}

// Scalar Scale a PV by a given value and return the result as a PV
func (p PV) Scalar(s float64) PV {
	return PV{p.X * s, p.Y * s, p.Z * s, p.W}
}

// Magnitude Returns the magnitude of the PV
func (p PV) Magnitude() float64 {
	return math.Sqrt((p.X * p.X) + (p.Y * p.Y) + (p.Z * p.Z) + (p.W * p.W))
}

// Normalize Returns the normalized version of the PV
func (p PV) Normalize() PV {
	magnitude := p.Magnitude()
	return PV{p.X / magnitude, p.Y / magnitude, p.Z / magnitude, p.W}
}

// DotProduct Return the dot product with the passed in PV
func (p PV) DotProduct(o PV) float64 {
	return (p.X * o.X) + (p.Y * o.Y) + (p.Z * o.Z) + (p.W * o.W)
}

// CrossProduct Returns the cross product with the PV passed in as a PV
func (p PV) CrossProduct(o PV) PV {
	return MakeVector((p.Y * o.Z) - (p.Z * o.Y),
			  		  (p.Z * o.X) - (p.X * o.Z),
			  		  (p.X * o.Y) - (p.Y * o.X))
}

// Reflect Calculate the reflection vector from a normal
func (p PV) Reflect(normal PV) PV {
	return p.Subtract(normal.Scalar(2 * p.DotProduct(normal)))
}
