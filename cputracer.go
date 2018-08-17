package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"

	"github.com/go-gl/mathgl/mgl32"
)

// PixelsPerMeter is exactly what it says.
const PixelsPerMeter = 100

func main() {

	sphere := &Sphere{mgl32.Vec3{3, 5, 5}, 2}

	ray := &Ray{direction: mgl32.Vec3{0, 0, 1}}

	const Width = 1600
	const Height = 900

	image := image.NewRGBA(image.Rect(0, 0, Width, Height))

	for y := 0; y < Height; y++ {

		// Set the y-coordinate of the start position of the ray
		ray.origin[1] = float32(y) / PixelsPerMeter
		for x := 0; x < Width; x++ {
			// Set the x-coordinate of the start position of the ray
			ray.origin[0] = float32(x) / PixelsPerMeter

			hit := raySphereIntersection(ray, sphere)
			var colValue color.Color
			if hit == nil {
				// fmt.Printf("\u001b[31m  \u001b[0m")
				colValue = ShadedColor{0, 0, 0, 1} //ShadedColor{0, 0, 0, 1}
			} else {
				// fmt.Printf("\u001b[42m  \u001b[0m")
				fmt.Printf("Hit %v", hit)
				colValue = ShadedColor{0, 0.5, 0, 0.75} //ShadedColor{0, 1, 0, 1}
			}
			// fmt.Printf("putting %v", colValue)
			image.Set(x, y, colValue)
		}
		// First row of the screen done, move to next row
		// fmt.Printf("\n")
	}

	file, err := os.Create("image.png")
	if err != nil {
		log.Fatal(err)
	}

	err = png.Encode(file, image)
	if err != nil {
		file.Close()
		log.Fatal(err)
	}

	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}

}
