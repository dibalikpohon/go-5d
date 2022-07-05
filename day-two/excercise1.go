package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
    iter := 1
	z := 1.0
	for iter < 16 {
	    z -= (z * z - x) / (2 * z)
		iter++
	}
	
	return z
}

func main() {
	fmt.Println(Sqrt(516961))
}
