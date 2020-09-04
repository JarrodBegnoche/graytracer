package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	_ "math"
	"os"
	"github.com/factorion/graytracer/pkg/primitives"
	"github.com/factorion/graytracer/pkg/shapes"
)

func main() {
	fmt.Println("Starting render")
	width := 100
	height := 100
	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}
	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})
	sphere := shapes.MakeSphere(1)
	transform := primitives.Shearing(1, 0, 0, 0, 0, 0)
	transform = transform.Multiply(primitives.Scaling(0.5, 1, 1))
	sphere.SetTransform(transform)
	wallZ := 10.0
	wallSize := 7.0
	canvasPixels := 100.0
	pixelSize := wallSize / canvasPixels
	half := wallSize / 2
	origin := primitives.MakePoint(0, 0, -5)
	for y := 0; y < height; y++ {
		worldY := half - (pixelSize * float64(y))
		altHeight := height - (y + 1)
		for x := 0; x < width; x++ {
			worldX := -half + (pixelSize * float64(x))
			position := primitives.MakePoint(worldX, worldY, wallZ)
			r := primitives.Ray{Origin:origin, Direction:position.Subtract(origin).Normalize()}
			intersections := shapes.Intersections{}
			hits := sphere.Intersect(r)
			for i := range hits {
				intersections[hits[i]] = shapes.Intersection{Distance:hits[i], Obj:sphere}
			}
			if len(intersections) != 0 {
				img.Set(x, altHeight, color.White)
			} else {
				img.Set(x, altHeight, color.Black)
			}
		}
	}
	f, _ := os.Create("image.png")
	png.Encode(f, img)
}
