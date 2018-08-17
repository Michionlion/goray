package main

import (
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
}

// Ray TODO.
type Ray struct {
	origin    mgl32.Vec3
	direction mgl32.Vec3
}

// HitType TODO.
type HitType int32

const (
	// EDGE TODO.
	EDGE HitType = 0

	// FULL TODO.
	FULL HitType = 1
)

// Hit TODO.
type Hit struct {
	position mgl32.Vec3
	hitType  HitType
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

	hit := &Hit{}

	if discriminant == 0 {
		hit.hitType = EDGE
	} else {
		hit.hitType = FULL
	}

	return hit
}
