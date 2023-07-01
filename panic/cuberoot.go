package main

import (
	"fmt"
	"math"
)

// A toy implementation of cube root using Newton's method.
func CubeRoot(x float64) float64 {
	z := x / 3 // Arbitrary initial value
	for i := 0; i < 1e6; i++ {
		prevz := z
		z -= (z*z*z - x) / (3 * z * z)
		if veryClose(z, prevz) {
			return z
		}
	}
	// A million iterations has not converged; something is wrong.
	panic(fmt.Sprintf("CubeRoot(%g) did not converge", x))
}

func veryClose(x, y float64) bool {
	return math.Abs(x-y) < 1.0e-10
}

func main() {
	fmt.Printf("%f ^ 1/3 = %f\n", 13.0, CubeRoot(13.0))
	fmt.Printf("%f ^ 1/3 = %f\n", -1.0, CubeRoot(-1.0))
}
