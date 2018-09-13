package main

import (
	"math"

	"github.com/go-gl/mathgl/mgl32"
)

// Object is the overall "physical object" interface.
// These will be checked by CastObjectRay.
type Object interface {
	intersects(*Ray) *Hit
	material() *Material
}

// Sphere TODO.
type Sphere struct {
	position      mgl32.Vec3
	radius        float32
	savedMaterial *Material
}

// Light TODO.
type Light struct {
	position mgl32.Vec3
	color    ShadedColor
}

func (sphere *Sphere) material() *Material {
	return sphere.savedMaterial

}

func (sphere *Sphere) intersects(ray *Ray) *Hit {

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
		hit.object = sphere
		return hit
	}

	return nil
}
