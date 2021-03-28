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
	"github.com/factorion/graytracer/pkg/patterns"
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
var ch chan (XY)
var wg sync.WaitGroup

// RenderPixel Goroutine to render in a multi-threaded environment
func RenderPixel() {
	defer wg.Done()
	open := true
	xyray := XY{}
	for open {
		xyray, open = <-ch
		ray := camera.RayForPixel(xyray.X, xyray.Y)
		col := world.ColorAt(ray, 5)
		imgMutex.Lock()
		img.Set(int(xyray.X), int(xyray.Y), col.ToImageRGBA())
		imgMutex.Unlock()
	}
}

// MakeHex Make a Hex group object
func MakeHex(mat patterns.Material, transform primitives.Matrix) shapes.Shape {
	// Hex group
	hex := shapes.MakeGroup()
	hex.SetTransform(transform)
	for i := 0.0; i < 6; i++ {
		corner := shapes.MakeSphere()
		corner.SetMaterial(mat)
		corner.SetTransform(primitives.RotationY(i * math.Pi / 3).Multiply(
							primitives.Translation(0, 0, -1).Multiply(
							primitives.Scaling(0.25, 0.25, 0.25))))
		edge := shapes.MakeCylinder(true)
		edge.SetMaterial(mat)
		edge.SetTransform(primitives.RotationY(i * math.Pi / 3).Multiply(
						  primitives.Translation(0, 0, -1).Multiply(
						  primitives.RotationY(-math.Pi / 6).Multiply(
						  primitives.RotationZ(-math.Pi / 2).Multiply(
						  primitives.Scaling(0.25, 1, 0.25))))))
		top := shapes.MakeCone(true)
		top.SetMaterial(mat)
		top.SetTransform(primitives.RotationY(i * math.Pi / 3).Multiply(
						 primitives.Translation(0, 1, 0).Multiply(
						 primitives.RotationX(math.Pi / 4).Multiply(
						 primitives.Scaling(0.25, math.Sqrt(2), 0.25)))))
		bottom := shapes.MakeCone(true)
		bottom.SetMaterial(mat)
		bottom.SetTransform(primitives.RotationY(i * math.Pi / 3).Multiply(
						    primitives.Translation(0, -1, 0).Multiply(
							primitives.RotationX(3 * math.Pi / 4).Multiply(
							primitives.Scaling(0.25, math.Sqrt(2), 0.25)))))
		hex.AddShape(corner)
		hex.AddShape(edge)
		hex.AddShape(top)
		hex.AddShape(bottom)
	}
	return hex
}

func main() {
	fmt.Println("Starting render")
	var width, height uint
	var threads int
	var fov float64
	flag.IntVar(&threads, "threads", runtime.NumCPU(), "Number of threads for rendering")
	flag.UintVar(&width, "width", 320, "Width of rendered image")
	flag.UintVar(&height, "height", 240, "Height of rendered image")
	flag.Float64Var(&fov, "fov", math.Pi/4, "Field of View (in Radians)")
	flag.Parse()
	camera = components.MakeCamera(width, height, fov)
	camera.ViewTransform(primitives.MakePoint(-20, 50, -35),
						 primitives.MakePoint(15, 10, 15),
						 primitives.MakeVector(0, 1, 0))
	ch = make(chan XY, 1000)
	imgMutex = &sync.Mutex{}
	start := time.Now()
	upLeft := image.Point{0, 0}
	lowRight := image.Point{int(width), int(height)}
	img = image.NewRGBA(image.Rectangle{upLeft, lowRight})
	world = components.MakeWorld()
	world.SetBackground(*patterns.MakeRGB(0.9, 0.9, 0.9))
	for x := 0.0; x < 10; x++ {
		for y := 0.0; y < 10; y++ {
			for z := 0.0; z < 10; z++ {
				mat := patterns.Material{Pat: patterns.MakeRGB(x * 0.1, y * 0.1, z * 0.1),
										 Ambient: 0.1, Diffuse: 0.7, Specular: 0.5,
										 Shininess: 200, Reflective: 0.1, Transparency: 0, RefractiveIndex: 0}
				hex := MakeHex(mat, primitives.Translation(x * 4, y * 3, z * 4))
				world.AddObject(hex)
			}
		}
	}
	
	// Render time without bounding boxes
	// Render finished : 1m34.8119446s
	// Render time with bounding boxes
	// Render finished : 2.4859576s

	light := components.PointLight{Intensity: patterns.MakeRGB(1, 1, 1),
		Position: primitives.MakePoint(-30, 60, -30)}
	world.AddLight(light)
	fmt.Println("Creating goroutines")
	wg.Add(threads)
	for t := 0; t < threads; t++ {
		go RenderPixel()
	}
	fmt.Println("Starting pixel calculations")
	for y := uint(0); y < height; y++ {
		for x := uint(0); x < width; x++ {
			ch <- XY{X: x, Y: y}
		}
	}
	fmt.Println("Closing channel")
	close(ch)
	wg.Wait()
	fmt.Printf("Render finished : %v\n", time.Since(start))
	f, _ := os.Create("image.png")
	png.Encode(f, img)
}
