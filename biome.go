package mapgen

import (
	"bufio"
	"image/color"
	"os"
	"strconv"
	"strings"

	"github.com/lucasb-eyer/go-colorful"
)

// GradientTable contains the "keypoints" of the colorgradient you want to generate.
// The position of each keypoint has to live in the range [0,1]
type GradientTable []struct {
	Color      colorful.Color
	Pos        float64
	Transition bool
}

func crunchSplitFunc(data []byte, atEOF bool) (advance int, token []byte, err error) {
	// Return nothing if at end of file and no data passed
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	// Find the index of the input of a ";" or "\n" accordingly
	i := strings.Index(string(data), ";")
	j := strings.Index(string(data), "\n")
	if j < i && j >= 0 {
		return j + 1, data[0:j], nil
	} else if i >= 0 {
		return i + 1, data[0:i], nil
	}

	// If at end of file with data return the data
	if atEOF {
		return len(data), data, nil
	}

	return
}

// GenerateGradient generates a gradient with the information given,
// in filename, by user
func GenerateGradient(filename string, transitionFlag bool) GradientTable {
	file, err := os.Open(filename)
	check(err)

	scannerLines := bufio.NewScanner(file)
	scanner := bufio.NewScanner(file)
	scanner.Split(crunchSplitFunc)

	lines := 0
	for scannerLines.Scan() {
		lines++
	}

	gradient := make(GradientTable, lines)

	i := 0
	file.Seek(0, 0)
	for scanner.Scan() {
		gradient[i].Color = ParseHex(scanner.Text())

		scanner.Scan()
		gradient[i].Pos, err = strconv.ParseFloat(scanner.Text(), 64)
		check(err)

		scanner.Scan()
		gradient[i].Transition, err = strconv.ParseBool(scanner.Text())
		check(err)
		gradient[i].Transition = gradient[i].Transition || transitionFlag

		i++
	}

	return gradient

}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

// ParseHex parses string to hex color
func ParseHex(s string) colorful.Color {
	c, err := colorful.Hex(s)
	check(err)

	return c
}

// GradientTable biome constrcuts the biome according to the gradientTable given
func (self GradientTable) biome(e float64) (color.Color, error) {
	for i := 0; i < len(self)-1; i++ {
		c1 := self[i]
		c2 := self[i+1]
		if c1.Pos <= e && e <= c2.Pos {
			if c1.Transition {
				e := (e - c1.Pos) / (c2.Pos - c1.Pos)
				return c1.Color.BlendLab(c2.Color, e).Clamped(), nil
			}

			if e-c1.Pos < c2.Pos-e {
				return c1.Color, nil
			} else {
				return c2.Color, nil
			}
		}
	}

	return self[len(self)-1].Color, nil
}
