package patterns_test

import(
	"testing"
	"github.com/factorion/graytracer/pkg/primitives"
	"github.com/factorion/graytracer/pkg/patterns"
)

func TestPatternTransformation(t *testing.T) {
	tables := []struct {
		patternTransform primitives.Matrix
		point primitives.PV
		result *patterns.RGB
	}{
		{primitives.Scaling(2, 2, 2),
		 primitives.MakePoint(2, 3, 4),
		 patterns.MakeRGB(1, 1.5, 2)},
		
		{primitives.Translation(0.5, 1, 1.5),
		 primitives.MakePoint(2.5, 3, 3.5),
		 patterns.MakeRGB(2, 2, 2)},
	}
	for _, table := range tables {
		material := patterns.MakeDefaultMaterial()
		material.Pat = &patterns.TestPattern{PatternBase:patterns.MakePatternBase()}
		material.Pat.SetTransform(table.patternTransform)
		result := material.Pat.ColorAt(table.point)
		if !result.Equals(*table.result) {
			t.Errorf("Expected %v, got %v", *table.result, result)
		}
	}
}
