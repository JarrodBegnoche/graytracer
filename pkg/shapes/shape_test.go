package shapes

import (
	"testing"
)

func TestSliceEquals(t *testing.T) {
	tables := []struct {
		a, b []float64
		equals bool
	}{
		{[]float64{0.0, 1.0}, []float64{0.0, 1.0}, true},
		{[]float64{0.0, 1.0}, []float64{1.0, 2.0}, false},
		{[]float64{0.0, 1.0, 2.0}, []float64{0.0, 1.0}, false},
		{[]float64{0.0, 1.0}, []float64{0.0, 1.0000000001}, true},
	}
	for _, table := range tables {
		equals := SliceEquals(table.a, table.b)
		if equals != table.equals {
			t.Errorf("Slice %v and %v returned %v as equals", table.a, table.b, equals)
		}
	}
}
