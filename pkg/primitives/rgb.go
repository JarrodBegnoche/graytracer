package primitives

// RGB represents red, green, and blue values for a color object
type RGB struct {red, green, blue float64}

// Red Returns red value
func (r RGB) Red() float64 {
	return r.red
}

// Green Returns green value
func (r RGB) Green() float64 {
	return r.green
}

// Blue Returns blue value
func (r RGB) Blue() float64 {
	return r.blue
}

// Add Adds one RGB color to another and returns as a new RGB object
func (r RGB) Add(g RGB) RGB {
	return RGB{r.red + g.Red(), r.green + g.Green(), r.blue + g.Blue()}
}

// Subtract Subtracts one RGB color from another and returns as a new RGB object
func (r RGB) Subtract(g RGB) RGB {
	return RGB{r.red - g.Red(), r.green - g.Green(), r.blue - g.Blue()}
}

// Multiply Multiples one RGB color to another and returns as a new RGB object
func (r RGB) Multiply(g RGB) RGB {
	return RGB{r.red * g.Red(), r.green * g.Green(), r.blue * g.Blue()}
}

// Scale Scale an RGB color by a single value and return as a new RGB object
func (r RGB) Scale(s float64) RGB {
	return RGB{r.red * s, r.green * s, r.blue * s}
}
