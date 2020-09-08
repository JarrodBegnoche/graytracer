package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"time"
	"github.com/factorion/graytracer/pkg/primitives"
	"github.com/factorion/graytracer/pkg/shapes"
)

func main() {
	fmt.Println("Starting render")
	start := time.Now()
	width := 1000
	height := 1000
	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}
	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})
	sphere := shapes.MakeSphere()
	sphere.SetMaterial(primitives.Material{Color:primitives.MakeRGB(1, 0.2, 1),
										   Ambient:0.1, Diffuse:0.9, Specular:0.9, Shininess:200})
	/*transform := primitives.Shearing(1, 0, 0, 0, 0, 0)
	transform = transform.Multiply(primitives.Scaling(0.5, 1, 1))
	sphere.SetTransform(transform)*/
	light := primitives.PointLight{Intensity:primitives.MakeRGB(1, 1, 1),
								   Position:primitives.MakePoint(-10, 10, -10)}
	wallZ := 10.0
	wallSize := 7.0
	canvasPixels := 1000.0
	pixelSize := wallSize / canvasPixels
	half := wallSize / 2
	origin := primitives.MakePoint(0, 0, -5)
	for y := 0; y < height; y++ {
		worldY := half - (pixelSize * float64(y))
		for x := 0; x < width; x++ {
			worldX := -half + (pixelSize * float64(x))
			position := primitives.MakePoint(worldX, worldY, wallZ)
			r := primitives.Ray{Origin:origin, Direction:position.Subtract(origin).Normalize()}
			intersections := []shapes.Intersection{}
			hits := sphere.Intersect(r)
			for i := range hits {
				intersections = append(intersections, shapes.Intersection{Distance:hits[i], Obj:sphere})
			}
			if intersection, hit := shapes.Hit(intersections); hit {
				point := r.Position(intersection.Distance)
				normal := intersection.Obj.Normal(point)
				eye := r.Direction.Negate()
				col := primitives.Lighting(intersection.Obj.Material(), light, point, eye, normal)
				img.Set(x, y, col.ToImageRGBA())
			} else {
				img.Set(x, y, color.Black)
			}
		}
	}
	fmt.Printf("Render finished : %v\n", time.Since(start))
	f, _ := os.Create("image.png")
	png.Encode(f, img)
}
