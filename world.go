package main

import (
	"container/list"

	"github.com/go-gl/mathgl/mgl32"
)

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
	object   Object
}

// World holds the objects and lights in a scene
type World struct {
	objects *list.List
	lights  *list.List
}

// NewWorld creates a new, empty world.
func NewWorld() *World {
	world := &World{
		objects: list.New(),
		lights:  list.New(),
	}
	return world
}

// Add adds an object to the world.
func (world *World) Add(object Object) {
	world.objects.PushBack(object)
}

// AddLight adds a light to the world.
func (world *World) AddLight(light *Light) {
	world.lights.PushBack(light)
}

// CastObjectRay casts a given ray to check for collisions with objects.
// It will return the nearest hit.
func (world *World) CastObjectRay(ray *Ray) *Hit {
	return nil
}

// CastLightRays casts rays to every light.
func (world *World) CastLightRays(from *Hit) *Hit {
	return nil
}
