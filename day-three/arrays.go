package main

import "fmt"

func main() {
	
	var str[2] string
	str[0] = "Hello"
	str[1] = "Go"

	fmt.Println(str[0], str[1])
	fmt.Println(str)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	slice := primes[2:4]
	fmt.Println(slice)
}
