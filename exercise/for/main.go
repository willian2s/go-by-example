package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for i := 1; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Printf("Iteration %d: z = %f\n", i+1, z)
	}
	return z
}

func main() {
	for _, x := range []float64{1, 2, 3, 4, 9, 16} {
		result := Sqrt(x)
		fmt.Printf("Computed sqrt(%f) = %f, math.Sqrt = %f\n", x, result, math.Sqrt(x))
	}
}
