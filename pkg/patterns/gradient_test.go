package patterns

import (
	"testing"
	"github.com/factorion/graytracer/pkg/primitives"
)

func TestGradientColorAt(t *testing.T) {
	tables := []struct {
		g *Gradient
		transform primitives.Matrix
		point primitives.PV
		result *RGB
	}{
		{MakeGradient(MakeRGB(1, 1, 1), MakeRGB(0, 0, 0)), primitives.MakeIdentityMatrix(4),
		 primitives.MakePoint(0, 0, 0), MakeRGB(1, 1, 1)},
		
		{MakeGradient(MakeRGB(1, 1, 1), MakeRGB(0, 0, 0)), primitives.MakeIdentityMatrix(4),
		 primitives.MakePoint(0.25, 0, 0), MakeRGB(0.75, 0.75, 0.75)},
		
		{MakeGradient(MakeRGB(1, 1, 1), MakeRGB(0, 0, 0)), primitives.MakeIdentityMatrix(4),
		 primitives.MakePoint(0.5, 0, 0), MakeRGB(0.5, 0.5, 0.5)},
		
		{MakeGradient(MakeRGB(1, 1, 1), MakeRGB(0, 0, 0)), primitives.MakeIdentityMatrix(4),
		 primitives.MakePoint(0.75, 0, 0), MakeRGB(0.25, 0.25, 0.25)},
		
		{MakeGradient(MakeRGB(1, 1, 1), MakeRGB(0, 0, 0)), primitives.Scaling(0.5, 0, 0),
		 primitives.MakePoint(0.75, 0, 0), MakeRGB(0.5, 0.5, 0.5)},
	}
	for _, table := range tables {
		table.g.SetTransform(table.transform)
		result := table.g.ColorAt(table.point)
		if !result.Equals(*table.result) {
			t.Errorf("Point: %v, Expected %v, got %v", table.point, table.result, result)
		}
	}
}