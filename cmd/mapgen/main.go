package main

import (
	"fmt"
	"image/png"
	"os"

	"github.com/hacdias/mapgen"
	flag "github.com/spf13/pflag"
)

var (
	width       int
	height      int
	octaves     int
	seed        int64
	scale       float64
	persistance float64
	lacunarity  float64
	filename    string
)

func init() {
	flag.IntVarP(&width, "width", "w", 100, "Width of image")
	flag.IntVarP(&height, "height", "h", 100, "Height of image")
	flag.IntVarP(&octaves, "octaves", "o", 5, "Number of octaves")
	flag.Int64VarP(&seed, "seed", "s", 0, "Seed to generate the map (default random)") //TODO: check if user defined
	flag.Float64VarP(&scale, "scale", "x", 20.0, "Scale")
	flag.Float64VarP(&persistance, "persistance", "p", 0.5, "Persistance")
	flag.Float64VarP(&lacunarity, "lacunarity", "l", 2.5, "Lacunarity")
	flag.StringVarP(&filename, "filename", "f", "img.png", "File name to output")
}

func main() {
	flag.Parse()

	if persistance < 0 || persistance > 1 {
		fmt.Println("Persistance must be between 0 and 1")
		os.Exit(1)
	}

	options := &mapgen.Options{
		Width:       width,
		Height:      height,
		Octaves:     octaves,
		Seed:        seed,
		Scale:       scale,
		Persistante: persistance,
		Lacunarity:  lacunarity,
	}

	img := mapgen.Generate(options)

	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	png.Encode(f, img)
}
