package patterns_test

import (
	"testing"
	"github.com/factorion/graytracer/pkg/patterns"
	"github.com/factorion/graytracer/pkg/primitives"
)

func TestGradientColorAt(t *testing.T) {
	tables := []struct {
		g *patterns.Gradient
		transform primitives.Matrix
		point primitives.PV
		result *patterns.RGB
	}{
		{patterns.MakeGradient(patterns.MakeRGB(1, 1, 1), patterns.MakeRGB(0, 0, 0)),
		 primitives.MakeIdentityMatrix(4),
		 primitives.MakePoint(0, 0, 0),
		 patterns.MakeRGB(1, 1, 1)},
		
		{patterns.MakeGradient(patterns.MakeRGB(1, 1, 1), patterns.MakeRGB(0, 0, 0)),
		 primitives.MakeIdentityMatrix(4),
		 primitives.MakePoint(0.25, 0, 0), patterns.MakeRGB(0.75, 0.75, 0.75)},
		
		{patterns.MakeGradient(patterns.MakeRGB(1, 1, 1), patterns.MakeRGB(0, 0, 0)),
		 primitives.MakeIdentityMatrix(4),
		 primitives.MakePoint(0.5, 0, 0),
		 patterns.MakeRGB(0.5, 0.5, 0.5)},
		
		{patterns.MakeGradient(patterns.MakeRGB(1, 1, 1), patterns.MakeRGB(0, 0, 0)),
		 primitives.MakeIdentityMatrix(4),
		 primitives.MakePoint(0.75, 0, 0),
		 patterns.MakeRGB(0.25, 0.25, 0.25)},
		
		{patterns.MakeGradient(patterns.MakeRGB(1, 1, 1), patterns.MakeRGB(0, 0, 0)),
		 primitives.Scaling(0.5, 0, 0),
		 primitives.MakePoint(0.75, 0, 0),
		 patterns.MakeRGB(0.5, 0.5, 0.5)},
	}
	for _, table := range tables {
		table.g.SetTransform(table.transform)
		result := table.g.ColorAt(table.point)
		if !result.Equals(*table.result) {
			t.Errorf("Point: %v, Expected %v, got %v", table.point, table.result, result)
		}
	}
}
