package primitives

// RGB represents red, green, and blue values for a color object
type RGB struct {Red, Green, Blue float64}

// Add Adds one RGB color to another and returns a new RGB object
func (r RGB) Add(g RGB) RGB {
	return RGB{r.Red + g.Red, r.Green + g.Green, r.Blue + g.Blue}
}

// Subtract Subtracts one RGB color from another and returns a new RGB object
func (r RGB) Subtract(g RGB) RGB {
	return RGB{r.Red - g.Red, r.Green - g.Green, r.Blue - g.Blue}
}

// Multiply Multiples one RGB color to another and returns a new RGB object
func (r RGB) Multiply(g RGB) RGB {
	return RGB{r.Red * g.Red, r.Green * g.Green, r.Blue * g.Blue}
}

// Scale Scale an RGB color by a single value and return the new RGB object
func (r RGB) Scale(s float64) RGB {
	return RGB{r.Red * s, r.Green * s, r.Blue * s}
}
