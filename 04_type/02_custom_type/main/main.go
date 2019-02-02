package main

import "fmt"

var a int

type hotdog int //Custom types
var b hotdog

func main() {
	a = 42
	b = 43
	c := 23.23
	fmt.Printf("%v\n%T\n%v\n%T\n", a, a, b, b)
	fmt.Println(int(b))
	fmt.Println(int(c))
	fmt.Printf("%T\n", float64(a))
	fmt.Printf("%T\n", int(b))
}