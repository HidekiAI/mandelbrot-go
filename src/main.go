package main

import (
	"fmt"
)

func main() {
	fmt.Println("ASCII/TEXT based mandelbrot (converted from RUST)!")
	mandelbrot := calculateMandelbrot(1000, -2.0, 1.0, -1.0, 1.0, 100, 24)

	render_mandelbrot(mandelbrot)
}
