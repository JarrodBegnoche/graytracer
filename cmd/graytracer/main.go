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
		col := world.ColorAt(ray, 5)
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
	floorMaterial := patterns.Material{Pat:patterns.MakeChecker(patterns.MakeRGB(0.05, 0.05, 0.05),
																patterns.MakeRGB(0.95, 0.95, 0.95)),
									   Ambient:0.1, Diffuse:0.9, Specular:0, Shininess:200,
									   Reflective:0, Transparency:0, RefractiveIndex:1}
	//floorMaterial.Pat.SetTransform(primitives.Scaling(0.25, 0, 0))
	//floorMaterial.Pat.SetTransform(primitives.RotationZ(math.Pi / 2))
	// Floor
	floor := shapes.MakePlane()
	floor.SetMaterial(floorMaterial)
	world.AddObject(floor)
	ceiling := shapes.MakePlane()
	ceiling.SetMaterial(floorMaterial)
	ceiling.SetTransform(primitives.Translation(0, 10, 0).Multiply(primitives.RotationX(math.Pi)))
	world.AddObject(ceiling)
	frontWall := shapes.MakePlane()
	frontWall.SetMaterial(floorMaterial)
	frontWall.SetTransform(primitives.Translation(0, 0, 5).Multiply(primitives.RotationX(math.Pi / 2)))
	world.AddObject(frontWall)
	/*backWall := shapes.MakePlane()
	backWall.SetMaterial(floorMaterial)
	backWall.SetTransform(primitives.Translation(0, 0, -5).Multiply(primitives.RotationX(-math.Pi / 2)))
	world.AddObject(backWall)*/
	leftWall := shapes.MakePlane()
	leftWall.SetMaterial(floorMaterial)
	leftWall.SetTransform(primitives.Translation(-5, 0, 0).Multiply(primitives.RotationZ(-math.Pi / 2)))
	world.AddObject(leftWall)
	rightWall := shapes.MakePlane()
	rightWall.SetMaterial(floorMaterial)
	rightWall.SetTransform(primitives.Translation(5, 0, 0).Multiply(primitives.RotationZ(math.Pi / 2)))
	world.AddObject(rightWall)
	// Middle
	middle := shapes.MakeSphere()
	checker := patterns.MakeChecker(patterns.MakeRGB(0.5, 1, 0.1), patterns.MakeRGB(0.9, 0.9, 0.1))
	checker.SetTransform(primitives.Scaling(0.125, 0.125, 0.125))
	middle.SetTransform(primitives.Translation(-0.5, 1, 0.5))
	middle.SetMaterial(patterns.Material{Pat:checker, Ambient:0.1, Diffuse:0.7, Specular:0.3,
										 Shininess:20, Reflective:0, Transparency:0, RefractiveIndex:1})
	world.AddObject(middle)
	// Right
	right := shapes.MakeSphere()
	right.SetTransform(primitives.Translation(1.5, 0.5, -0.5).Multiply(
					   primitives.Scaling(0.5, 0.5, 0.5)))
	right.SetMaterial(patterns.Material{Pat:patterns.MakeRGB(0.01, 0.01, 0.01), Ambient:0.1, Diffuse:0.7, Specular:0.7,
										Shininess:200, Reflective:0.1, Transparency:0.9, RefractiveIndex:1.333333})
	world.AddObject(right)
	// Left
	left := shapes.MakeCube()
	left.SetTransform(primitives.Translation(-1.5, 0.33, -0.75).Multiply(
					  primitives.RotationY(math.Pi / 6).Multiply(
					  primitives.Scaling(0.33, 0.33, 0.33))))
	left.SetMaterial(patterns.Material{Pat:patterns.MakeRGB(0.05, 0.05, 0.05), Ambient:0.1,
									   Diffuse:0.7, Specular:0.3, Shininess:200, Reflective:1,
									   Transparency:0, RefractiveIndex:1})
	world.AddObject(left)
	light := components.PointLight{Intensity:patterns.MakeRGB(1, 1, 1),
								   Position:primitives.MakePoint(-4.5, 4.5, -4.5)}
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
