package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func shortened_add(x, y int) int {
	return x + y
}

// return multiple results
func swap(x, y string) (string, string) {
	return y, x
}

// named-return ?? this is cool
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x

	return
}

func main() {
	fmt.Println("2 + 3 = ", add(2, 3))
	fmt.Println("3 + 3 = ", shortened_add(3, 3))
	a, b := swap("hello", "world")
	fmt.Println("swap ('hello', 'world')", b, a)
	fmt.Println(split(17))
}
