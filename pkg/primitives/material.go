package primitives

// Material Basic Phong material
type Material struct {
	Color RGB
	Ambient, Diffuse, Specular, Shininess float64
}

func MakeDefaultMaterial() Material {
	return Material{Color:MakeRGB(1, 1, 1), Ambient:0.1, Diffuse:0.9, Specular:0.9, Shininess:200}
}