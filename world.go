package main

import (
	"image/color"

	"github.com/stojg/vector"
)

// Voxel struct.
type Voxel struct {
	col color.Color
}

// World struct.
type World struct {
	voxels [][][]Voxel
}

func (world *World) getVoxel(x, y, z int) *Voxel {
	return &world.voxels[x][y][z]
}

// WorldRenderer struct.
type WorldRenderer struct {
	// world *World
	floorY  int
	sphereX int
	sphereY int
	sphereZ int
	sphereR int
}

func (wr *WorldRenderer) intersect(origin, dir *vector.Vector3) *vector.Vector3 {

	return nil
}
