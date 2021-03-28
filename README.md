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
- Chapter 15: Triangles
- Chapter 16: Constructive Solid Geometry

## Latest Render

<img src="./image.png" width="800"/>

## Benchmark stats

### Benchmarks of different bounding box setups with 4096 spheres

pkg: github.com/factorion/graytracer/pkg/components
cpu: AMD Ryzen 7 2700 Eight-Core Processor
| Function | Iterations | Speed | Memory | Allocations |
| -------- | ---------- | ----- | ------ | ----------- |
| BenchmarkNoBoundingBoxes-16 | 1711 | 614977 ns/op | 5960 B/op | 113 allocs/op |
| Benchmark8BoundingBoxes-16 | 10000 | 117799 ns/op | 7336 B/op | 124 allocs/op |
| Benchmark64BoundingBoxes-16 | 32006 | 37650 ns/op | 7432 B/op | 129 allocs/op |

### Benchmarks of basic shape calculations

pkg: github.com/factorion/graytracer/pkg/shapes
cpu: AMD Ryzen 7 2700 Eight-Core Processor          
| Function | Iterations | Speed | Memory | Allocations |
| -------- | ---------- | ----- | ------ | ----------- |
| BenchmarkConeIntersection-16 | 5701870 | 207.7 ns/op | 72 B/op | 2 allocs/op |
| BenchmarkConeNormal-16 | 3989524 | 305.8 ns/op | 224 B/op | 5 allocs/op |
| BenchmarkCubeIntersection-16 | 9322422 | 130.6 ns/op | 48 B/op | 1 allocs/op |
| BenchmarkCubeNormal-16 | 3650877 | 319.7 ns/op | 224 B/op | 5 allocs/op |
| BenchmarkCylinderIntersection-16 | 5483966 | 214.8 ns/op | 72 B/op | 2 allocs/op |
| BenchmarkCylinderNormal-16 | 3969734 | 301.9 ns/op | 224 B/op | 5 allocs/op |
| BenchmarkPlaneIntersection-16 | 10618462 | 112.5 ns/op | 24 B/op | 1 allocs/op |
| BenchmarkPlaneNormal-16 | 4438746 | 273.1 ns/op | 224 B/op | 5 allocs/op |
| BenchmarkSphereIntersection-16 | 6149660 | 200.8 ns/op | 72 B/op | 2 allocs/op |
| BenchmarkSphereNormal-16 | 3968947 | 300.2 ns/op | 224 B/op | 5 allocs/op |
