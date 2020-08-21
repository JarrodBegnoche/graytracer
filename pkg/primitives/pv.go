package primitives

import (
	"math"
)


// PV represents 3D coordinates and a w variable for distinction between point and vector
type PV struct {x, y, z, w float64}

// X Return X coordinate
func (pv PV) X() float64 {
	return pv.x
}

// Y Return Y coordinate
func (pv PV) Y() float64 {
	return pv.y
}

// Z Return Z coordinate
func (pv PV) Z() float64 {
	return pv.z
}

// W Return W value
func (pv PV) W() float64 {
	return pv.w
}

// MakeVector Create a vector PV type
func MakeVector(x, y, z float64) PV {
	return PV{x:x, y:y, z:z, w:0.0}
}

// MakePoint Create a point PV type
func MakePoint(x, y, z float64) PV {
	return PV{x:x, y:y, z:z, w:1.0}
}

// Add adds one PV to another and returns the result
func (pv PV) Add(o PV) PV {
	return PV{pv.x + o.X(), pv.y + o.Y(), pv.z + o.Z(), pv.w + o.W()}
}

// Subtract subtracts one PV from another and returns the result
func (pv PV) Subtract(o PV) PV {
	return PV{pv.x - o.X(), pv.y - o.Y(), pv.z - o.Z(), pv.w - o.W()}
}

// Negate Negate the PV to return its opposite
func (pv PV) Negate() PV {
	return PV{0 - pv.x, 0 - pv.y, 0 - pv.z, pv.w}
}

// Scalar Scale a PV by a given value and return the result as a PV
func (pv PV) Scalar(s float64) PV {
	return PV{pv.x * s, pv.y * s, pv.z * s, pv.w * s}
}

// Magnitude Returns the magnitude of the PV
func (pv PV) Magnitude() float64 {
	return math.Sqrt((pv.x * pv.x) + (pv.y * pv.y) + (pv.z * pv.z) + (pv.w * pv.w))
}

// Normalize Returns the normalized version of the PV
func (pv PV) Normalize() PV {
	magnitude := pv.Magnitude()
	return PV{pv.x / magnitude, pv.y / magnitude, pv.z / magnitude, pv.w / magnitude}
}

// DotProduct Return the dot product with the passed in PV
func (pv PV) DotProduct(o PV) float64 {
	return (pv.x * o.X()) + (pv.y * o.Y()) + (pv.z * o.Z()) + (pv.w * o.W())
}

// CrossProduct Returns the cross product with the PV passed in as a PV
func (pv PV) CrossProduct(o PV) PV {
	return PV{(pv.y * o.Z()) - (pv.z * o.Y()),
			  (pv.z * o.X()) - (pv.x * o.Z()),
			  (pv.x * o.Y()) - (pv.y * o.X()),
			  0.0}
}