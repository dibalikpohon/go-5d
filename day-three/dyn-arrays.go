package main

import "fmt"

func main() {

	pow := make([]int, 10)

	// iterate index, ignoring the value
	for i, _ := range pow {
		pow[i] = 1 << uint(i)
	}

	// iterate the value, ignoring the index
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}
