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
	m := primitives.MakeMatrix(4)
	fmt.Printf("Matrix = %v\n", m)
	fmt.Println(len(m))
	m2 := primitives.MakeMatrix(3)
	fmt.Printf("Matrix3x3 = %v\n", m2)
	fmt.Println(len(m2))
	im := primitives.MakeIdentityMatrix(4)
	fmt.Printf("Identity Matrix 2x2 = %v\n", im)
}
