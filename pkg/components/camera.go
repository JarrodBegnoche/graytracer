package components

import (
	"math"
	"github.com/factorion/graytracer/pkg/primitives"
)

// Camera Camera object to determine the view of the render
type Camera struct {
	width, height uint
	fieldOfView, halfWidth, halfHeight, pixelSize float64
	transform primitives.Matrix
}

// MakeCamera Create a camera object from the width, height, and field of view
func MakeCamera(width, height uint, fieldOfView float64) *Camera {
	c := Camera{width:width, height:height, fieldOfView:fieldOfView,
				transform:primitives.MakeIdentityMatrix(4)}
	halfView := math.Tan(fieldOfView / 2.0)
	aspect := float64(width) / float64(height)
	if aspect >= 1 {
		c.halfWidth = halfView
		c.halfHeight = halfView / aspect
	} else {
		c.halfWidth = halfView * aspect
		c.halfHeight = halfView
	}
	c.pixelSize = (c.halfWidth * 2.0) / float64(c.width)
	return &c
}

// ViewTransform Create a view transformation matrix for camera usage
func (c *Camera) ViewTransform(from, to, up primitives.PV) {
	orientation := primitives.MakeIdentityMatrix(4)
	forward := to.Subtract(from).Normalize()
	left := forward.CrossProduct(up.Normalize())
	trueUp := left.CrossProduct(forward)
	orientation[0][0] = left.X
	orientation[0][1] = left.Y
	orientation[0][2] = left.Z
	orientation[1][0] = trueUp.X
	orientation[1][1] = trueUp.Y
	orientation[1][2] = trueUp.Z
	orientation[2][0] = -forward.X
	orientation[2][1] = -forward.Y
	orientation[2][2] = -forward.Z
	c.transform = orientation.Multiply(primitives.Translation(-from.X, -from.Y, -from.Z))
}

// RayForPixel Calculate the ray for the given x, y coordinates
func (c Camera) RayForPixel(x, y uint) primitives.Ray {
	inverse, _ := c.transform.Inverse()
	pixel := primitives.MakePoint(c.halfWidth - ((float64(x) + 0.5) * c.pixelSize),
								  c.halfHeight - ((float64(y) + 0.5) * c.pixelSize), -1).Transform(inverse)
	origin := primitives.MakePoint(0, 0, 0).Transform(inverse)
	return primitives.Ray{Origin:origin, Direction:pixel.Subtract(origin).Normalize()}
}
