package mandelbrot

import (
	"fmt"
	"math/cmplx"
)

// "math" is not enough, we need "math/cmplx" for complex numbers

func calculateMandelbrot(maxIters int, xMin float64, xMax float64, yMin float64, yMax float64, width int, height int) [][]int {
	allRows := make([][]int, width)
	for imgY := 0; imgY < height; imgY++ {
		row := make([]int, height)
		for imgX := 0; imgX < width; imgX++ {
			cx := xMin + (xMax-xMin)*(float64(imgX)/float64(width))
			cy := yMin + (yMax-yMin)*(float64(imgY)/float64(height))
			escapedAt := mandelbrotAtPoint(cx, cy, maxIters)
			row[imgX] = escapedAt
		}
		allRows[imgY] = row
	}
	return allRows
}

//	fn mandelbrot_at_point(cx: f64, cy: f64, max_iters: usize) -> usize {
//	  let mut z = Complex { re: 0.0, im: 0.0 };
//	  let c = Complex::new(cx, cy);
//
//	  for i in 0..=max_iters {
//	      if z.norm() > 2.0 {
//	          return i;
//	      }
//	      z = z * z + c;
//	  }
//	  return max_iters;
//	}
func mandelbrotAtPoint(cx float64, cy float64, maxIters int) int {
	var z complex128
	c := complex(cx, cy) // complex is built-in type in Go
	minNorm := 2.0
	for i := 0; i <= maxIters; i++ {
		//zNorm := real(z)*real(z)+imag(z)*imag(z)    // Complex.norm() nor Complex.abs() are not available in Go, so we have to do it manually here or use math.Sqrt(real(z)*real(z)+imag(z)*imag(z))
		//zNorm := math.Sqrt(real(z)*real(z) + imag(z)*imag(z))
		zNorm := cmplx.Abs(z) // using cmplx package for clarity
		if zNorm > minNorm {
			return i
		}
		z = z*z + c
	}
	return maxIters
}

func render_mandelbrot(escapeVals [][]int) {
	for _, row := range escapeVals {
		line := ""
		for _, column := range row {
			val := ' '
			switch {
			case column <= 2:
				val = ' '
			case column <= 5:
				val = '.'
			case column <= 10:
				val = '•'
			case column <= 30:
				val = '*'
			case column <= 100:
				val = '+'
			case column <= 200:
				val = 'x'
			case column <= 400:
				val = '$'
			case column <= 700:
				val = '#'
			default:
				val = '%'
			}
			line += string(val)
		}
		fmt.Println(line)
	}
}

//func main() {
//    fmt.Println("ASCII/TEXT based mandelbrot (converted from RUST)!")
//    mandelbrot := calculateMandelbrot(1000, -2.0, 1.0, -1.0, 1.0, 100, 24)
//
//    render_mandelbrot(mandelbrot)
//}
