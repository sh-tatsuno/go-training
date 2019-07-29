package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

const (
	xmin, ymin, xmax, ymax = -2, -2, +2, +2
	width, height          = 1024, 1024
)

func main() {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	colors := []color.Color{}
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)

			// 色の平均をとる
			if len(colors) < 4 {
				colors = append(colors, mandelbrot(z))
			} else {
				colors = append(colors[1:], mandelbrot(z))
			}
			img.Set(px, py, avgColor(colors))
		}
	}
	png.Encode(os.Stdout, img)
}

func avgColor(colors []color.Color) color.Color {
	r, g, b, a := 0, 0, 0, 0
	for _, color := range colors {
		rr, gg, bb, aa := color.RGBA()
		r += int(rr)
		g += int(gg)
		b += int(bb)
		a += int(aa)
	}
	return color.RGBA{uint8(r / len(colors)), uint8(g / len(colors)),
		uint8(b / len(colors)), uint8(a / len(colors))}
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			r := uint8(contrast * n)
			g := uint8(contrast * 2 * n)
			b := uint8(contrast * 3 * n)
			return color.RGBA{r, g, b, 128}
		}
	}
	return color.Black
}
