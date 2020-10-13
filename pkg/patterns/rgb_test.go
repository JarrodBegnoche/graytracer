package patterns_test

import(
	"testing"
	"github.com/factorion/graytracer/pkg/patterns"
)

func TestLightEquals(t *testing.T) {
	tables := []struct {
		rgb1, rgb2 *patterns.RGB
		equals bool
	}{
		{patterns.MakeRGB(1, 2, 3), patterns.MakeRGB(1, 2, 3), true},
		{patterns.MakeRGB(1, 1, 1), patterns.MakeRGB(0, 1, 1), false},
		{patterns.MakeRGB(1, 1, 1), patterns.MakeRGB(1, 0, 1), false},
		{patterns.MakeRGB(1, 1, 1), patterns.MakeRGB(1, 1, 0), false},
		{patterns.MakeRGB(1, 2, 3.123456789), patterns.MakeRGB(1, 2, 3.123456788), true},
	}
	for _, table := range tables {
		equals := table.rgb1.Equals(*table.rgb2)
		if equals != table.equals {
			t.Errorf("PV %v and %v returned %v for Equals", table.rgb1, table.rgb2, equals)
		}
	}
}

func TestRGBAdd(t *testing.T) {
	tables := []struct {
		c1, c2, sum *patterns.RGB
	}{
		{patterns.MakeRGB(0.9, 0.6, 0.75), patterns.MakeRGB(0.7, 0.1, 0.25), patterns.MakeRGB(1.6, 0.7, 1.0)},
	}
	for _, table := range tables {
		sum := table.c1.Add(*table.c2)
		if !sum.Equals(*table.sum) {
			t.Errorf("Expected %v, got %v", *table.sum, sum)
		}
	}
}

func TestRGBSubtract(t *testing.T) {
	tables := []struct {
		c1, c2, diff *patterns.RGB
	}{
		{patterns.MakeRGB(0.9, 0.6, 0.75), patterns.MakeRGB(0.5, 0.1, 0.25), patterns.MakeRGB(0.4, 0.5, 0.5)},
	}
	for _, table := range tables {
		diff := table.c1.Subtract(*table.c2)
		if !diff.Equals(*table.diff) {
			t.Errorf("Expected %v, got %v", *table.diff, diff)
		}
	}
}

func TestRGBMultiply(t *testing.T) {
	tables := []struct {
		c1, c2, prod *patterns.RGB
	}{
		{patterns.MakeRGB(1.0, 0.2, 0.4), patterns.MakeRGB(0.9, 1.0, 0.5), patterns.MakeRGB(0.9, 0.2, 0.2)},
	}
	for _, table := range tables {
		prod := table.c1.Multiply(*table.c2)
		if !prod.Equals(*table.prod) {
			t.Errorf("Expected %v, got %v", *table.prod, prod)
		}
	}
}

func TestRGBScale(t * testing.T) {
	tables := []struct {
		c1 *patterns.RGB
		s float64
		scale *patterns.RGB
	}{
		{patterns.MakeRGB(0.2, 0.3, 0.4), 2.0, patterns.MakeRGB(0.4, 0.6, 0.8)},
	}
	for _, table := range tables {
		scale := table.c1.Scale(table.s)
		if !scale.Equals(*table.scale) {
			t.Errorf("Expected %v, got %v", *table.scale, scale)
		}
	}
}
