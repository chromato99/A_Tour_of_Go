package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for i := z*z - x; i > 0.000001 || i < -0.000001; i = z*z - x {
		z -= i / (2 * z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
