package main

import "fmt"

const (
	A = iota
	B = iota
	C = iota
)

const (
	d = iota
	e
	f
)

const (
	g = iota
	h = iota * 10
	i = iota * 10
)

const (
	j = 1 << 1 // bit shifting operation
	k = 1 << 2
	l = 1 << 3
)

const (
	_ = iota
	KB = 1 << (iota*10)
	MB = 1 << (iota*10)
	GB = 1 << (iota*10)
)


func main() {
	fmt.Printf("%v\n%T\n", A, A)
	fmt.Printf("%v\n%T\n", B, B)
	fmt.Printf("%v\n%T\n", C, C)

	fmt.Printf("%v\n%T\n", d, d)
	fmt.Printf("%v\n%T\n", e, e)
	fmt.Printf("%v\n%T\n", f, f)

	fmt.Printf("%v\n%T\n", g, g)
	fmt.Printf("%v\n%T\n", h, h)
	fmt.Printf("%v\n%T\n", i, i)

	fmt.Printf("%b\n%v\n%T\n", j, j, j)
	fmt.Printf("%b\n%v\n%T\n", k, k, k)
	fmt.Printf("%b\n%v\n%T\n", l, l, l)

	fmt.Printf("%b\n%v\n%T\n", KB, KB, KB)
	fmt.Printf("%b\n%v\n%T\n", MB, MB, MB)
	fmt.Printf("%b\n%v\n%T\n", GB, GB, GB)
}
