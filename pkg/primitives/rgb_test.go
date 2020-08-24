package primitives

import (
	"testing"
)

func TestRGBAccessors(t *testing.T) {
	tables := []struct {
		rgb RGB
		red float64
		green float64
		blue float64
	}{
		{RGB{red:0.9, green:0.6, blue:0.75}, 0.9, 0.6, 0.75},
	}
	for _, table := range tables {
		if table.rgb.Red() != table.red {
			t.Errorf("Expected Red value %v, got %v", table.red, table.rgb.Red())
		}
		if table.rgb.Green() != table.green {
			t.Errorf("Expected Green value %v, got %v", table.green, table.rgb.Green())
		}
		if table.rgb.Blue() != table.blue {
			t.Errorf("Expected Blue value %v, got %v", table.blue, table.rgb.Blue())
		}
	}
}

func TestRGBAdd(t *testing.T) {
	tables := []struct {
		c1 RGB
		c2 RGB
		sum RGB
	}{
		{RGB{red:0.9, green:0.6, blue:0.75}, RGB{red:0.7, green:0.1, blue:0.25}, RGB{red:1.6, green:0.7, blue:1.0}},
	}
	for _, table := range tables {
		sum := table.c1.Add(table.c2)
		if sum != table.sum {
			t.Errorf("Expected %v, got %v", table.sum, sum)
		}
	}
}

func TestRGBSubtract(t *testing.T) {
	tables := []struct {
		c1 RGB
		c2 RGB
		diff RGB
	}{
		{RGB{red:0.9, green:0.6, blue:0.75}, RGB{red:0.5, green:0.1, blue:0.25}, RGB{red:0.4, green:0.5, blue:0.5}},
	}
	for _, table := range tables {
		diff := table.c1.Subtract(table.c2)
		if diff != table.diff {
			t.Errorf("Expected %v, got %v", table.diff, diff)
		}
	}
}

func TestRGBMultiply(t *testing.T) {
	tables := []struct {
		c1 RGB
		c2 RGB
		prod RGB
	}{
		{RGB{red:1.0, green:0.2, blue:0.4}, RGB{red:0.9, green:1.0, blue:0.5}, RGB{red:0.9, green:0.2, blue:0.2}},
	}
	for _, table := range tables {
		prod := table.c1.Multiply(table.c2)
		if prod != table.prod {
			t.Errorf("Expected %v, got %v", table.prod, prod)
		}
	}
}

func TestRGBScale(t * testing.T) {
	tables := []struct {
		c1 RGB
		s float64
		scale RGB
	}{
		{RGB{red:0.2, green:0.3, blue:0.4}, 2.0, RGB{red:0.4, green:0.6, blue:0.8}},
	}
	for _, table := range tables {
		scale := table.c1.Scale(table.s)
		if scale != table.scale {
			t.Errorf("Expected %v, got %v", table.scale, scale)
		}
	}
}
