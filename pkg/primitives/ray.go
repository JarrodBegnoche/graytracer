package primitives

// Ray Contains a point and vector
type Ray struct {
	Origin, Direction PV
}

// Equals Compare two rays with an amount for approximation
func (ray Ray) Equals(o Ray) bool {
	return ray.Origin.Equals(o.Origin) && ray.Direction.Equals(o.Direction)
}

// Position Calculate ray's position after a set amount of time
func (ray Ray) Position(time float64) PV {
	return ray.Origin.Add(ray.Direction.Scalar(time))
}

// Transform Transform the origin and direciton by a matrix
func (ray Ray) Transform(m Matrix) Ray {
	if len(m) != 4 {
		return Ray{}
	}
	return Ray{Origin:ray.Origin.Transform(m), Direction:ray.Direction.Transform(m)}
}
