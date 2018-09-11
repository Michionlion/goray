package main

import (
	"container/list"
	"math"

	"github.com/go-gl/mathgl/mgl32"
)

// Voxel struct.
// type Voxel struct {
// 	col color.Color
// }
//
// // World struct.
// type World struct {
// 	voxels [][][]Voxel
// }
//
// func (world *World) getVoxel(x, y, z int) *Voxel {
// 	return &world.voxels[x][y][z]
// }

// Sphere TODO.
type Sphere struct {
	position mgl32.Vec3
	radius   float32
	material *Material
}

// Light TODO.
type Light struct {
	position mgl32.Vec3
	color    ShadedColor
}

// Material TODO.
type Material struct {
	diffuse      ShadedColor
	reflectivity float32
}

// Ray TODO.
type Ray struct {
	origin    mgl32.Vec3
	direction mgl32.Vec3
}

// Hit TODO.
type Hit struct {
	position mgl32.Vec3
	distance float32
}

type World struct {
	spheres *list.List
}

func NewWorld() {
	world := World{spheres: list.New()}
}

// // WorldRenderer struct.
// type WorldRenderer struct {
// 	// world *World
// 	sphereRadius int
// 	spherePos    mgl32.Vec3
// }

// func (wr *WorldRenderer) ray(dir *mgl32.Vec3) *mgl32.Vec3 {
//
// 	// first pass -- ignore dir
//
// 	return nil
// }

func raySphereIntersection(ray *Ray, sphere *Sphere) *Hit {

	A := ray.direction.Dot(ray.direction)

	dist := ray.origin.Sub(sphere.position)

	B := 2 * ray.direction.Dot(dist)

	C := dist.Dot(dist) - sphere.radius*sphere.radius

	discriminant := B*B - 4*A*C

	if discriminant < 0 {
		return nil
	}

	discriminantSqrt := float32(math.Sqrt(float64(discriminant)))

	t0 := (-B + discriminantSqrt) / 2
	t1 := (-B - discriminantSqrt) / 2

	if t0 > t1 {
		t0 = t1
	}

	if t0 > 0.0001 {
		hit := &Hit{}
		hit.position = ray.origin.Add(ray.direction.Mul(t0))
		hit.distance = t0
		return hit
	}

	return nil
}
