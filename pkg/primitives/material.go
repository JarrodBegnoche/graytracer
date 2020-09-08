package primitives

// Material Basic Phong material
type Material struct {
	Color RGB
	Ambient, Diffuse, Specular, Shininess float64
}
