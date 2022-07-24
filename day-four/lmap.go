package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string] Vertex;

var ml = map[string]Vertex{
	// bisa di-remove juga type literal di value nya
	"Bell Labs": { 40.68433, -74.39967 },
	"Google": Vertex{ 37.42202, -122.08408 },
}

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{ 40.68433, -74.39967 }
	fmt.Println(m["Bell Labs"])

	value, isPresent := ml["Apple Inc."]
	if isPresent {
		fmt.Println("'Apple Inc.' exists!")
		fmt.Println(value)
	} else {
		fmt.Println("'Apple Inc.' does not exist")
	}

	value, isPresent = ml["Google"]
	if isPresent {
		fmt.Println("'Google' exists!")
		fmt.Println(value)
	} else {
		fmt.Println("'Google' does not exist")
	}
}
