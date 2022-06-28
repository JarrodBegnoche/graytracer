package components

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/factorion/graytracer/pkg/patterns"
	"github.com/factorion/graytracer/pkg/primitives"
	"github.com/factorion/graytracer/pkg/shapes"
)

var Default_name string = "DEFAULT"

type parsed_obj struct {
	Vertices []primitives.PV
	Normals  []primitives.PV
	Faces    map[string][]*shapes.Triangle
}

// Parse up to three values separated by a /
func ParseInts(value_string string) []int64 {
	parsed_ints := []int64{0, 0, 0}
	var parse_err error
	values := strings.Split(value_string, "/")
	for i := 0; i < 3; i++ {
		if len(values) > i {
			if values[i] == "" {
				continue
			}
			parsed_ints[i], parse_err = strconv.ParseInt(values[i], 10, 64)
			if parse_err != nil {
				fmt.Fprintf(os.Stderr, "Error parsing values: %s", value_string)
				os.Exit(1)
			}
		} else {
			break
		}
	}
	return parsed_ints
}

// Parse vertices and triangles from Wavefront OBJ file
func ParseObjFile(filename string, smooth bool, mats map[string]patterns.Material) *parsed_obj {
	name := Default_name
	mat_groups := make(map[string]uint64)
	material := mats[Default_name]
	result := &parsed_obj{
		Vertices: make([]primitives.PV, 0),
		Normals:  make([]primitives.PV, 0),
		Faces:    make(map[string][]*shapes.Triangle)}

	// Open wavefront OBJ file for parsing
	f, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing file %s", filename)
		os.Exit(1)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		// Get everything before a #
		line := strings.Split(scanner.Text(), "#")[0]
		fields := strings.Fields(line)
		field_length := len(fields)
		if field_length == 0 {
			continue
		}
		switch statement := fields[0]; statement {
		case "v":
			// Parse and verify the x, y, and z values of a vertex
			if field_length < 4 {
				fmt.Fprintf(os.Stderr, "Insufficient values for a vertex in line: %s", line)
				os.Exit(1)
			}
			x, x_err := strconv.ParseFloat(fields[1], 64)
			y, y_err := strconv.ParseFloat(fields[2], 64)
			z, z_err := strconv.ParseFloat(fields[3], 64)
			if (x_err != nil) || (y_err != nil) || (z_err != nil) {
				fmt.Fprintf(os.Stderr, "Error converting numbers in line: %s", line)
				os.Exit(1)
			}
			result.Vertices = append(result.Vertices, primitives.MakePoint(x, y, z))
		case "vn":
			// Parse and verify the x, y, and z values of a vertex normal
			if field_length < 4 {
				fmt.Fprintf(os.Stderr, "Insufficient values for a vertex normal in line: %s", line)
				os.Exit(1)
			}
			x, x_err := strconv.ParseFloat(fields[1], 64)
			y, y_err := strconv.ParseFloat(fields[2], 64)
			z, z_err := strconv.ParseFloat(fields[3], 64)
			if (x_err != nil) || (y_err != nil) || (z_err != nil) {
				fmt.Fprintf(os.Stderr, "Error converting numbers in line: %s", line)
				os.Exit(1)
			}
			result.Normals = append(result.Normals, primitives.MakeVector(x, y, z))
		case "f":
			field1 := ParseInts(fields[1])
			field2 := ParseInts(fields[2])
			for _, p3 := range fields[3:] {
				field3 := ParseInts(p3)
				var triangle *shapes.Triangle
				if field1[2] > 0 && field2[2] > 0 && field3[2] > 0 && smooth {
					triangle = shapes.MakeSmoothTriangle(result.Vertices[field1[0]-1],
						result.Vertices[field2[0]-1], result.Vertices[field3[0]-1],
						result.Normals[field1[2]-1], result.Normals[field2[2]-1],
						result.Normals[field3[2]-1])
				} else {
					triangle = shapes.MakeTriangle(result.Vertices[field1[0]-1],
						result.Vertices[field2[0]-1],
						result.Vertices[field3[0]-1])
				}
				triangle.SetMaterial(material)
				result.Faces[name] = append(result.Faces[name], triangle)
				field2 = field3
			}
		case "g":
			if fields[1] == "" {
				fmt.Fprintf(os.Stderr, "Blank group name given in line: %s", line)
				os.Exit(1)
			}
			name = fields[1]
		case "usemtl":
			if fields[1] == "" {
				fmt.Fprintf(os.Stderr, "Blank material name given in line: %s", line)
				os.Exit(1)
			}
			name = fields[1]
			if mat, ok := mats[name]; ok {
				material = mat
			} else {
				material = mats[Default_name]
			}
			if _, ok := result.Faces[name]; ok {
				mat_groups[name] += 1
				name = name + strconv.FormatUint(mat_groups[name], 10)
			} else {
				mat_groups[name] = 0
			}
		}
	}
	return result
}
