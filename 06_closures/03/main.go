package main

import "fmt"

var x int = 0

func main() {
	//var increment = func () int {
	//	x++
	//	return x
	//}

	increment := func () int {
		x++
		return x
	}

	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment)
	fmt.Printf("%T\n", increment)
}



