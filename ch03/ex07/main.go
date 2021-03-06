package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"math/cmplx"
	"os"
)

const (
	xmin, ymin, xmax, ymax = -2, -2, +2, +2
	width, height          = 1024, 1024
	threshold              = 1e-20
)

// +-1, +-iが解
var roots = []complex128{(1 + 0i), (-1 + 0i), (0 + 1i), (0 - 1i)}

func main() {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, newton(z))
		}
	}
	png.Encode(os.Stdout, img)
}

func newton(z complex128) color.Color {
	const iterations = 200
	for n := uint8(0); n < iterations; n++ {
		for i, r := range roots {
			abs := cmplx.Abs(z - r)
			if abs < threshold {
				return getColor(i, abs)
			}
		}
		z = (3 * z / 4) + (1 / (4 * z * z * z))
	}
	return color.Black
}

func getColor(i int, abs float64) color.Color {
	value := math.MaxUint8 * uint8((threshold-abs)/threshold)
	if i == 0 {
		return color.RGBA{value, 255 - value, 0, 128}
	} else if i == 1 {
		return color.RGBA{0, value, 255 - value, 128}
	} else if i == 2 {
		return color.RGBA{255 - value, 0, value, 128}
	} else if i == 3 {
		return color.RGBA{0, 0, 0, value}
	}
	fmt.Fprintf(os.Stderr, "Invalid index [%d]\n", i)
	return color.RGBA{128, 128, 128, 128}
}
