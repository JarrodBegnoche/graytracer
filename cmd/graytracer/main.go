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

	"github.com/schollz/progressbar/v3"
)

// XY X and Y values for pixel generation
type XY struct {
	X, Y uint64
}

var world *components.World
var img *image.RGBA
var imgMutex *sync.Mutex
var camera *components.Camera
var ch chan (XY)
var wg sync.WaitGroup
var bar progressbar.ProgressBar
var prog_count = uint64(0)

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
		prog_count++
		if prog_count >= 100 {
			bar.Add(int(prog_count))
			prog_count = 0
		}
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
	var width, height uint64
	var threads int
	var fov float64
	flag.IntVar(&threads, "threads", runtime.NumCPU(), "Number of threads for rendering")
	flag.Uint64Var(&width, "width", 320, "Width of rendered image")
	flag.Uint64Var(&height, "height", 180, "Height of rendered image")
	flag.Float64Var(&fov, "fov", math.Pi/3, "Field of View (in Radians)")
	flag.Parse()
	camera = components.MakeCamera(width, height, fov)
	camera.ViewTransform(primitives.MakePoint(-6, 6, -10),
		primitives.MakePoint(6, 0, 6),
		primitives.MakeVector(-0.45, 1, 0))
	ch = make(chan XY, 1000)
	imgMutex = &sync.Mutex{}
	start := time.Now()
	img = image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{int(width), int(height)}})
	world = components.MakeWorld()
	world.SetBackground(*patterns.MakeRGB(0, 0, 0))
	light1 := components.PointLight{Intensity: patterns.MakeRGB(1, 1, 1),
		Position: primitives.MakePoint(50, 100, -50)}
	light2 := components.PointLight{Intensity: patterns.MakeRGB(0.2, 0.2, 0.2),
		Position: primitives.MakePoint(-400, 50, -10)}
	world.AddLight(light1)
	world.AddLight(light2)
	plane := shapes.MakePlane()
	plane.SetMaterial(patterns.Material{Pat: patterns.MakeRGB(1, 1, 1), Ambient: 1, Diffuse: 0, Specular: 0,
		Shininess: 200, Reflective: 0, Transparency: 0, RefractiveIndex: 1.0})
	plane.SetTransform(primitives.Translation(0, 0, 500).Multiply(primitives.RotationX(math.Pi / 2)))
	world.AddObject(plane)
	sphere := shapes.MakeSphere()
	sphere.SetMaterial(patterns.Material{Pat: patterns.MakeRGB(0.373, 0.404, 0.550),
		Ambient: 0, Diffuse: 0.2, Specular: 1, Shininess: 200, Reflective: 0.7, Transparency: 0.7, RefractiveIndex: 1.5})
	sphere.SetTransform(primitives.Scaling(3.5, 3.5, 3.5).Multiply(primitives.Scaling(0.5, 0.5, 0.5).Multiply(primitives.Translation(1, -1, 1))))
	world.AddObject(sphere)
	white_mat := patterns.Material{Pat: patterns.MakeRGB(1, 1, 1), Ambient: 0.1, Diffuse: 0.7, Specular: 0,
		Shininess: 20, Reflective: 0.1, Transparency: 0, RefractiveIndex: 1.0}
	blue_mat := patterns.Material{Pat: patterns.MakeRGB(0.537, 0.831, 0.914), Ambient: 0.1, Diffuse: 0.7, Specular: 0,
		Shininess: 20, Reflective: 0.1, Transparency: 0, RefractiveIndex: 1.0}
	red_mat := patterns.Material{Pat: patterns.MakeRGB(0.941, 0.322, 0.388), Ambient: 0.1, Diffuse: 0.7, Specular: 0,
		Shininess: 20, Reflective: 0.1, Transparency: 0, RefractiveIndex: 1.0}
	purple_mat := patterns.Material{Pat: patterns.MakeRGB(0.373, 0.404, 0.550), Ambient: 0.1, Diffuse: 0.7, Specular: 0,
		Shininess: 20, Reflective: 0.1, Transparency: 0, RefractiveIndex: 1.0}
	cube01 := shapes.MakeCube()
	cube01.SetMaterial(white_mat)
	cube01.SetTransform(primitives.Translation(4, 0, 0).Multiply(primitives.Scaling(3, 3, 3).Multiply(primitives.Scaling(0.5, 0.5, 0.5).Multiply(primitives.Translation(1, -1, 1)))))
	world.AddObject(cube01)
	cube02 := shapes.MakeCube()
	cube02.SetMaterial(blue_mat)
	cube02.SetTransform(primitives.Translation(8.5, 1.5, -0.5).Multiply(primitives.Scaling(3.5, 3.5, 3.5).Multiply(primitives.Scaling(0.5, 0.5, 0.5).Multiply(primitives.Translation(1, -1, 1)))))
	world.AddObject(cube02)
	cube03 := shapes.MakeCube()
	cube03.SetMaterial(red_mat)
	cube03.SetTransform(primitives.Translation(0, 0, 4).Multiply(primitives.Scaling(3.5, 3.5, 3.5).Multiply(primitives.Scaling(0.5, 0.5, 0.5).Multiply(primitives.Translation(1, -1, 1)))))
	world.AddObject(cube03)
	cube04 := shapes.MakeCube()
	cube04.SetMaterial(white_mat)
	cube04.SetTransform(primitives.Translation(4, 0, 4).Multiply(primitives.Scaling(2, 2, 2).Multiply(primitives.Scaling(0.5, 0.5, 0.5).Multiply(primitives.Translation(1, -1, 1)))))
	world.AddObject(cube04)
	cube05 := shapes.MakeCube()
	cube05.SetMaterial(purple_mat)
	cube05.SetTransform(primitives.Translation(7.5, 0.5, 4).Multiply(primitives.Scaling(3, 3, 3).Multiply(primitives.Scaling(0.5, 0.5, 0.5).Multiply(primitives.Translation(1, -1, 1)))))
	world.AddObject(cube05)
	cube06 := shapes.MakeCube()
	cube06.SetMaterial(white_mat)
	cube06.SetTransform(primitives.Translation(-0.25, 0.25, 8).Multiply(primitives.Scaling(3, 3, 3).Multiply(primitives.Scaling(0.5, 0.5, 0.5).Multiply(primitives.Translation(1, -1, 1)))))
	world.AddObject(cube06)
	cube07 := shapes.MakeCube()
	cube07.SetMaterial(blue_mat)
	cube07.SetTransform(primitives.Translation(4, 1, 7.5).Multiply(primitives.Scaling(3.5, 3.5, 3.5).Multiply(primitives.Scaling(0.5, 0.5, 0.5).Multiply(primitives.Translation(1, -1, 1)))))
	world.AddObject(cube07)
	cube08 := shapes.MakeCube()
	cube08.SetMaterial(red_mat)
	cube08.SetTransform(primitives.Translation(10, 2, 7.5).Multiply(primitives.Scaling(3, 3, 3).Multiply(primitives.Scaling(0.5, 0.5, 0.5).Multiply(primitives.Translation(1, -1, 1)))))
	world.AddObject(cube08)
	cube09 := shapes.MakeCube()
	cube09.SetMaterial(white_mat)
	cube09.SetTransform(primitives.Translation(8, 2, 12).Multiply(primitives.Scaling(2, 2, 2).Multiply(primitives.Scaling(0.5, 0.5, 0.5).Multiply(primitives.Translation(1, -1, 1)))))
	world.AddObject(cube09)
	cube10 := shapes.MakeCube()
	cube10.SetMaterial(white_mat)
	cube10.SetTransform(primitives.Translation(20, 1, 9).Multiply(primitives.Scaling(2, 2, 2).Multiply(primitives.Scaling(0.5, 0.5, 0.5).Multiply(primitives.Translation(1, -1, 1)))))
	world.AddObject(cube10)
	cube11 := shapes.MakeCube()
	cube11.SetMaterial(blue_mat)
	cube11.SetTransform(primitives.Translation(-0.5, -5, 0.25).Multiply(primitives.Scaling(3.5, 3.5, 3.5).Multiply(primitives.Scaling(0.5, 0.5, 0.5).Multiply(primitives.Translation(1, -1, 1)))))
	world.AddObject(cube11)
	cube12 := shapes.MakeCube()
	cube12.SetMaterial(red_mat)
	cube12.SetTransform(primitives.Translation(4, -4, 0).Multiply(primitives.Scaling(3.5, 3.5, 3.5).Multiply(primitives.Scaling(0.5, 0.5, 0.5).Multiply(primitives.Translation(1, -1, 1)))))
	world.AddObject(cube12)
	cube13 := shapes.MakeCube()
	cube13.SetMaterial(white_mat)
	cube13.SetTransform(primitives.Translation(8.5, -4, 0).Multiply(primitives.Scaling(3.5, 3.5, 3.5).Multiply(primitives.Scaling(0.5, 0.5, 0.5).Multiply(primitives.Translation(1, -1, 1)))))
	world.AddObject(cube13)
	cube14 := shapes.MakeCube()
	cube14.SetMaterial(white_mat)
	cube14.SetTransform(primitives.Translation(0, -4, 4).Multiply(primitives.Scaling(3.5, 3.5, 3.5).Multiply(primitives.Scaling(0.5, 0.5, 0.5).Multiply(primitives.Translation(1, -1, 1)))))
	world.AddObject(cube14)
	cube15 := shapes.MakeCube()
	cube15.SetMaterial(purple_mat)
	cube15.SetTransform(primitives.Translation(-0.5, -4.5, 8).Multiply(primitives.Scaling(3.5, 3.5, 3.5).Multiply(primitives.Scaling(0.5, 0.5, 0.5).Multiply(primitives.Translation(1, -1, 1)))))
	world.AddObject(cube15)
	cube16 := shapes.MakeCube()
	cube16.SetMaterial(white_mat)
	cube16.SetTransform(primitives.Translation(0, -8, 4).Multiply(primitives.Scaling(3.5, 3.5, 3.5).Multiply(primitives.Scaling(0.5, 0.5, 0.5).Multiply(primitives.Translation(1, -1, 1)))))
	world.AddObject(cube16)
	cube17 := shapes.MakeCube()
	cube17.SetMaterial(white_mat)
	cube17.SetTransform(primitives.Translation(-0.5, -8.5, 8).Multiply(primitives.Scaling(3.5, 3.5, 3.5).Multiply(primitives.Scaling(0.5, 0.5, 0.5).Multiply(primitives.Translation(1, -1, 1)))))
	world.AddObject(cube17)
	fmt.Println("Creating goroutines")
	wg.Add(threads)
	for t := 0; t < threads; t++ {
		go RenderPixel()
	}
	fmt.Println("Starting pixel calculations")
	bar = *progressbar.Default(int64(width * height))
	for y := uint64(0); y < height; y++ {
		for x := uint64(0); x < width; x++ {
			ch <- XY{X: x, Y: y}
		}
	}
	close(ch)
	wg.Wait()
	bar.Add(int(prog_count))
	fmt.Printf("Render finished : %v\n", time.Since(start))
	f, _ := os.Create("image.png")
	png.Encode(f, img)
}
