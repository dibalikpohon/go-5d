package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	// Struct literal
	fmt.Println(Vertex{1, 2})
	
	// accessing struct
	v := Vertex{30, 40}
	fmt.Println(v.Y)

	// pointer to struct 👀
	pv := &v
	
	// use dot to access the pointer
	pv.X = 11
	
	// or C-style dereferencing
	(*pv).Y = 33

	fmt.Println(*pv)
	fmt.Println(pv)

	v1 := Vertex{Y: 40, X: 1}
	fmt.Println(v1)
}
