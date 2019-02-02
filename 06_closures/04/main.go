package main

import "fmt"

var x int

func wrapper() func() int {
	x=10
	return func() int {
		x++
		return x
	}
}

func main() {
	increment := wrapper();

	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment)
	fmt.Printf("%T\n", increment)
}



