package primitives

import (
	"testing"
)

func TestRGBAdd(t *testing.T) {
	tables := []struct {
		c1 RGB
		c2 RGB
		sum RGB
	}{
		{RGB{Red:0.9, Green:0.6, Blue:0.75}, RGB{Red:0.7, Green:0.1, Blue:0.25}, RGB{Red:1.6, Green:0.7, Blue:1.0}},
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
		{RGB{Red:0.9, Green:0.6, Blue:0.75}, RGB{Red:0.5, Green:0.1, Blue:0.25}, RGB{Red:0.4, Green:0.5, Blue:0.5}},
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
		{RGB{Red:1.0, Green:0.2, Blue:0.4}, RGB{Red:0.9, Green:1.0, Blue:0.5}, RGB{Red:0.9, Green:0.2, Blue:0.2}},
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
		{RGB{Red:0.2, Green:0.3, Blue:0.4}, 2.0, RGB{Red:0.4, Green:0.6, Blue:0.8}},
	}
	for _, table := range tables {
		scale := table.c1.Scale(table.s)
		if scale != table.scale {
			t.Errorf("Expected %v, got %v", table.scale, scale)
		}
	}
}