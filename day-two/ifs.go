package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))

	// short-hand if statement
	if result := false; result {
		fmt.Println(!result)
		fmt.Println("true")
	} else {
		fmt.Println(!result)
		fmt.Println("false")
	}

	// `result` cannot be used
	fmt.Println(result)
}
