package primitives

import (
	"math"
)

// PV represents 3D coordinates and a w variable for distinction between point and vector
type PV struct {x, y, z, w float64}

// MakeVector Create a vector PV type
func MakeVector(x, y, z float64) PV {
	return PV{x:x, y:y, z:z, w:0.0}
}

// MakePoint Create a point PV type
func MakePoint(x, y, z float64) PV {
	return PV{x:x, y:y, z:z, w:1.0}
}

// Equals Compares two PVs with an amount for approximation
func (p PV) Equals(o PV) bool {
	EPSILON := 0.00000001
	if math.Abs(p.x - o.x) > EPSILON {
		return false
	}
	if math.Abs(p.y - o.y) > EPSILON {
		return false
	}
	if math.Abs(p.z - o.z) > EPSILON {
		return false
	}
	if math.Abs(p.w - o.w) > EPSILON {
		return false
	}
	return true
}

// Add adds one PV to another and returns the result
func (p PV) Add(o PV) PV {
	return PV{p.x + o.x, p.y + o.y, p.z + o.z, p.w + o.w}
}

// Subtract subtracts one PV from another and returns the result
func (p PV) Subtract(o PV) PV {
	return PV{p.x - o.x, p.y - o.y, p.z - o.z, p.w - o.w}
}

// Negate Negate the PV to return its opposite
func (p PV) Negate() PV {
	return PV{-p.x, -p.y, -p.z, p.w}
}

// Scalar Scale a PV by a given value and return the result as a PV
func (p PV) Scalar(s float64) PV {
	return PV{p.x * s, p.y * s, p.z * s, p.w}
}

// Magnitude Returns the magnitude of the PV
func (p PV) Magnitude() float64 {
	return math.Sqrt((p.x * p.x) + (p.y * p.y) + (p.z * p.z) + (p.w * p.w))
}

// Normalize Returns the normalized version of the PV
func (p PV) Normalize() PV {
	magnitude := p.Magnitude()
	return PV{p.x / magnitude, p.y / magnitude, p.z / magnitude, p.w}
}

// DotProduct Return the dot product with the passed in PV
func (p PV) DotProduct(o PV) float64 {
	return (p.x * o.x) + (p.y * o.y) + (p.z * o.z) + (p.w * o.w)
}

// CrossProduct Returns the cross product with the PV passed in as a PV
func (p PV) CrossProduct(o PV) PV {
	return MakeVector((p.y * o.z) - (p.z * o.y),
			  		  (p.z * o.x) - (p.x * o.z),
			  		  (p.x * o.y) - (p.y * o.x))
}
