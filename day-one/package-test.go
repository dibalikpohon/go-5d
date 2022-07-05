package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("print")
	// This line won't compile!
	// math.pow(2, 3)
	math.Pow(2, 3)
}
