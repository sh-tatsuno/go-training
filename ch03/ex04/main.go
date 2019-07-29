package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"net/url"
	"os"
)

const (
	width, height = 600, 320
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/svg+xml")
		showSvg(w, r.URL)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func showSvg(w io.Writer, u *url.URL) {
	q, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", q["width"], q["height"])
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, af := corner(i+1, j)
			bx, by, bf := corner(i, j)
			cx, cy, cf := corner(i, j+1)
			dx, dy, df := corner(i+1, j+1)
			if af && bf && cf && df {
				fmt.Fprintf(w, "<polygon points='%g,%g %g,%g %g,%g %g,%g' fill=\"%s\"/>\n",
					ax, ay, bx, by, cx, cy, dx, dy, q["color"][0])
			}
		}
	}
	fmt.Fprintf(w, "</svg>")
}

func corner(i, j int) (float64, float64, bool) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := f(x, y)
	if math.IsNaN(z) || math.IsInf(z, 0) {
		fmt.Fprintf(os.Stderr, "warning: return invalid value\n")
		return 0, 0, false
	}

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, true
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}
