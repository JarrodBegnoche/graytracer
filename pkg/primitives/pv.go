package primitives

import (
	"math"
)

// pv represents 3D coordinates and a w variable for distinction between point and vector
type pv struct {x, y, z, w float64}

// X Return X coordinate
func (p pv) X() float64 {
	return p.x
}

// Y Return Y coordinate
func (p pv) Y() float64 {
	return p.y
}

// Z Return Z coordinate
func (p pv) Z() float64 {
	return p.z
}

// W Return W value
func (p pv) W() float64 {
	return p.w
}

// MakeVector Create a vector pv type
func MakeVector(x, y, z float64) pv {
	return pv{x:x, y:y, z:z, w:0.0}
}

// MakePoint Create a point pv type
func MakePoint(x, y, z float64) pv {
	return pv{x:x, y:y, z:z, w:1.0}
}

// Equals Compares two pvs with an amount for approximation
func (p pv) Equals(o pv) bool {
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

// Add adds one pv to another and returns the result
func (p pv) Add(o pv) pv {
	return pv{p.x + o.X(), p.y + o.Y(), p.z + o.Z(), p.w + o.W()}
}

// Subtract subtracts one pv from another and returns the result
func (p pv) Subtract(o pv) pv {
	return pv{p.x - o.X(), p.y - o.Y(), p.z - o.Z(), p.w - o.W()}
}

// Negate Negate the pv to return its opposite
func (p pv) Negate() pv {
	return pv{-p.x, -p.y, -p.z, p.w}
}

// Scalar Scale a pv by a given value and return the result as a pv
func (p pv) Scalar(s float64) pv {
	return pv{p.x * s, p.y * s, p.z * s, p.w}
}

// Magnitude Returns the magnitude of the pv
func (p pv) Magnitude() float64 {
	return math.Sqrt((p.x * p.x) + (p.y * p.y) + (p.z * p.z) + (p.w * p.w))
}

// Normalize Returns the normalized version of the pv
func (p pv) Normalize() pv {
	magnitude := p.Magnitude()
	return pv{p.x / magnitude, p.y / magnitude, p.z / magnitude, p.w}
}

// DotProduct Return the dot product with the passed in pv
func (p pv) DotProduct(o pv) float64 {
	return (p.x * o.X()) + (p.y * o.Y()) + (p.z * o.Z()) + (p.w * o.W())
}

// CrossProduct Returns the cross product with the pv passed in as a pv
func (p pv) CrossProduct(o pv) pv {
	return MakeVector((p.y * o.Z()) - (p.z * o.Y()),
			  		  (p.z * o.X()) - (p.x * o.Z()),
			  		  (p.x * o.Y()) - (p.y * o.X()))
}
