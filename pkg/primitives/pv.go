package primitives

import (
	"math"
)


// PV represents a PV in 3D space
type PV struct {X, Y, Z, W float64}

// MakeVector Create a vector PV type
func MakeVector(x, y, z float64) PV {
	return PV{x, y, z, 0.0}
}

// MakePoint Create a point PV type
func MakePoint(x, y, z float64) PV {
	return PV{x, y, z, 1.0}
}

// Add adds one PV to another and returns the result
func (v PV) Add(q PV) PV {
	return PV{v.X + q.X, v.Y + q.Y, v.Z + q.Z, v.W + q.W}
}

// Subtract subtracts one PV from another and returns the result
func (v PV) Subtract(q PV) PV {
	return PV{v.X - q.X, v.Y - q.Y, v.Z - q.Z, v.W - q.W}
}

// Negate Negate the PV to return one going in the opposite direction
func (v PV) Negate() PV {
	return PV{0 - v.X, 0 - v.Y, 0 - v.Z, v.W}
}

// Scalar Scale a PV by a given value and return the result as a PV
func (v PV) Scalar(s float64) PV {
	return PV{v.X * s, v.Y * s, v.Z * s, v.W * s}
}

// Magnitude Returns the magnitude of the PV
func (v PV) Magnitude() float64 {
	return math.Sqrt((v.X * v.X) + (v.Y * v.Y) + (v.Z * v.Z) + (v.W * v.W))
}

// Normalize Returns the normalized version of the PV
func (v PV) Normalize() PV {
	magnitude := v.Magnitude()
	return PV{v.X / magnitude, v.Y / magnitude, v.Z / magnitude, v.W / magnitude}
}

// DotProduct Return the dot product with the passed in PV
func (v PV) DotProduct(q PV) float64 {
	return (v.X * q.X) + (v.Y * q.Y) + (v.Z * q.Z) + (v.W * q.W)
}

// CrossProduct Returns the cross product with the PV passed in as a PV
func (v PV) CrossProduct(q PV) PV {
	return PV{(v.Y * q.Z) - (v.Z * q.Y),
			  (v.Z * q.X) - (v.X * q.Z),
			  (v.X * q.Y) - (v.Y * q.X),
			  0.0}
}