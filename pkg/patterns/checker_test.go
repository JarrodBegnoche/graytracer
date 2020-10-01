package patterns

import (
	"testing"
	"github.com/factorion/graytracer/pkg/primitives"
)

func TestCheckerColorAt(t *testing.T) {
	tables := []struct {
		c *Checker
		transform primitives.Matrix
		point primitives.PV
		result *RGB
	}{
		{MakeChecker(MakeRGB(1, 1, 1), MakeRGB(0, 0, 0)), primitives.MakeIdentityMatrix(4),
		 primitives.MakePoint(0, 0, 0), MakeRGB(1, 1, 1)},
		
		{MakeChecker(MakeRGB(1, 1, 1), MakeRGB(0, 0, 0)), primitives.MakeIdentityMatrix(4),
		 primitives.MakePoint(0.99, 0, 0), MakeRGB(1, 1, 1)},
		
		{MakeChecker(MakeRGB(1, 1, 1), MakeRGB(0, 0, 0)), primitives.MakeIdentityMatrix(4),
		 primitives.MakePoint(1.01, 0, 0), MakeRGB(0, 0, 0)},
		
		{MakeChecker(MakeRGB(1, 1, 1), MakeRGB(0, 0, 0)), primitives.MakeIdentityMatrix(4),
		 primitives.MakePoint(0, 0.99, 0), MakeRGB(1, 1, 1)},
		
		 {MakeChecker(MakeRGB(1, 1, 1), MakeRGB(0, 0, 0)), primitives.MakeIdentityMatrix(4),
		  primitives.MakePoint(0, 1.01, 0), MakeRGB(0, 0, 0)},
		
		{MakeChecker(MakeRGB(1, 1, 1), MakeRGB(0, 0, 0)), primitives.Scaling(0.5, 0.5, 0),
		 primitives.MakePoint(0.51, 0, 0), MakeRGB(0, 0, 0)},
	}
	for _, table := range tables {
		table.c.SetTransform(table.transform)
		result := table.c.ColorAt(table.point)
		if !result.Equals(*table.result) {
			t.Errorf("Point: %v, Expected %v, got %v", table.point, table.result, result)
		}
	}
}