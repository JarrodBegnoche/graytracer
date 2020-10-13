package patterns

// Material Basic Phong material
type Material struct {
	Pat Pattern
	Ambient, Diffuse, Specular, Shininess, Reflective, Transparency, RefractiveIndex float64
}

// MakeDefaultMaterial Create a basic material
func MakeDefaultMaterial() Material {
	return Material{Pat:MakeRGB(1, 1, 1), Ambient:0.1, Diffuse:0.9, Specular:0.9, Shininess:200,
					Reflective:0, Transparency:0, RefractiveIndex:1}
}
