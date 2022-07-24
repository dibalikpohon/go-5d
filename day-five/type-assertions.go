package main

import "fmt"

func main() {
	var i interface{} = "hello"

	s := i.(string)     // s returns the underlying value
	fmt.Println(s)

	s, ok := i.(string)  // ok reports whether the assertion
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	f = i.(float64) // trying to assert a type without assertion value,
	                // Go wil panic
	fmt.Println(f)
}
