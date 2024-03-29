# graytracer
[![Build](https://github.com/factorion/graytracer/actions/workflows/Build.yml/badge.svg)](https://github.com/factorion/graytracer/actions/workflows/Build.yml)

Raytracer written in Go, based on the book [The Ray Tracer Challenge by Jamis Buck](https://pragprog.com/book/jbtracer/the-ray-tracer-challenge).  

## Current progress
- ✅ Chapter 1: Tuples, points and vectors
- ✅ Chapter 2: Drawing on a canvas
- ✅ Chapter 3: Matrices
- ✅ Chapter 4: Matrix transformations
- ✅ Chapter 5: [Ray-Sphere intersections](https://user-images.githubusercontent.com/40322086/108282866-54440b80-7150-11eb-886e-b7dce6254328.png)
- ✅ Chapter 6: [Light and Shading](https://user-images.githubusercontent.com/40322086/108282483-b3555080-714f-11eb-8ff8-66dd50fbd801.png)
- ✅ Chapter 7: [Making a Scene](https://user-images.githubusercontent.com/40322086/108283129-bdc41a00-7150-11eb-9f5c-587fb78044d9.png)
- ✅ Chapter 8: [Shadows](https://user-images.githubusercontent.com/40322086/108283364-214e4780-7151-11eb-9a9d-317127989193.png)
- ✅ Chapter 9: [Planes](https://user-images.githubusercontent.com/40322086/108283490-55c20380-7151-11eb-80ec-dfbab565d7d3.png)
- ✅ Chapter 10: [Patterns](https://user-images.githubusercontent.com/40322086/108283582-83a74800-7151-11eb-8810-708903002f40.png)
- ✅ Chapter 11: [Reflection and Refraction](https://user-images.githubusercontent.com/40322086/108283705-c832e380-7151-11eb-92f5-0ca6fe5b3bf3.png)
- ✅ Chapter 12: [Cubes](https://user-images.githubusercontent.com/40322086/108283784-ec8ec000-7151-11eb-9726-eb9bd1f61be9.png)
- ✅ Chapter 13: [Cylinders and Cones](https://user-images.githubusercontent.com/40322086/108651820-98a51380-7490-11eb-8519-c72a496c025c.png)
- ✅ Chapter 14: [Groups](https://user-images.githubusercontent.com/40322086/110737622-b5b14480-81fb-11eb-8b70-ff4517a84bac.png)
- ✅ Chapter 14.5: [Bounding Boxes](https://user-images.githubusercontent.com/40322086/112742776-6e4ae800-8f5f-11eb-8a4e-66a5d145fc3f.png)
- ✅ Chapter 15: [Triangles](https://user-images.githubusercontent.com/40322086/113527126-8be50680-958a-11eb-8521-6a738a7189c3.png)
- ✅ Chapter 15.5: OBJ files
    - [Teapot Low](https://user-images.githubusercontent.com/40322086/176288756-29e5b895-0347-4db1-9a5c-a869d2f04f4a.png)
    - [Teapot](https://user-images.githubusercontent.com/40322086/176288784-fdfb14af-df75-4a65-a749-6ed7335a1c7f.png)
    - [Teapot Low Smooth](https://user-images.githubusercontent.com/40322086/176288804-6b0b3cd4-062b-4e89-9655-d022c88dd911.png)
    - [Teapot Smooth](https://user-images.githubusercontent.com/40322086/176288815-4b13b22a-3c79-4ac7-89fb-fd8bd88351fa.png)
    - [Model 3](https://user-images.githubusercontent.com/40322086/176288834-cdb78bef-d27e-4db2-83dc-c09bedc88b3a.png)
        - [Model 3 OBJ source](http://dmi.chez-alice.fr/)
        - 737,161 Triangles
- ✅ Chapter 16: [Constructive Solid Geometry](https://user-images.githubusercontent.com/40322086/176577897-7eda8539-804b-4378-a0dc-ef09b2518232.png)

## Latest Render

<img src="./image.png" width="800"/>

## Benchmark stats

### Benchmarks of primitives calculation

pkg: github.com/factorion/graytracer/pkg/primitives
cpu: AMD Ryzen 7 5800H with Radeon Graphics
| Function | Iterations | Speed | Memory | Allocations |
| -------- | ---------: | ----: | -----: | ----------: |
| BenchmarkSubmatrix4x4-16 | 10416955 | 114.4 ns/op | 152 B/op | 4 allocs/op |
| BenchmarkMatrixMultiply-16 | 6347768 | 189.2 ns/op | 224 B/op | 5 allocs/op |
| BenchmarkMatrixDeterminant4x4-16 | 799813 | 1514 ns/op | 1568 B/op | 52 allocs/op |
| BenchmarkMatrixMinor3x3-16 | 14114458 | 83.44 ns/op | 80 B/op | 3 allocs/op |
| BenchmarkMatrixCofactor3x3-16 | 14621076 | 81.79 ns/op | 80 B/op | 3 allocs/op |
| BenchmarkMatrixInverse4x4-16 | 151344 | 7859 ns/op | 8064 B/op | 265 allocs/op |
| BenchmarkPVTransform-16 | 255785548 | 4.582 ns/op | 0 B/op | 0 allocs/op |
| BenchmarkPVReflect-16 | 707641048 | 1.715 ns/op | 0 B/op | 0 allocs/op |
| BenchmarkRayTransform-16 | 88858610 | 13.39 ns/op | 0 B/op | 0 allocs/op |

### Benchmarks of basic shape calculations

pkg: github.com/factorion/graytracer/pkg/shapes
cpu: AMD Ryzen 7 5800H with Radeon Graphics
| Function | Iterations | Speed | Memory | Allocations |
| -------- | ---------: | ----: | -----: | ----------: |
| BenchmarkConeIntersection-16 | 9989410 | 121.4 ns/op | 128 B/op | 2 allocs/op |
| BenchmarkConeNormal-16 | 7841389 | 155.3 ns/op | 224 B/op | 5 allocs/op |
| BenchmarkCubeIntersection-16 | 15381164 | 70.61 ns/op | 80 B/op | 1 allocs/op |
| BenchmarkCubeNormal-16 | 7521614 | 156.3 ns/op | 224 B/op | 5 allocs/op |
| BenchmarkCylinderIntersection-16 | 9446713 | 124.3 ns/op | 128 B/op | 2 allocs/op |
| BenchmarkCylinderNormal-16 | 7790467 | 151.1 ns/op | 224 B/op | 5 allocs/op |
| BenchmarkPlaneIntersection-16 | 19712168 | 60.18 ns/op | 48 B/op | 1 allocs/op |
| BenchmarkPlaneNormal-16 | 8302458 | 148.8 ns/op | 224 B/op | 5 allocs/op |
| BenchmarkSphereIntersection-16 | 9833896 | 120.2 ns/op | 128 B/op | 2 allocs/op |
| BenchmarkSphereNormal-16 | 7740237 | 150.4 ns/op | 224 B/op | 5 allocs/op |
| BenchmarkTriangleIntersection-16 | 16434596 | 71.02 ns/op | 48 B/op | 1 allocs/op |
| BenchmarkTriangleNormal-16 | 8217364 | 143.5 ns/op | 224 B/op | 5 allocs/op |

### Benchmarks of different bounding box setups with 4096 spheres

pkg: github.com/factorion/graytracer/pkg/components
cpu: AMD Ryzen 7 5800H with Radeon Graphics
| Function | Iterations | Speed | Memory | Allocations |
| -------- | ---------: | ----: | -----: | ----------: |
| BenchmarkNoBoundingBoxes-16 | 4136 | 291257 ns/op | 7936 B/op | 113 allocs/op |
| Benchmark8BoundingBoxes-16 | 18379 | 64370 ns/op | 10080 B/op | 124 allocs/op |
| Benchmark64BoundingBoxes-16 | 60115 | 20114 ns/op | 10240 B/op | 129 allocs/op |
