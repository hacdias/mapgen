package main

import (
	"image"
	"image/png"
	"math"
	"os"

	noise "github.com/ojrac/opensimplex-go"
)

func makeMap(width, height int) [][]float64 {
	slice := make([][]float64, height)
	for i := range slice {
		slice[i] = make([]float64, width)
	}

	return slice
}

func lerp(v0, v1, t float64) float64 {
	t = t - v0
	v1 = v1 - v0
	return t / v1
}

func main() {
	width := 3840
	height := 2160
	octaves := 5
	scale := 1080.0
	persistance := 0.5
	lacunarity := 2.6

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	p := noise.NewNormalized(16)

	elev := makeMap(width, height)

	min := math.MaxFloat64
	max := -min

	for y := range elev {
		elev[y] = make([]float64, width)
		for x := range elev[y] {
			amplitude := 1.0
			frequency := 1.0
			noise := 0.0

			for i := 0; i < octaves; i++ {
				sX := float64(x) / scale * frequency
				sY := float64(y) / scale * frequency
				noise += p.Eval2(sX, sY) * amplitude

				amplitude *= persistance
				frequency *= lacunarity
			}

			elev[y][x] = noise
			max = math.Max(noise, max)
			min = math.Min(noise, min)
		}
	}

	for y := range elev {
		for x := range elev[y] {
			noise := lerp(min, max, elev[y][x])

			color, _ := biome(noise)
			img.Set(x, y, color)
		}
	}

	f, err := os.Create("img.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	png.Encode(f, img)

}
