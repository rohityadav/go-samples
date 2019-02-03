package main

import "fmt"

const name string = "Rohit" // single constant intialization tyoe is optional it can be inferred
const (
	fname = "Rohit"
	lname = "yadav"
)
func main() {
	const innermethodconstant  = 12
	fmt.Println(name)
	fmt.Println(fname)
	fmt.Println(lname)
	fmt.Println(innermethodconstant)
}
