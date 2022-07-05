package main

import "fmt"

func main() {
	defer fmt.Println("abcd")
	defer fmt.Println("world")

	fmt.Println("hello")
}
