package components_test

import (
	"testing"

	"github.com/factorion/graytracer/pkg/components"
	"github.com/factorion/graytracer/pkg/patterns"
)

var DEFAULT string = components.Default_name

var mats map[string]patterns.Material = map[string]patterns.Material{DEFAULT: patterns.MakeDefaultMaterial()}

func TestGibberish(t *testing.T) {
	parsed_obj := components.ParseObjFile("gibberish.obj", false, mats)
	if len(parsed_obj.Vertices) != 0 {
		t.Error("Unexpected vertices found")
	}
	if len(parsed_obj.Faces) > 0 {
		t.Error("Unexpected triangle faces found")
	}
}

func TestVertices(t *testing.T) {
	parsed_obj := components.ParseObjFile("vertices.obj", false, mats)
	if len(parsed_obj.Vertices) != 4 {
		t.Errorf("Incorrect amount of vertices, found %v, expected 4", len(parsed_obj.Vertices))
		t.Errorf("Vertices: %v", parsed_obj.Vertices)
	}
	if len(parsed_obj.Faces) > 0 {
		t.Error("Unexpected triangle faces found")
	}
}

func TestFaces(t *testing.T) {
	parsed_obj := components.ParseObjFile("triangles.obj", false, mats)
	if len(parsed_obj.Vertices) != 4 {
		t.Errorf("Incorrect amount of vertices, found %v, expected 4", len(parsed_obj.Vertices))
		t.Errorf("Vertices: %v", parsed_obj.Vertices)
	}
	if len(parsed_obj.Faces[DEFAULT]) != 2 {
		t.Errorf("Incorrect amount of faces, found %v, expected 2", len(parsed_obj.Faces[DEFAULT]))
		t.Errorf("Faces: %v", parsed_obj.Faces[DEFAULT])
	}
	if !(parsed_obj.Faces[DEFAULT][0].Point1.Equals(parsed_obj.Vertices[0])) ||
		!(parsed_obj.Faces[DEFAULT][0].Point2.Equals(parsed_obj.Vertices[1])) ||
		!(parsed_obj.Faces[DEFAULT][0].Point3.Equals(parsed_obj.Vertices[2])) {
		t.Errorf("Incorrect vertices on first triangle: %v", parsed_obj.Faces[DEFAULT][0])
	}
	if !(parsed_obj.Faces[DEFAULT][1].Point1.Equals(parsed_obj.Vertices[0])) ||
		!(parsed_obj.Faces[DEFAULT][1].Point2.Equals(parsed_obj.Vertices[2])) ||
		!(parsed_obj.Faces[DEFAULT][1].Point3.Equals(parsed_obj.Vertices[3])) {
		t.Errorf("Incorrect vertices on second triangle: %v", parsed_obj.Faces[DEFAULT][0])
	}
}

func TestPolygons(t *testing.T) {
	parsed_obj := components.ParseObjFile("polygons.obj", false, mats)
	if len(parsed_obj.Vertices) != 5 {
		t.Errorf("Incorrect amount of vertices, found %v, expected 5", len(parsed_obj.Vertices))
		t.Errorf("Vertices: %v", parsed_obj.Vertices)
	}
	if len(parsed_obj.Faces[DEFAULT]) != 3 {
		t.Errorf("Incorrect amount of faces, found %v, expected 3", len(parsed_obj.Faces[DEFAULT]))
		t.Errorf("Faces: %v", parsed_obj.Faces[DEFAULT])
	}
	if !(parsed_obj.Faces[DEFAULT][0].Point1.Equals(parsed_obj.Vertices[0])) ||
		!(parsed_obj.Faces[DEFAULT][0].Point2.Equals(parsed_obj.Vertices[1])) ||
		!(parsed_obj.Faces[DEFAULT][0].Point3.Equals(parsed_obj.Vertices[2])) {
		t.Errorf("Incorrect vertices on first triangle: %v", parsed_obj.Faces[DEFAULT][0])
	}
	if !(parsed_obj.Faces[DEFAULT][1].Point1.Equals(parsed_obj.Vertices[0])) ||
		!(parsed_obj.Faces[DEFAULT][1].Point2.Equals(parsed_obj.Vertices[2])) ||
		!(parsed_obj.Faces[DEFAULT][1].Point3.Equals(parsed_obj.Vertices[3])) {
		t.Errorf("Incorrect vertices on second triangle: %v", parsed_obj.Faces[DEFAULT][0])
	}
	if !(parsed_obj.Faces[DEFAULT][2].Point1.Equals(parsed_obj.Vertices[0])) ||
		!(parsed_obj.Faces[DEFAULT][2].Point2.Equals(parsed_obj.Vertices[3])) ||
		!(parsed_obj.Faces[DEFAULT][2].Point3.Equals(parsed_obj.Vertices[4])) {
		t.Errorf("Incorrect vertices on third triangle: %v", parsed_obj.Faces[DEFAULT][0])
	}
}

func TestGroups(t *testing.T) {
	parsed_obj := components.ParseObjFile("groups.obj", false, mats)
	group1 := "FirstGroup"
	group2 := "SecondGroup"
	if len(parsed_obj.Vertices) != 4 {
		t.Errorf("Incorrect amount of vertices, found %v, expected 4", len(parsed_obj.Vertices))
		t.Errorf("Vertices: %v", parsed_obj.Vertices)
	}
	if len(parsed_obj.Faces[group1]) != 1 {
		t.Errorf("Incorrect amount of faces, found %v, expected 1", len(parsed_obj.Faces[group1]))
		t.Errorf("Faces: %v", parsed_obj.Faces[group1])
	}
	if len(parsed_obj.Faces[group2]) != 1 {
		t.Errorf("Incorrect amount of faces, found %v, expected 1", len(parsed_obj.Faces[group2]))
		t.Errorf("Faces: %v", parsed_obj.Faces[group2])
	}
	if !(parsed_obj.Faces[group1][0].Point1.Equals(parsed_obj.Vertices[0])) ||
		!(parsed_obj.Faces[group1][0].Point2.Equals(parsed_obj.Vertices[1])) ||
		!(parsed_obj.Faces[group1][0].Point3.Equals(parsed_obj.Vertices[2])) {
		t.Errorf("Incorrect vertices on %v triangle: %v", group1, parsed_obj.Faces[group1][0])
	}
	if !(parsed_obj.Faces[group2][0].Point1.Equals(parsed_obj.Vertices[0])) ||
		!(parsed_obj.Faces[group2][0].Point2.Equals(parsed_obj.Vertices[2])) ||
		!(parsed_obj.Faces[group2][0].Point3.Equals(parsed_obj.Vertices[3])) {
		t.Errorf("Incorrect vertices on %v triangle: %v", group2, parsed_obj.Faces[group2][0])
	}
}

func TestVertexNormals(t *testing.T) {
	parsed_obj := components.ParseObjFile("vertex_normals.obj", true, mats)
	if len(parsed_obj.Normals) != 3 {
		t.Errorf("Incorrect amount of vertex normals, found %v, expected 3", len(parsed_obj.Normals))
		t.Errorf("Vertex Normals: %v", parsed_obj.Normals)
	}
}

func TestSmoothTriangles(t *testing.T) {
	parsed_obj := components.ParseObjFile("smooth_triangles.obj", true, mats)
	if len(parsed_obj.Vertices) != 3 {
		t.Errorf("Incorrect amount of vertices, found %v, expected 3", len(parsed_obj.Vertices))
		t.Errorf("Vertices: %v", parsed_obj.Vertices)
	}
	if len(parsed_obj.Normals) != 3 {
		t.Errorf("Incorrect amount of vertex normals, found %v, expected 3", len(parsed_obj.Normals))
		t.Errorf("Vertex Normals: %v", parsed_obj.Normals)
	}
	if len(parsed_obj.Faces[DEFAULT]) != 2 {
		t.Errorf("Incorrect amount of faces, found %v, expected 3", len(parsed_obj.Faces[DEFAULT]))
		t.Errorf("Faces: %v", parsed_obj.Faces[DEFAULT])
	}
	if !(parsed_obj.Faces[DEFAULT][0].Point1.Equals(parsed_obj.Vertices[0])) ||
		!(parsed_obj.Faces[DEFAULT][0].Point2.Equals(parsed_obj.Vertices[1])) ||
		!(parsed_obj.Faces[DEFAULT][0].Point3.Equals(parsed_obj.Vertices[2])) {
		t.Errorf("Incorrect vertices on first triangle: %v", parsed_obj.Faces[DEFAULT][0])
	}
	if !(parsed_obj.Faces[DEFAULT][0].Normal1.Equals(parsed_obj.Normals[2])) ||
		!(parsed_obj.Faces[DEFAULT][0].Normal2.Equals(parsed_obj.Normals[0])) ||
		!(parsed_obj.Faces[DEFAULT][0].Normal3.Equals(parsed_obj.Normals[1])) {
		t.Errorf("Incorrect vertex normals on first triangle: %v", parsed_obj.Faces[DEFAULT][0])
	}
	if !(parsed_obj.Faces[DEFAULT][1].Point1.Equals(parsed_obj.Vertices[0])) ||
		!(parsed_obj.Faces[DEFAULT][1].Point2.Equals(parsed_obj.Vertices[1])) ||
		!(parsed_obj.Faces[DEFAULT][1].Point3.Equals(parsed_obj.Vertices[2])) {
		t.Errorf("Incorrect vertices on second triangle: %v", parsed_obj.Faces[DEFAULT][0])
	}
	if !(parsed_obj.Faces[DEFAULT][1].Normal1.Equals(parsed_obj.Normals[2])) ||
		!(parsed_obj.Faces[DEFAULT][1].Normal2.Equals(parsed_obj.Normals[0])) ||
		!(parsed_obj.Faces[DEFAULT][1].Normal3.Equals(parsed_obj.Normals[1])) {
		t.Errorf("Incorrect vertex normals on second triangle: %v", parsed_obj.Faces[DEFAULT][0])
	}
}
