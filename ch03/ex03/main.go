package main

import (
	"fmt"
	"math"
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
var minz = getMin()

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, az, af := corner(i+1, j)
			bx, by, bz, bf := corner(i, j)
			cx, cy, cz, cf := corner(i, j+1)
			dx, dy, dz, df := corner(i+1, j+1)
			zMean := (az + bz + cz + dz) / 4
			if af && bf && cf && df {
				fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' fill=\"#%s\"/>\n",
					ax, ay, bx, by, cx, cy, dx, dy, getColor(zMean))
			}
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64, float64, bool) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := f(x, y)
	if math.IsNaN(z) || math.IsInf(z, 0) {
		fmt.Fprintf(os.Stderr, "warning: return invalid value\n")
		return 0, 0, 0, false
	}

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, z, true
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

func getMin() float64 {
	minZ := 0.0
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			_, _, az, _ := corner(i+1, j)
			_, _, bz, _ := corner(i, j)
			_, _, cz, _ := corner(i, j+1)
			_, _, dz, _ := corner(i+1, j+1)
			zMean := (az + bz + cz + dz) / 4
			if zMean < minZ {
				minZ = zMean
			}
		}
	}
	return minZ
}

func getColor(z float64) string {
	// convert 0-1 -> color code
	redRateCode := fmt.Sprintf("%02X", int(16*16*(z-minz)/(1-minz)))
	blueRateCode := fmt.Sprintf("%02X", int(16*16*(1-((z-minz)/(1-minz)))))
	code := redRateCode + "00" + blueRateCode
	return code
}
