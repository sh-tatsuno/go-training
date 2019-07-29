// package main

// import (
// 	"image"
// 	"image/color"
// 	"image/png"
// 	"io"
// 	"log"
// 	"math/cmplx"
// 	"net/http"
// 	"net/url"
// 	"os"
// 	"strconv"
// )

// func main() {
// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		w.Header().Set("Content-Type", "image/svg+xml")
// 		showPng(w, r.URL)
// 	})
// 	log.Fatal(http.ListenAndServe("localhost:8000", nil))
// }

// func showPng(w io.Writer, u *url.URL) {
// 	q, err := url.ParseQuery(u.RawQuery)
// 	if err != nil {
// 		panic(err)
// 	}
// 	r, err := strconv.Atoi(q["r"][0])
// 	if err != nil {
// 		panic(err)
// 	}
// 	width, err := strconv.Atoi(q["width"][0])
// 	if err != nil {
// 		panic(err)
// 	}
// 	height, err := strconv.Atoi(q["height"][0])
// 	if err != nil {
// 		panic(err)
// 	}

// 	xmin, ymin, xmax, ymax := -r, -r, r, r
// 	img := image.NewRGBA(image.Rect(0, 0, width, height))
// 	for py := 0; py < height; py++ {
// 		y := float64(py) / float64(height*(ymax-ymin)+ymin)
// 		for px := 0; px < width; px++ {
// 			x := float64(px) / float64(width*(xmax-xmin)+xmin)
// 			z := complex(x, y)
// 			img.Set(px, py, mandelbrot(z))
// 		}
// 	}
// 	png.Encode(os.Stdout, img)
// }

// func mandelbrot(z complex128) color.Color {
// 	const iterations = 200
// 	const contrast = 15

// 	var v complex128
// 	for n := uint8(0); n < iterations; n++ {
// 		v = v*v + z
// 		if cmplx.Abs(v) > 2 {
// 			r := uint8(contrast * n)
// 			g := uint8(contrast * 2 * n)
// 			b := uint8(contrast * 3 * n)
// 			return color.RGBA{r, g, b, 128}
// 		}
// 	}
// 	return color.Black
// }
