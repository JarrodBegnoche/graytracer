package patterns_test

import (
	"testing"
	"github.com/factorion/graytracer/pkg/patterns"
	"github.com/factorion/graytracer/pkg/primitives"
)

func TestStripeColorAt(t *testing.T) {
	tables := []struct {
		s *patterns.Stripe
		transform primitives.Matrix
		point primitives.PV
		result *patterns.RGB
	}{
		{patterns.MakeStripe(patterns.MakeRGB(1, 1, 1), patterns.MakeRGB(0, 0, 0)),
		 primitives.MakeIdentityMatrix(4),
		 primitives.MakePoint(0, 0, 0),
		 patterns.MakeRGB(1, 1, 1)},
		
		{patterns.MakeStripe(patterns.MakeRGB(1, 1, 1), patterns.MakeRGB(0, 0, 0)),
		 primitives.MakeIdentityMatrix(4),
		 primitives.MakePoint(0, 1, 0),
		 patterns.MakeRGB(1, 1, 1)},
		
		{patterns.MakeStripe(patterns.MakeRGB(1, 1, 1), patterns.MakeRGB(0, 0, 0)),
		 primitives.MakeIdentityMatrix(4),
		 primitives.MakePoint(0, 2, 0),
		 patterns.MakeRGB(1, 1, 1)},
		
		{patterns.MakeStripe(patterns.MakeRGB(1, 1, 1), patterns.MakeRGB(0, 0, 0)),
		 primitives.MakeIdentityMatrix(4),
		 primitives.MakePoint(0.9, 0, 0),
		 patterns.MakeRGB(1, 1, 1)},
		
		{patterns.MakeStripe(patterns.MakeRGB(1, 1, 1), patterns.MakeRGB(0, 0, 0)),
		 primitives.MakeIdentityMatrix(4),
		 primitives.MakePoint(1, 0, 0),
		 patterns.MakeRGB(0, 0, 0)},
		
		{patterns.MakeStripe(patterns.MakeRGB(1, 1, 1), patterns.MakeRGB(0, 0, 0)),
		 primitives.MakeIdentityMatrix(4),
		 primitives.MakePoint(-0.1, 0, 0),
		 patterns.MakeRGB(0, 0, 0)},
		
		{patterns.MakeStripe(patterns.MakeRGB(1, 1, 1), patterns.MakeRGB(0, 0, 0)),
		 primitives.MakeIdentityMatrix(4),
		 primitives.MakePoint(-1, 0, 0),
		 patterns.MakeRGB(0, 0, 0)},
		
		{patterns.MakeStripe(patterns.MakeRGB(1, 1, 1), patterns.MakeRGB(0, 0, 0)),
		 primitives.MakeIdentityMatrix(4),
		 primitives.MakePoint(-1.1, 0, 0),
		 patterns.MakeRGB(1, 1, 1)},
		
		{patterns.MakeStripe(patterns.MakeRGB(1, 1, 1), patterns.MakeRGB(0, 0, 0)),
		 primitives.Scaling(0.5, 0, 0),
		 primitives.MakePoint(0.6, 0, 0),
		 patterns.MakeRGB(0, 0, 0)},
	}
	for _, table := range tables {
		table.s.SetTransform(table.transform)
		result := table.s.ColorAt(table.point)
		if !result.Equals(*table.result) {
			t.Errorf("Point: %v, Expected %v, got %v", table.point, table.result, result)
		}
	}
}
