package main

import "fmt"

var c, python, java = true, false, "no!"

func main() {
	var i, j int = 1, 2

	// shorthand declarations???
	// inside a function only
	cpp, matlab := 2, true

	fmt.Println(i, c, python, java, j, matlab, cpp)
}
