package main

import "fmt"

func main() {
	sum := 0
// syntax	    [init]  [cond] [post]
			for i := 0; i < 10; i++ {

				sum += i
			}
	fmt.Println(sum)

// init and post are optional? hmm
	sum = 1
	for ; sum < 1000; {

		sum += sum
	}
	fmt.Println(sum)


	// for is Go's while
	sum = 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)


	// forever (lol)
	// for { }
}
