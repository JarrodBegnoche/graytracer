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
	flag.UintVar(&width, "width", 480, "Width of rendered image")
	flag.UintVar(&height, "height", 270, "Height of rendered image")
	flag.Float64Var(&fov, "fov", math.Pi/3, "Field of View (in Radians)")
	flag.Parse()
	camera = components.MakeCamera(width, height, fov)
	camera.ViewTransform(primitives.MakePoint(0, 1.5, -4.5),
						 primitives.MakePoint(0, 1, 0),
						 primitives.MakeVector(0, 1, 0))
	ch = make(chan XY, 1000)
	imgMutex = &sync.Mutex{}
	start := time.Now()
	upLeft := image.Point{0, 0}
	lowRight := image.Point{int(width), int(height)}
	img = image.NewRGBA(image.Rectangle{upLeft, lowRight})
	world = components.MakeWorld()
	world.SetBackground(*patterns.MakeRGB(0.9, 0.9, 0.9))
	d4 := shapes.MakeGroup()
	d4.SetTransform(primitives.Translation(1.75, 0.5, 0.5).Multiply(
					primitives.RotationY(math.Pi / 2).Multiply(
					primitives.Translation(0, 1.0 / 3.0, 0))))
	d4_mat := patterns.Material{Pat: patterns.MakeRGB(0.8, 0, 0.8), Ambient: 0.1, Diffuse: 0.7, Specular: 0.5,
								Shininess: 200, Reflective: 0.1, Transparency: 0, RefractiveIndex: 0}
	d4_points := []primitives.PV{primitives.MakePoint(0, -1.0 / 3.0, math.Sqrt(8.0 / 9.0)),
								 primitives.MakePoint(math.Sqrt(2.0 / 3.0), -1.0 / 3.0, -math.Sqrt(2.0 / 9.0)),
								 primitives.MakePoint(-math.Sqrt(2.0 / 3.0), -1.0 / 3.0, -math.Sqrt(2.0 / 9.0)),
								 primitives.MakePoint(0, 1, 0)}
	d4_triangles := []*shapes.Triangle{shapes.MakeTriangle(d4_points[0], d4_points[1], d4_points[2]),
						 			   shapes.MakeTriangle(d4_points[0], d4_points[1], d4_points[3]),
									   shapes.MakeTriangle(d4_points[0], d4_points[2], d4_points[3]),
									   shapes.MakeTriangle(d4_points[1], d4_points[2], d4_points[3])}
	for _, triangle := range d4_triangles {
		triangle.SetMaterial(d4_mat)
		d4.AddShape(triangle)
	}
	world.AddObject(d4)
	d8 := shapes.MakeGroup()
	d8.SetTransform(primitives.Translation(0, 1, 0.5).Multiply(
					primitives.Scaling(0.75, 0.75, 0.75)))
	d8_mat := patterns.Material{Pat: patterns.MakeRGB(0, 0.8, 0), Ambient: 0.1, Diffuse: 0.7, Specular: 0.5,
								Shininess: 200, Reflective: 0.1, Transparency: 0, RefractiveIndex: 0}
	d8_points := []primitives.PV{primitives.MakePoint(0, 1, 0),
								 primitives.MakePoint(1, 0, 0),
								 primitives.MakePoint(0, 0, 1),
								 primitives.MakePoint(-1, 0, 0),
								 primitives.MakePoint(0, 0, -1),
								 primitives.MakePoint(0, -1, 0)}
	d8_triangles := []*shapes.Triangle{shapes.MakeTriangle(d8_points[0], d8_points[1], d8_points[2]),
									   shapes.MakeTriangle(d8_points[0], d8_points[2], d8_points[3]),
									   shapes.MakeTriangle(d8_points[0], d8_points[3], d8_points[4]),
									   shapes.MakeTriangle(d8_points[0], d8_points[4], d8_points[1]),
									   shapes.MakeTriangle(d8_points[5], d8_points[1], d8_points[2]),
									   shapes.MakeTriangle(d8_points[5], d8_points[2], d8_points[3]),
									   shapes.MakeTriangle(d8_points[5], d8_points[3], d8_points[4]),
									   shapes.MakeTriangle(d8_points[5], d8_points[4], d8_points[1])}
	for _, triangle := range d8_triangles {
		triangle.SetMaterial(d8_mat)
		d8.AddShape(triangle)
	}
	world.AddObject(d8)
	d20 := shapes.MakeGroup()
	d20.SetTransform(primitives.Translation(-2, 1, 2).Multiply(
					 primitives.Scaling(0.5, 0.5, 0.5)))
	d20_mat := patterns.Material{Pat: patterns.MakeRGB(0.9, 0, 0), Ambient: 0.1, Diffuse: 0.7, Specular: 0.5,
								 Shininess: 200, Reflective: 0.1, Transparency: 0, RefractiveIndex: 0}
	d20_points := []primitives.PV{primitives.MakePoint(0, -1, -math.Phi), // 0
								  primitives.MakePoint(0, -1, math.Phi),  // 1
								  primitives.MakePoint(0, 1, -math.Phi),  // 2
								  primitives.MakePoint(0, 1, math.Phi),   // 3
								  primitives.MakePoint(-1, -math.Phi, 0), // 4
								  primitives.MakePoint(-1, math.Phi, 0),  // 5
								  primitives.MakePoint(1, -math.Phi, 0),  // 6
								  primitives.MakePoint(1, math.Phi, 0),   // 7
								  primitives.MakePoint(-math.Phi, 0, -1), // 8
								  primitives.MakePoint(math.Phi, 0, -1),  // 9
								  primitives.MakePoint(-math.Phi, 0, 1),  // 10
								  primitives.MakePoint(math.Phi, 0, 1)}   // 11
	d20_triangles := []*shapes.Triangle{shapes.MakeTriangle(d20_points[0], d20_points[2], d20_points[8]), // front
										shapes.MakeTriangle(d20_points[0], d20_points[2], d20_points[9]),
										shapes.MakeTriangle(d20_points[1], d20_points[3], d20_points[10]), // back
										shapes.MakeTriangle(d20_points[1], d20_points[3], d20_points[11]),
										shapes.MakeTriangle(d20_points[9], d20_points[11], d20_points[6]), // right
										shapes.MakeTriangle(d20_points[9], d20_points[11], d20_points[7]),
										shapes.MakeTriangle(d20_points[8], d20_points[10], d20_points[4]), // left
										shapes.MakeTriangle(d20_points[8], d20_points[10], d20_points[5]),
										shapes.MakeTriangle(d20_points[4], d20_points[6], d20_points[0]), // bottom
										shapes.MakeTriangle(d20_points[4], d20_points[6], d20_points[1]),
										shapes.MakeTriangle(d20_points[4], d20_points[8], d20_points[0]), // bottom split
										shapes.MakeTriangle(d20_points[4], d20_points[10], d20_points[1]),
										shapes.MakeTriangle(d20_points[6], d20_points[9], d20_points[0]),
										shapes.MakeTriangle(d20_points[6], d20_points[11], d20_points[1]),
										shapes.MakeTriangle(d20_points[5], d20_points[7], d20_points[2]), // top
										shapes.MakeTriangle(d20_points[5], d20_points[7], d20_points[3]),
										shapes.MakeTriangle(d20_points[5], d20_points[8], d20_points[2]), //top split
										shapes.MakeTriangle(d20_points[5], d20_points[10], d20_points[3]),
										shapes.MakeTriangle(d20_points[7], d20_points[9], d20_points[2]),
										shapes.MakeTriangle(d20_points[7], d20_points[11], d20_points[3])}
	for _, triangle := range d20_triangles {
		triangle.SetMaterial(d20_mat)
		d20.AddShape(triangle)
	}
	world.AddObject(d20)
	light := components.PointLight{Intensity: patterns.MakeRGB(1, 1, 1),
								   Position: primitives.MakePoint(-1, 0.5, -5)}
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
