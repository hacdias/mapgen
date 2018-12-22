package main

import (
	"fmt"
	"image/gif"
	"image/jpeg"
	"image/png"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/hacdias/mapgen"
	"github.com/spf13/pflag"
	"golang.org/x/image/bmp"
	"golang.org/x/image/tiff"
)

var (
	width       int
	height      int
	octaves     int
	seed        int64
	scale       float64
	persistence float64
	lacunarity  float64
	transition 	bool
	colorsFile 	string
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
	pflag.BoolVarP(&transition, "color transition", "t", false, "Color Transition in map generation overrites color palette transition")
	pflag.StringVarP(&colorsFile, "colors filename", "c", "defaultPalette.txt", "File name with the color palette")
	pflag.StringVarP(&filename, "filename", "f", "img.png", "File name to output")
}

func seedDefined() bool {
	v := false

	pflag.Visit(func(f *pflag.Flag) {
		if f.Name == "seed" {
			v = true
		}
	})

	return v
}

func main() {
	pflag.Parse()

	if persistence < 0 || persistence > 1 {
		fmt.Println("persistence must be between 0 and 1")
		os.Exit(1)
	}

	if !seedDefined() {
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		seed = int64(r1.Intn(100))
	}

	options := &mapgen.Options{
		Width:       	width,
		Height:      	height,
		Octaves:     	octaves,
		Seed:        	seed,
		Scale:       	scale,
		Persistence: 	persistence,
		Lacunarity:  	lacunarity,
		Transition:	 	transition,
		ColorsFile:		colorsFile,
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

	if strings.HasSuffix(filename, ".png") {
		err = png.Encode(f, img)
	} else if strings.HasSuffix(filename, ".gif") {
		err = gif.Encode(f, img, nil)
	} else if strings.HasSuffix(filename, ".jpg") || strings.HasSuffix(filename, ".jpeg") {
		err = jpeg.Encode(f, img, nil)
	} else if strings.HasSuffix(filename, ".bmp") {
		err = bmp.Encode(f, img)
	} else if strings.HasSuffix(filename, ".tiff") {
		err = tiff.Encode(f, img, nil)
	} else {
		fmt.Println("No support for the defined format. The file was saved as a PNG.")
		err = png.Encode(f, img)
	}

	if err != nil {
		panic(err)
	}
}
