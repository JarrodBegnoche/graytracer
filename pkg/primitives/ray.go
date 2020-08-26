package primitives

// Ray Contains a point and vector
type Ray struct {
	origin, direction pv
}

// Position Calculate ray's position after a set amount of time
func (ray Ray) Position(time float64) pv {
	return ray.origin.Add(ray.direction.Scalar(time))
}
