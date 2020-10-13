package patterns_test

import (
	"testing"
	"github.com/factorion/graytracer/pkg/patterns"
	"github.com/factorion/graytracer/pkg/primitives"
)

func TestCheckerColorAt(t *testing.T) {
	tables := []struct {
		c *patterns.Checker
		transform primitives.Matrix
		point primitives.PV
		result *patterns.RGB
	}{
		{patterns.MakeChecker(patterns.MakeRGB(1, 1, 1), patterns.MakeRGB(0, 0, 0)),
		 primitives.MakeIdentityMatrix(4),
		 primitives.MakePoint(0, 0, 0),
		 patterns.MakeRGB(1, 1, 1)},
		
		{patterns.MakeChecker(patterns.MakeRGB(1, 1, 1), patterns.MakeRGB(0, 0, 0)),
		 primitives.MakeIdentityMatrix(4),
		 primitives.MakePoint(0.99, 0, 0),
		 patterns.MakeRGB(1, 1, 1)},
		
		{patterns.MakeChecker(patterns.MakeRGB(1, 1, 1), patterns.MakeRGB(0, 0, 0)),
		 primitives.MakeIdentityMatrix(4),
		 primitives.MakePoint(1.01, 0, 0),
		 patterns.MakeRGB(0, 0, 0)},
		
		{patterns.MakeChecker(patterns.MakeRGB(1, 1, 1), patterns.MakeRGB(0, 0, 0)),
		 primitives.MakeIdentityMatrix(4),
		 primitives.MakePoint(0, 0.99, 0),
		 patterns.MakeRGB(1, 1, 1)},
		
		{patterns.MakeChecker(patterns.MakeRGB(1, 1, 1), patterns.MakeRGB(0, 0, 0)),
		 primitives.MakeIdentityMatrix(4),
		 primitives.MakePoint(0, 1.01, 0),
		 patterns.MakeRGB(0, 0, 0)},
		
		{patterns.MakeChecker(patterns.MakeRGB(1, 1, 1), patterns.MakeRGB(0, 0, 0)),
		 primitives.Scaling(0.5, 0.5, 0),
		 primitives.MakePoint(0.51, 0, 0),
		 patterns.MakeRGB(0, 0, 0)},
	}
	for _, table := range tables {
		table.c.SetTransform(table.transform)
		result := table.c.ColorAt(table.point)
		if !result.Equals(*table.result) {
			t.Errorf("Point: %v, Expected %v, got %v", table.point, table.result, result)
		}
	}
}
