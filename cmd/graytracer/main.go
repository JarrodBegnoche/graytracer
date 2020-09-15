package main

import (
	"flag"
	"fmt"
	"image"
	"image/png"
	"math"
	"os"
	"runtime"
	"sync"
	"time"
	"github.com/factorion/graytracer/pkg/components"
	"github.com/factorion/graytracer/pkg/primitives"
	"github.com/factorion/graytracer/pkg/shapes"
)

// XY X and Y values for pixel generation
type XY struct {
	X, Y uint
}

var world *components.World
var img *image.RGBA
var imgMutex *sync.Mutex
var camera *components.Camera
var ch chan(XY)
var wg sync.WaitGroup

// RenderPixel Goroutine to render in a multi-threaded environment
func RenderPixel() {
	defer wg.Done()
	open := true
	xyray := XY{}
	for open {
		xyray, open = <- ch
		ray := camera.RayForPixel(xyray.X, xyray.Y)
		col := world.ColorAt(ray)
		imgMutex.Lock()
		img.Set(int(xyray.X), int(xyray.Y), col.ToImageRGBA())
		imgMutex.Unlock()
	}
}

func main() {
	fmt.Println("Starting render")
	var width, height uint
	var threads int
	var fov float64
	flag.IntVar(&threads, "threads", runtime.NumCPU(), "Number of threads for rendering")
	flag.UintVar(&width, "width", 480, "Width of rendered image")
	flag.UintVar(&height, "height", 270, "Height of rendered image")
	flag.Float64Var(&fov, "fov", math.Pi / 3, "Field of View (in Radians)")
	flag.Parse()
	camera = components.MakeCamera(width, height, fov)
	camera.ViewTransform(primitives.MakePoint(0, 1.5, -5),
						 primitives.MakePoint(0, 1, 0),
						 primitives.MakeVector(0, 1, 0))
	ch = make(chan XY, 1000)
	imgMutex = &sync.Mutex{}
	start := time.Now()
	upLeft := image.Point{0, 0}
	lowRight := image.Point{int(width), int(height)}
	img = image.NewRGBA(image.Rectangle{upLeft, lowRight})
	world = &components.World{}
	floorMaterial := primitives.Material{Color:primitives.MakeRGB(1, 0.9, 0.9), Ambient:0.1,
										 Diffuse:0.9, Specular:0, Shininess:200}
	// Floor
	floor := shapes.MakePlane()
	floor.SetMaterial(floorMaterial)
	world.AddObject(floor)
	// Middle
	middle := shapes.MakeSphere()
	middle.SetTransform(primitives.Translation(-0.5, 1, 0.5))
	middle.SetMaterial(primitives.Material{Color:primitives.MakeRGB(0.1, 1, 0.5), Ambient:0.1,
										   Diffuse:0.7, Specular:0.3, Shininess:200})
	world.AddObject(middle)
	// Right
	right := shapes.MakeSphere()
	right.SetTransform(primitives.Translation(1.5, 0.5, -0.5).Multiply(
					   primitives.Scaling(0.5, 0.5, 0.5)))
	right.SetMaterial(primitives.Material{Color:primitives.MakeRGB(0.5, 1, 0.1), Ambient:0.1,
										  Diffuse:0.7, Specular:0.3, Shininess:200})
	world.AddObject(right)
	// Left
	left := shapes.MakeSphere()
	left.SetTransform(primitives.Translation(-1.5, 0.33, -0.75).Multiply(
					  primitives.Scaling(0.33, 0.33, 0.33)))
	left.SetMaterial(primitives.Material{Color:primitives.MakeRGB(1.0, 0.8, 0.1), Ambient:0.1,
		  								 Diffuse:0.7, Specular:0.3, Shininess:200})
	world.AddObject(left)
	light := components.PointLight{Intensity:primitives.MakeRGB(1, 1, 1),
								   Position:primitives.MakePoint(-10, 10, -10)}
	world.AddLight(light)
	fmt.Println("Creating goroutines")
	wg.Add(threads)
	for t := 0; t < threads; t++ {
		go RenderPixel()
	}
	fmt.Println("Starting pixel calculations")
	for y := uint(0); y < height; y++ {
		for x := uint(0); x < width; x++ {
			ch <- XY{X:x, Y:y}
		}
	}
	fmt.Println("Closing channel")
	close(ch)
	wg.Wait()
	fmt.Printf("Render finished : %v\n", time.Since(start))
	f, _ := os.Create("image.png")
	png.Encode(f, img)
}
