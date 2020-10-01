package patterns

import (
	"github.com/factorion/graytracer/pkg/primitives"
	"testing"
)

func TestStripeColorAt(t *testing.T) {
	tables := []struct {
		s *Stripe
		transform primitives.Matrix
		point primitives.PV
		result *RGB
	}{
		{MakeStripe(MakeRGB(1, 1, 1), MakeRGB(0, 0, 0)), primitives.MakeIdentityMatrix(4),
		 primitives.MakePoint(0, 0, 0), MakeRGB(1, 1, 1)},
		
		{MakeStripe(MakeRGB(1, 1, 1), MakeRGB(0, 0, 0)), primitives.MakeIdentityMatrix(4),
		 primitives.MakePoint(0, 1, 0), MakeRGB(1, 1, 1)},
		
		{MakeStripe(MakeRGB(1, 1, 1), MakeRGB(0, 0, 0)), primitives.MakeIdentityMatrix(4),
		 primitives.MakePoint(0, 2, 0), MakeRGB(1, 1, 1)},
		
		{MakeStripe(MakeRGB(1, 1, 1), MakeRGB(0, 0, 0)), primitives.MakeIdentityMatrix(4),
		 primitives.MakePoint(0.9, 0, 0), MakeRGB(1, 1, 1)},
		
		{MakeStripe(MakeRGB(1, 1, 1), MakeRGB(0, 0, 0)), primitives.MakeIdentityMatrix(4),
		 primitives.MakePoint(1, 0, 0), MakeRGB(0, 0, 0)},
		
		{MakeStripe(MakeRGB(1, 1, 1), MakeRGB(0, 0, 0)), primitives.MakeIdentityMatrix(4),
		 primitives.MakePoint(-0.1, 0, 0), MakeRGB(0, 0, 0)},
		
		{MakeStripe(MakeRGB(1, 1, 1), MakeRGB(0, 0, 0)), primitives.MakeIdentityMatrix(4),
		 primitives.MakePoint(-1, 0, 0), MakeRGB(0, 0, 0)},
		
		{MakeStripe(MakeRGB(1, 1, 1), MakeRGB(0, 0, 0)), primitives.MakeIdentityMatrix(4),
		 primitives.MakePoint(-1.1, 0, 0), MakeRGB(1, 1, 1)},
		
		{MakeStripe(MakeRGB(1, 1, 1), MakeRGB(0, 0, 0)), primitives.Scaling(0.5, 0, 0),
		 primitives.MakePoint(0.6, 0, 0), MakeRGB(0, 0, 0)},
	}
	for _, table := range tables {
		table.s.SetTransform(table.transform)
		result := table.s.ColorAt(table.point)
		if !result.Equals(*table.result) {
			t.Errorf("Point: %v, Expected %v, got %v", table.point, table.result, result)
		}
	}
}
