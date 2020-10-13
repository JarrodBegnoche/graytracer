package patterns_test

import (
	"testing"
	"github.com/factorion/graytracer/pkg/patterns"
	"github.com/factorion/graytracer/pkg/primitives"
)

func TestMakeMaterial(t *testing.T) {
	tables := []struct {
		mat patterns.Material
		col *patterns.RGB
		ambient, diffuse, specular, shininess, reflective, transparency, refractiveIndex float64
	}{
		{patterns.MakeDefaultMaterial(), patterns.MakeRGB(1, 1, 1), 0.1, 0.9, 0.9, 200, 0, 0, 1},
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
		if table.mat.Reflective != table.reflective {
			t.Errorf("Reflective: Expected %v, got %v", table.reflective, table.mat.Reflective)
		}
		if table.mat.Transparency != table.transparency {
			t.Errorf("Transparency: Expected %v, got %v", table.transparency, table.mat.Transparency)
		}
		if table.mat.RefractiveIndex != table.refractiveIndex {
			t.Errorf("RefractiveIndex: Expected %v, got %v", table.refractiveIndex, table.mat.RefractiveIndex)
		}
	}
}
