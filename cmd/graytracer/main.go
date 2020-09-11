package main

import (
	"flag"
	"fmt"
	"image"
	"image/png"
	"os"
	"runtime"
	"sync"
	"time"
	"github.com/factorion/graytracer/pkg/components"
	"github.com/factorion/graytracer/pkg/primitives"
	"github.com/factorion/graytracer/pkg/shapes"
)

// XYRay Values needed for 
type XYRay struct {
	X, Y int
	Ray primitives.Ray
}

// RenderPixel Goroutine to render in a multi-threaded environment
func RenderPixel(world *components.World, img *image.RGBA, ch chan XYRay, imgMutex *sync.Mutex) {
	open := true
	xyray := XYRay{}
	for open {
		xyray, open = <- ch
		col := world.ColorAt(xyray.Ray)
		imgMutex.Lock()
		img.Set(xyray.X, xyray.Y, col.ToImageRGBA())
		imgMutex.Unlock()
	}
}

func main() {
	fmt.Println("Starting render")
	var threads, width, height int
	flag.IntVar(&threads, "threads", runtime.NumCPU(), "Number of threads for rendering")
	flag.IntVar(&height, "height", 1000, "Height of rendered image")
	flag.IntVar(&width, "width", 1000, "Width of rendered image")
	ch := make(chan XYRay, 1000)
	imgMutex := &sync.Mutex{}
	start := time.Now()
	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}
	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})
	world := &components.World{}
	sphere := shapes.MakeSphere()
	sphere.SetMaterial(primitives.Material{Color:primitives.MakeRGB(1, 0.2, 1),
										   Ambient:0.1, Diffuse:0.9, Specular:0.9, Shininess:200})
	world.AddObject(sphere)
	/*transform := primitives.Shearing(1, 0, 0, 0, 0, 0)
	transform = transform.Multiply(primitives.Scaling(0.5, 1, 1))
	sphere.SetTransform(transform)*/
	light := components.PointLight{Intensity:primitives.MakeRGB(1, 1, 1),
								   Position:primitives.MakePoint(-10, 10, -10)}
	world.AddLight(light)
	wallZ := 10.0
	wallSize := 7.0
	canvasPixels := 1000.0
	pixelSize := wallSize / canvasPixels
	half := wallSize / 2
	origin := primitives.MakePoint(0, 0, -5)
	fmt.Println("Creating goroutines")
	for t := 0; t < threads; t++ {
		go RenderPixel(world, img, ch, imgMutex)
	}
	fmt.Println("Starting pixel calculations")
	for y := 0; y < height; y++ {
		worldY := half - (pixelSize * float64(y))
		for x := 0; x < width; x++ {
			worldX := -half + (pixelSize * float64(x))
			position := primitives.MakePoint(worldX, worldY, wallZ)
			ray := primitives.Ray{Origin:origin, Direction:position.Subtract(origin).Normalize()}
			ch <- XYRay{X:x, Y:y, Ray:ray}
		}
	}
	fmt.Println("Closing channel")
	close(ch)
	fmt.Printf("Render finished : %v\n", time.Since(start))
	f, _ := os.Create("image.png")
	png.Encode(f, img)
}
