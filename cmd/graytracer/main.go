package main

import (
	"fmt"
	"github.com/factorion/graytracer/pkg/primitives"
)

func main() {
	vector := primitives.MakeVector(1.0, 0.0, 0.0)
	fmt.Printf("Vector = %v\n", vector)
	point := primitives.MakePoint(2.0, 3.0, 4.0)
	fmt.Printf("Point = %v\n", point)
}