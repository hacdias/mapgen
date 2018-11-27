package mapgen

import (
	"image/color"

	"github.com/lucasb-eyer/go-colorful"
)

func biome(e float64) (color.Color, error) {
	if e < 0.3 {
		return colorful.Hex("#2c52a0")
	} else if e < 0.4 {
		return colorful.Hex("#3766c8")
	} else if e < 0.45 {
		return colorful.Hex("#d0d080")
	} else if e < 0.55 {
		return colorful.Hex("#589619")
	} else if e < 0.60 {
		return colorful.Hex("#426220")
	} else if e < 0.70 {
		return colorful.Hex("#5c453e")
	} else if e < 0.90 {
		return colorful.Hex("#4d3b39")
	} else {
		return colorful.Hex("#ffffff")
	}
}
