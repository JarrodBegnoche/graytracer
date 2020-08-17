package primitives

import (
	"math"
)

// Vector represents a vector in 3D space
type Vector struct {X, Y, Z float64}

// Add adds one vector to another and returns the result
func (v Vector) Add(q Vector) Vector {
	return Vector{v.X + q.X, v.Y + q.Y, v.Z + q.Z}
}

// Subtract subtracts one vector from another and returns the result
func (v Vector) Subtract(q Vector) Vector {
	return Vector{v.X - q.X, v.Y - q.Y, v.Z - q.Z}
}

// Negate Negate the vector to return one going in the opposite direction
func (v Vector) Negate() Vector {
	return Vector{0 - v.X, 0 - v.Y, 0 - v.Z}
}

// Scalar Scale a vector by a given value and return the result as a Vector
func (v Vector) Scalar(s float64) Vector {
	return Vector{v.X * s, v.Y * s, v.Z * s}
}

// Magnitude Returns the magnitude of the vector
func (v Vector) Magnitude() float64 {
	return math.Sqrt((v.X * v.X) + (v.Y * v.Y) + (v.Z * v.Z))
}

// Normalize Returns the normalized version of the vector
func (v Vector) Normalize() Vector {
	magnitude := v.Magnitude()
	return Vector{v.X / magnitude, v.Y / magnitude, v.Z / magnitude}
}

// DotProduct Return the dot product with the passed in vector
func (v Vector) DotProduct(q Vector) float64 {
	return (v.X * q.X) + (v.Y * q.Y) + (v.Z * q.Z)
}

// CrossProduct Returns the cross product with the vector passed in as a vector
func (v Vector) CrossProduct(q Vector) Vector {
	return Vector{(v.Y * q.Z) - (v.Z * q.Y),
				  (v.Z * q.X) - (v.X * q.Z),
				  (v.X * q.Y) - (v.Y * q.X)}
}