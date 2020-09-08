package primitives

import (
	"testing"
)

func TestMakeRGB(t * testing.T) {
	tables := []struct {
		rgb1, rgb2 RGB
	}{
		{MakeRGB(1, 2, 3), RGB{1, 2, 3}},
		{MakeRGB(3, 2, 1), RGB{3, 2, 1}},
	}
	for _, table := range tables {
		if table.rgb1 != table.rgb2 {
			t.Errorf("Expected %v, got %v", table.rgb1, table.rgb2)
		}
	}
}

func TestLightEquals(t *testing.T) {
	tables := []struct {
		rgb1, rgb2 RGB
		equals bool
	}{
		{MakeRGB(1, 2, 3), MakeRGB(1, 2, 3), true},
		{MakeRGB(1, 1, 1), MakeRGB(0, 1, 1), false},
		{MakeRGB(1, 1, 1), MakeRGB(1, 0, 1), false},
		{MakeRGB(1, 1, 1), MakeRGB(1, 1, 0), false},
		{MakeRGB(1, 2, 3.123456789), MakeRGB(1, 2, 3.123456788), true},
	}
	for _, table := range tables {
		equals := table.rgb1.Equals(table.rgb2)
		if equals != table.equals {
			t.Errorf("PV %v and %v returned %v for Equals", table.rgb1, table.rgb2, equals)
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
