package main

import (
	"fmt"
	"image/png"
	"os"

	"github.com/hacdias/mapgen"
	"github.com/spf13/pflag"
)

var (
	width       int
	height      int
	octaves     int
	seed        int64
	scale       float64
	persistence float64
	lacunarity  float64
	filename    string
)

func init() {
	pflag.IntVarP(&width, "width", "w", 100, "Width of image")
	pflag.IntVarP(&height, "height", "h", 100, "Height of image")
	pflag.IntVarP(&octaves, "octaves", "o", 5, "Number of octaves")
	pflag.Int64VarP(&seed, "seed", "s", 0, "Seed to generate the map (default random)") //TODO: check if user defined
	pflag.Float64VarP(&scale, "scale", "x", 20.0, "Scale")
	pflag.Float64VarP(&persistence, "persistence", "p", 0.5, "persistence")
	pflag.Float64VarP(&lacunarity, "lacunarity", "l", 2.5, "Lacunarity")
	pflag.StringVarP(&filename, "filename", "f", "img.png", "File name to output")
}

func main() {
	pflag.Parse()

	if persistence < 0 || persistence > 1 {
		fmt.Println("persistence must be between 0 and 1")
		os.Exit(1)
	}

	options := &mapgen.Options{
		Width:       width,
		Height:      height,
		Octaves:     octaves,
		Seed:        seed,
		Scale:       scale,
		Persistence: persistence,
		Lacunarity:  lacunarity,
	}

	img, err := mapgen.Generate(options)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	f, err := os.Create(filename)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer f.Close()
	png.Encode(f, img)
}
