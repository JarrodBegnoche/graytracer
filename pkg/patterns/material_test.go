package patterns

import (
	"testing"
	"github.com/factorion/graytracer/pkg/primitives"
)

func TestMakeMaterial(t *testing.T) {
	tables := []struct {
		mat Material
		col *RGB
		ambient, diffuse, specular, shininess float64
	}{
		{MakeDefaultMaterial(), MakeRGB(1, 1, 1), 0.1, 0.9, 0.9, 200},
	}
	for _, table := range tables {
		col := table.mat.Pat.ColorAt(primitives.MakePoint(0, 0, 0))
		if !col.Equals(*table.col) {
			t.Errorf("Color: Expected %v, got %v", table.col, col)
		}
		if table.mat.Ambient != table.ambient {
			t.Errorf("Ambient: Expected %v, got %v", table.ambient, table.mat.Ambient)
		}
		if table.mat.Diffuse != table.diffuse {
			t.Errorf("Diffuse: Expected %v, got %v", table.diffuse, table.mat.Diffuse)
		}
		if table.mat.Specular != table.specular {
			t.Errorf("Specular: Expected %v, got %v", table.specular, table.mat.Specular)
		}
		if table.mat.Shininess != table.shininess {
			t.Errorf("Shininess: Expected %v, got %v", table.shininess, table.mat.Shininess)
		}
	}
}