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
	camera.ViewTransform(primitives.MakePoint(-7, 1, 10),
		primitives.MakePoint(2, -2.5, 0),
		primitives.MakeVector(0, 1, 0))
	ch = make(chan XY, 1000)
	imgMutex = &sync.Mutex{}
	start := time.Now()
	upLeft := image.Point{0, 0}
	lowRight := image.Point{int(width), int(height)}
	img = image.NewRGBA(image.Rectangle{upLeft, lowRight})
	world = components.MakeWorld()
	world.SetBackground(*patterns.MakeRGB(0.9, 0.9, 0.9))
	floor := shapes.MakePlane()
	floor.SetTransform(primitives.Translation(0, -2.964808, 0))
	floor.SetMaterial(patterns.Material{Pat: patterns.MakeChecker(patterns.MakeRGB(0.95, 0.95, 0.95), patterns.MakeRGB(0.45, 0.45, 0.45)),
		Ambient: 0.4, Diffuse: 1.0, Specular: 0.6, Shininess: 200, Reflective: 0, Transparency: 0, RefractiveIndex: 1.0})
	world.AddObject(floor)
	tesla := shapes.MakeGroup()
	tesla.SetTransform(primitives.RotationX(-math.Pi / 2).Multiply(primitives.Scaling(-1, 1, 1)))
	tesla_mats := make(map[string]patterns.Material, 0)
	tesla_mats[components.Default_name] = patterns.Material{Pat: patterns.MakeRGB(1, 0.25, 1.0),
		Ambient: 0, Diffuse: 0.8, Specular: 0.5, Shininess: 200, Reflective: 0.1, Transparency: 0, RefractiveIndex: 1.0}
	tesla_mats["Body"] = patterns.Material{Pat: patterns.MakeRGB(0, 0, 0),
		Ambient: 0, Diffuse: 0.8, Specular: 0.6, Shininess: 16, Reflective: 0.05, Transparency: 0, RefractiveIndex: 1.0}
	tesla_mats["Body_noir"] = patterns.Material{Pat: patterns.MakeRGB(0.0980392, 0.0980392, 0.0980392),
		Ambient: 0, Diffuse: 0.8, Specular: 0.4, Shininess: 64, Reflective: 0, Transparency: 0, RefractiveIndex: 1.0}
	tesla_mats["Chrome"] = patterns.Material{Pat: patterns.MakeRGB(0.321569, 0.321569, 0.321569),
		Ambient: 0, Diffuse: 1.0, Specular: 0.6, Shininess: 18, Reflective: 0.4, Transparency: 0, RefractiveIndex: 1.0}
	tesla_mats["Dessous"] = patterns.Material{Pat: patterns.MakeRGB(0.180392, 0.180392, 0.180392),
		Ambient: 0, Diffuse: 1.0, Specular: 0, Shininess: 8, Reflective: 0, Transparency: 0, RefractiveIndex: 1.0}
	tesla_mats["Disk"] = patterns.Material{Pat: patterns.MakeRGB(0.235294, 0.239216, 0.227451),
		Ambient: 0, Diffuse: 0.8, Specular: 0.25, Shininess: 64, Reflective: 0, Transparency: 0, RefractiveIndex: 1.0}
	tesla_mats["Feux_blanc"] = patterns.Material{Pat: patterns.MakeRGB(0.152941, 0.14902, 0.14902),
		Ambient: 0, Diffuse: 1.0, Specular: 0.6, Shininess: 8, Reflective: 0, Transparency: 0.28, RefractiveIndex: 1.0}
	tesla_mats["Feux_chrome"] = patterns.Material{Pat: patterns.MakeRGB(0.192157, 0.192157, 0.192157),
		Ambient: 0, Diffuse: 1.0, Specular: 0.6, Shininess: 64, Reflective: 0.4, Transparency: 0, RefractiveIndex: 1.0}
	tesla_mats["Feux_glass"] = patterns.Material{Pat: patterns.MakeRGB(0.145098, 0.145098, 0.145098),
		Ambient: 0, Diffuse: 1.0, Specular: 0.5, Shininess: 64, Reflective: 0.1, Transparency: 0.505, RefractiveIndex: 1.0}
	tesla_mats["Feux_noir"] = patterns.Material{Pat: patterns.MakeRGB(0.145098, 0.145098, 0.145098),
		Ambient: 0, Diffuse: 1.0, Specular: 0.14, Shininess: 64, Reflective: 0, Transparency: 0, RefractiveIndex: 1.0}
	tesla_mats["Feux_rouge"] = patterns.Material{Pat: patterns.MakeRGB(0.138235, 0.0156863, 0.0156863),
		Ambient: 0, Diffuse: 1.0, Specular: 0.5, Shininess: 8, Reflective: 0.1, Transparency: 0.185, RefractiveIndex: 1.0}
	tesla_mats["Feux_rouge_b"] = patterns.Material{Pat: patterns.MakeRGB(0.138235, 0.0156863, 0.0156863),
		Ambient: 0, Diffuse: 0.8, Specular: 0.4, Shininess: 8, Reflective: 0.1, Transparency: 0.2, RefractiveIndex: 1.0}
	tesla_mats["Feux_stop"] = patterns.Material{Pat: patterns.MakeRGB(0.9, 0.1, 0.1),
		Ambient: 0, Diffuse: 1.0, Specular: 0.6, Shininess: 8, Reflective: 0.1, Transparency: 0.2, RefractiveIndex: 1.0}
	tesla_mats["Frein"] = patterns.Material{Pat: patterns.MakeRGB(0.372549, 0.105882, 0.129412),
		Ambient: 0, Diffuse: 0.9, Specular: 0.1, Shininess: 35, Reflective: 0, Transparency: 0, RefractiveIndex: 1.0}
	tesla_mats["Interieur"] = patterns.Material{Pat: patterns.MakeRGB(0.176471, 0.176471, 0.176471),
		Ambient: 0, Diffuse: 1.0, Specular: 0.1, Shininess: 64, Reflective: 0, Transparency: 0, RefractiveIndex: 1.0}
	tesla_mats["Interieur_bois"] = patterns.Material{Pat: patterns.MakeRGB(1.0, 1.0, 1.0),
		Ambient: 0, Diffuse: 1.0, Specular: 0.02, Shininess: 64, Reflective: 0, Transparency: 0, RefractiveIndex: 1.0}
	tesla_mats["Interieur_ceinture"] = patterns.Material{Pat: patterns.MakeRGB(0.219608, 0.219608, 0.219608),
		Ambient: 0, Diffuse: 1.0, Specular: 0.1, Shininess: 22, Reflective: 0, Transparency: 0, RefractiveIndex: 1.0}
	tesla_mats["Interieur_display"] = patterns.Material{Pat: patterns.MakeRGB(1, 1, 1),
		Ambient: 0, Diffuse: 0.8, Specular: 0.2, Shininess: 64, Reflective: 0, Transparency: 0, RefractiveIndex: 1.0}
	tesla_mats["Interieur_gris"] = patterns.Material{Pat: patterns.MakeRGB(0.752941, 0.752941, 0.752941),
		Ambient: 0, Diffuse: 1.0, Specular: 0, Shininess: 16, Reflective: 0, Transparency: 0, RefractiveIndex: 1.0}
	tesla_mats["Interieur_HP"] = patterns.Material{Pat: patterns.MakeRGB(0.0162745, 0.0162745, 0.0162745),
		Ambient: 0, Diffuse: 0.8, Specular: 0.2, Shininess: 16, Reflective: 0, Transparency: 0, RefractiveIndex: 1.0}
	tesla_mats["Interieur_light"] = patterns.Material{Pat: patterns.MakeRGB(0.784314, 0.784314, 0.784314),
		Ambient: 0, Diffuse: 0.9, Specular: 0.1, Shininess: 8, Reflective: 0, Transparency: 0.18, RefractiveIndex: 1.0}
	tesla_mats["Interieur_noir_b"] = patterns.Material{Pat: patterns.MakeRGB(0.04, 0.04, 0.04),
		Ambient: 0, Diffuse: 1.0, Specular: 0.6, Shininess: 16, Reflective: 0, Transparency: 0, RefractiveIndex: 1.0}
	tesla_mats["Interieur_porte"] = patterns.Material{Pat: patterns.MakeRGB(0.244706, 0.244706, 0.244706),
		Ambient: 0, Diffuse: 1.0, Specular: 0, Shininess: 4, Reflective: 0, Transparency: 0, RefractiveIndex: 1.0}
	tesla_mats["Interieur_sol"] = patterns.Material{Pat: patterns.MakeRGB(1, 1, 1),
		Ambient: 0, Diffuse: 1.0, Specular: 0, Shininess: 4, Reflective: 0, Transparency: 0, RefractiveIndex: 1.0}
	tesla_mats["Jante"] = patterns.Material{Pat: patterns.MakeRGB(0.515098, 0.515098, 0.515098),
		Ambient: 0, Diffuse: 1.0, Specular: 0.6, Shininess: 16, Reflective: 0.1, Transparency: 0, RefractiveIndex: 1.0}
	tesla_mats["Jante_noir"] = patterns.Material{Pat: patterns.MakeRGB(0.145098, 0.145098, 0.145098),
		Ambient: 0, Diffuse: 0.8, Specular: 0.1, Shininess: 64, Reflective: 0, Transparency: 0, RefractiveIndex: 1.0}
	tesla_mats["Mecanique"] = patterns.Material{Pat: patterns.MakeRGB(0.211765, 0.211765, 0.211765),
		Ambient: 0, Diffuse: 0.8, Specular: 0.25, Shininess: 64, Reflective: 0, Transparency: 0, RefractiveIndex: 1.0}
	tesla_mats["Metal_alu"] = patterns.Material{Pat: patterns.MakeRGB(0.498039, 0.498039, 0.498039),
		Ambient: 0, Diffuse: 1.0, Specular: 0.02, Shininess: 64, Reflective: 0, Transparency: 0, RefractiveIndex: 1.0}
	tesla_mats["Metal_noir"] = patterns.Material{Pat: patterns.MakeRGB(0.188235, 0.188235, 0.188235),
		Ambient: 0, Diffuse: 0.8, Specular: 0.25, Shininess: 64, Reflective: 0, Transparency: 0, RefractiveIndex: 1.0}
	tesla_mats["Miroir"] = patterns.Material{Pat: patterns.MakeRGB(0.254902, 0.254902, 0.254902),
		Ambient: 0, Diffuse: 0.8, Specular: 0.5, Shininess: 64, Reflective: 0.9, Transparency: 0, RefractiveIndex: 1.0}
	tesla_mats["Partie_noir"] = patterns.Material{Pat: patterns.MakeRGB(0, 0, 0),
		Ambient: 0, Diffuse: 1.0, Specular: 0, Shininess: 8, Reflective: 0, Transparency: 0, RefractiveIndex: 1.0}
	tesla_mats["Phare"] = patterns.Material{Pat: patterns.MakeRGB(0.294118, 0.294118, 0.294118),
		Ambient: 0, Diffuse: 0.8, Specular: 0.6, Shininess: 56, Reflective: 0, Transparency: 0, RefractiveIndex: 1.0}
	tesla_mats["Phare_alu"] = patterns.Material{Pat: patterns.MakeRGB(0.490196, 0.490196, 0.490196),
		Ambient: 0, Diffuse: 0.9, Specular: 0.4, Shininess: 16, Reflective: 0, Transparency: 0, RefractiveIndex: 1.0}
	tesla_mats["Phare_blanc"] = patterns.Material{Pat: patterns.MakeRGB(0.152471, 0.152471, 0.152471),
		Ambient: 0, Diffuse: 1.0, Specular: 0.2, Shininess: 8, Reflective: 0, Transparency: 0.14, RefractiveIndex: 1.0}
	tesla_mats["Phare_orange"] = patterns.Material{Pat: patterns.MakeRGB(0.466667, 0.227451, 0.12549),
		Ambient: 0, Diffuse: 1.0, Specular: 0.2, Shininess: 8, Reflective: 0, Transparency: 0, RefractiveIndex: 1.0}
	tesla_mats["Phare_feux_b"] = patterns.Material{Pat: patterns.MakeRGB(0.533333, 0.552941, 0.568627),
		Ambient: 0, Diffuse: 1.0, Specular: 0.2, Shininess: 16, Reflective: 0, Transparency: 0.52, RefractiveIndex: 1.0}
	tesla_mats["Phare_vitre"] = patterns.Material{Pat: patterns.MakeRGB(0.317647, 0.329412, 0.341176),
		Ambient: 0, Diffuse: 1.0, Specular: 0.25, Shininess: 16, Reflective: 0.1, Transparency: 0.6, RefractiveIndex: 1.0}
	tesla_mats["Plastique_noir"] = patterns.Material{Pat: patterns.MakeRGB(0.121569, 0.121569, 0.121569),
		Ambient: 0, Diffuse: 0.9, Specular: 0.15, Shininess: 64, Reflective: 0, Transparency: 0, RefractiveIndex: 1.0}
	tesla_mats["Plastique_rouge"] = patterns.Material{Pat: patterns.MakeRGB(0.752941, 0.247059, 0.172549),
		Ambient: 0, Diffuse: 0.9, Specular: 0.15, Shininess: 8, Reflective: 0, Transparency: 0, RefractiveIndex: 1.0}
	tesla_mats["Pneu_text"] = patterns.Material{Pat: patterns.MakeRGB(0.0203922, 0.0203922, 0.0203922),
		Ambient: 0, Diffuse: 0.9, Specular: 0.1, Shininess: 11, Reflective: 0, Transparency: 0, RefractiveIndex: 1.0}
	tesla_mats["Siege"] = patterns.Material{Pat: patterns.MakeRGB(0.164706, 0.164706, 0.164706),
		Ambient: 0, Diffuse: 0.8, Specular: 0.3, Shininess: 35, Reflective: 0, Transparency: 0, RefractiveIndex: 1.0}
	tesla_mats["Vitre"] = patterns.Material{Pat: patterns.MakeRGB(0.227451, 0.278431, 0.27451),
		Ambient: 0, Diffuse: 1.0, Specular: 0.4, Shininess: 16, Reflective: 0, Transparency: 0.66, RefractiveIndex: 1.0}
	tesla_mats["Vitre_toit"] = patterns.Material{Pat: patterns.MakeRGB(0.0392157, 0.0392157, 0.0392157),
		Ambient: 0, Diffuse: 1.0, Specular: 0.4, Shininess: 16, Reflective: 0, Transparency: 0.2, RefractiveIndex: 1.0}
	tesla_pieces := make(map[string]*shapes.Group, 0)
	tesla_obj := components.ParseObjFile("Tesla Model 3.obj", true, tesla_mats)
	for key, element := range tesla_obj.Faces {
		tesla_pieces[key] = shapes.MakeGroup()
		for _, triangle := range element {
			tesla_pieces[key].AddShape(triangle)
		}
		tesla.AddShape(tesla_pieces[key])
	}
	world.AddObject(tesla)
	light := components.PointLight{Intensity: patterns.MakeRGB(1, 1, 1),
		Position: primitives.MakePoint(-10, 5, 15)}
	world.AddLight(light)
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
