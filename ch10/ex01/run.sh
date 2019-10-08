go run mandelbrot/main.go | go run convert/main.go --fmt png  >> out.png
go run mandelbrot/main.go | go run convert/main.go --fmt gif >> out.gif
go run mandelbrot/main.go | go run convert/main.go --fmt jpeg >> out.jpeg