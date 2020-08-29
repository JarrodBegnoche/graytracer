package primitives

// Ray Contains a point and vector
type Ray struct {
	Origin, Direction PV
}

// Position Calculate ray's position after a set amount of time
func (ray Ray) Position(time float64) PV {
	return ray.Origin.Add(ray.Direction.Scalar(time))
}
