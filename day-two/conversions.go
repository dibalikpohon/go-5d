package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)

	// explicit type conversion
	fi := int(f)
	zf := float64(z)

	// cannot use imlpicity type conversion
	var abc int = int(math.Sqrt(80))


	fmt.Println(fi, zf, abc)
}
