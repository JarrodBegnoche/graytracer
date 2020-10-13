package components

import (
	"testing"
	"github.com/factorion/graytracer/pkg/primitives"
)

func TestMakeCamera(t *testing.T) {
	tables := []struct {
		width, height uint
		fieldOfView, pixelSize float64
	}{
		{200, 125, 1.5707963267948966, 0.01},
		{125, 200, 1.5707963267948966, 0.01},
	}
	for _, table := range tables {
		camera := MakeCamera(table.width, table.height, table.fieldOfView)
		if camera.width != table.width {
			t.Errorf("Incorrect width, expected %v, got %v", table.width, camera.width)
		}
		if camera.height != table.height {
			t.Errorf("Incorrect wiheightdth, expected %v, got %v", table.height, camera.height)
		}
		if camera.fieldOfView != table.fieldOfView {
			t.Errorf("Incorrect fieldOfView, expected %v, got %v", table.fieldOfView, camera.fieldOfView)
		}
		if camera.pixelSize != table.pixelSize {
			t.Errorf("Incorrect pixelSize, expected %v, got %v", table.pixelSize, camera.pixelSize)
		}
	}
}

func TestViewTransform(t *testing.T) {
	tables := []struct {
		from, to, up primitives.PV
		transform primitives.Matrix
	}{
		{primitives.MakePoint(0, 0, 0), primitives.MakePoint(0, 0, -1), primitives.MakeVector(0, 1, 0),
		 primitives.MakeIdentityMatrix(4)},

		{primitives.MakePoint(0, 0, 0), primitives.MakePoint(0, 0, 1), primitives.MakeVector(0, 1, 0),
		 primitives.Scaling(-1, 1, -1)},

		{primitives.MakePoint(0, 0, 8), primitives.MakePoint(0, 0, 0), primitives.MakeVector(0, 1, 0),
		 primitives.Translation(0, 0, -8)},
	}
	for _, table := range tables {
		camera := MakeCamera(100, 100, 1)
		camera.ViewTransform(table.from, table.to, table.up)
		if !camera.transform.Equals(table.transform) {
			t.Errorf("Expect %v, got %v", table.transform, camera.transform)
		}
	}
}

func TestRayForPixel(t *testing.T) {
	tables := []struct {
		x, y uint
		c *Camera
		transform primitives.Matrix
		ray primitives.Ray
	}{
		{100, 50, MakeCamera(201, 101, 1.5707963267948966), primitives.MakeIdentityMatrix(4),
		 primitives.Ray{Origin:primitives.MakePoint(0, 0, 0), Direction:primitives.MakeVector(0, 0, -1)}},

		{0, 0, MakeCamera(201, 101, 1.5707963267948966), primitives.MakeIdentityMatrix(4),
		 primitives.Ray{Origin:primitives.MakePoint(0, 0, 0),
						Direction:primitives.MakeVector(0.66518642611945, 0.332593213059725, -0.66851235825004)}},

		{100, 50, MakeCamera(201, 101, 1.5707963267948966),
		 primitives.RotationY(0.7853981633974483).Multiply(primitives.Translation(0, -2, 5)),
		 primitives.Ray{Origin:primitives.MakePoint(0, 2, -5),
						Direction:primitives.MakeVector(0.7071067811865476, 0, -0.7071067811865476)}},
	}
	for _, table := range tables {
		table.c.transform = table.transform
		ray := table.c.RayForPixel(table.x, table.y)
		if !ray.Equals(table.ray) {
			t.Errorf("Expected %v, got %v", table.ray, ray)
		}
	}
}
