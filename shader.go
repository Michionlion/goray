package main

// ShadedColor is a percentage based color implementation
type ShadedColor struct {
	r, g, b, a float32
}

func clamp(value, min, max float32) float32 {
	if value > max {
		return max
	} else if value < min {
		return min
	}

	return value
}

// RGBA converts shaded color to RGBA
func (sc ShadedColor) RGBA() (r, g, b, a uint32) {
	r = uint32(clamp(sc.r, 0, 1) * 255)
	r |= r << 8
	g = uint32(clamp(sc.g, 0, 1) * 255)
	g |= g << 8
	b = uint32(clamp(sc.b, 0, 1) * 255)
	b |= b << 8
	a = uint32(clamp(sc.a, 0, 1) * 255)
	a |= a << 8
	return
}
