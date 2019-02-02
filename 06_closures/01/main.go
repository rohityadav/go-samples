package main

import "fmt"

var x int
func main() {
	x = 0

	{
		fmt.Println(x)
		y:="test"
		fmt.Println(y)
	}
	//fmt.Println(y) not going to work outside scope of the closure
}
