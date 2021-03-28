package shapes

import (
	"math"
	"sort"
	"github.com/factorion/graytracer/pkg/primitives"
)

// MinMax Sort an array of floats and return the minimum and maximum values
func MinMax(values []float64) (float64, float64) {
	sort.Float64s(values)
	return values[0], values[len(values) - 1]
}

// CombineBounds Combine a slice of bounds into one
func CombineBounds(bounds_slice []*Bounds) *Bounds {
	if len(bounds_slice) == 0 {
		return nil
	}
	x_list := make([]float64, len(bounds_slice) * 2)
	y_list := make([]float64, len(bounds_slice) * 2)
	z_list := make([]float64, len(bounds_slice) * 2)
	for index, bounds := range bounds_slice {
		x_list[index * 2] = bounds.Min.X
		x_list[(index * 2) + 1] = bounds.Max.X
		y_list[index * 2] = bounds.Min.Y
		y_list[(index * 2) + 1] = bounds.Max.Y
		z_list[index * 2] = bounds.Min.Z
		z_list[(index * 2) + 1] = bounds.Max.Z
	}
	x_min, x_max := MinMax(x_list)
	y_min, y_max := MinMax(y_list)
	z_min, z_max := MinMax(z_list)
	return &Bounds{Min:primitives.MakePoint(x_min, y_min, z_min), Max:primitives.MakePoint(x_max, y_max, z_max)}
}

// AddBounds Add bounds with another bounds
func (b* Bounds) AddBounds(bounds *Bounds) {
	x_min := math.Min(b.Min.X, bounds.Min.X)
	x_max := math.Max(b.Max.X, bounds.Max.X)
	y_min := math.Min(b.Min.Y, bounds.Min.Y)
	y_max := math.Max(b.Max.Y, bounds.Max.Y)
	z_min := math.Min(b.Min.Z, bounds.Min.Z)
	z_max := math.Max(b.Max.Z, bounds.Max.Z)
	b.Min = primitives.MakePoint(x_min, y_min, z_min)
	b.Max = primitives.MakePoint(x_max, y_max, z_max)
}

// Bounds Bounding box structure with a minimum and maximum point representing an axis-aligned cube
type Bounds struct {
	Min, Max primitives.PV
}

func (b* Bounds) Transform(transform primitives.Matrix) *Bounds {
	x_list := make([]float64, 8)
	y_list := make([]float64, 8)
	z_list := make([]float64, 8)
	points := make([]primitives.PV, 8)
	// Create the eight points of a cube
	points[0] = primitives.MakePoint(b.Min.X, b.Min.Y, b.Min.Z)
	points[1] = primitives.MakePoint(b.Min.X, b.Min.Y, b.Max.Z)
	points[2] = primitives.MakePoint(b.Min.X, b.Max.Y, b.Min.Z)
	points[3] = primitives.MakePoint(b.Min.X, b.Max.Y, b.Max.Z)
	points[4] = primitives.MakePoint(b.Max.X, b.Min.Y, b.Min.Z)
	points[5] = primitives.MakePoint(b.Max.X, b.Min.Y, b.Max.Z)
	points[6] = primitives.MakePoint(b.Max.X, b.Max.Y, b.Min.Z)
	points[7] = primitives.MakePoint(b.Max.X, b.Max.Y, b.Max.Z)
	// Transform the eight points and add their coordinates to their respective slices
	for i := 0; i < 8; i++ {
		point := points[i].Transform(transform)
		x_list[i] = point.X
		y_list[i] = point.Y
		z_list[i] = point.Z
	}
	// Get the minimum and maximum, resulting in our new bounds
	x_min, x_max := MinMax(x_list)
	y_min, y_max := MinMax(y_list)
	z_min, z_max := MinMax(z_list)
	return &Bounds{Min:primitives.MakePoint(x_min, y_min, z_min), Max:primitives.MakePoint(x_max, y_max, z_max)}
}

func (b* Bounds) Intersect(ray primitives.Ray) bool {
	xtmin, xtmax := CheckAxis(ray.Origin.X, ray.Direction.X, b.Min.X, b.Max.X)
	ytmin, ytmax := CheckAxis(ray.Origin.Y, ray.Direction.Y, b.Min.Y, b.Max.Y)
	ztmin, ztmax := CheckAxis(ray.Origin.Z, ray.Direction.Z, b.Min.Z, b.Max.Z)
	tmin := math.Max(math.Max(xtmin, ytmin), ztmin)
	tmax := math.Min(math.Min(xtmax, ytmax), ztmax)
	return tmin <= tmax;
}
