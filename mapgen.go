package mapgen

import (
	"image"
	"math"

	noise "github.com/ojrac/opensimplex-go"
)

// Options are used to generate the image map. You can set
// the width, height, octaves, seed, scale, persistence
// and lacunatiry.
type Options struct {
	Width       int
	Height      int
	Octaves     int
	Seed        int64
	Scale       float64
	Persistence float64
	Lacunarity  float64
}

// Generate generates a image using the options object you defined.
// It returns an image that can be saved to a file or used to be
// later parsed.
func Generate(o *Options) (image.Image, error) {
	img := image.NewRGBA(image.Rect(0, 0, o.Width, o.Height))
	p := noise.NewNormalized(o.Seed)
	elev := makeMap(o.Width, o.Height)

	min := math.MaxFloat64
	max := -min

	for y := range elev {
		for x := range elev[y] {
			amplitude := 1.0
			frequency := 1.0
			noise := 0.0

			for i := 0; i < o.Octaves; i++ {
				sX := float64(x) / o.Scale * frequency
				sY := float64(y) / o.Scale * frequency
				noise += p.Eval2(sX, sY) * amplitude

				amplitude *= o.Persistence
				frequency *= o.Lacunarity
			}

			max = math.Max(noise, max)
			min = math.Min(noise, min)
			elev[y][x] = noise
		}
	}

	for y := range elev {
		for x := range elev[y] {
			noise := lerp(min, max, elev[y][x])
			color, err := biome(noise)

			if err != nil {
				return nil, err
			}

			img.Set(x, y, color)
		}
	}

	return img, nil
}

func makeMap(width, height int) [][]float64 {
	slice := make([][]float64, height)
	for i := range slice {
		slice[i] = make([]float64, width)
	}

	return slice
}

func lerp(v0, v1, t float64) float64 {
	return (t - v0) / (v1 - v0)
}
