package main

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"log"
	"math/cmplx"
	"net/http"
	"net/url"
	"strconv"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		showPng(w, r.URL)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

const (
	width, height          = 1024, 1024
	xmin, ymin, xmax, ymax = -2, -2, 2, 2
)

func showPng(w io.Writer, u *url.URL) {
	q, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		log.Fatal(err)
	}
	width, _ := strconv.ParseFloat(q["width"][0], 64)
	height, _ := strconv.ParseFloat(q["height"][0], 64)
	r, _ := strconv.ParseFloat(q["rate"][0], 64)
	img := image.NewRGBA(image.Rect(0, 0, int(width), int(height)))
	for py := 0; py < int(height); py++ {
		y := (float64(py)/float64(height)*(ymax-ymin) + ymin) * r
		for px := 0; px < int(width); px++ {
			x := (float64(px)/float64(width)*(xmax-xmin) + xmin) * r
			z := complex(x, y)
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(w, img)
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
