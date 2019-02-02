package main

import (
	"fmt"
	"go-samples/03_variables/utils"
)


var test_package_scope = "Rohit"

func main() {
	shorthand()
	zero_default_intialization()
	declare_intialize_variable()
	declare_intialize_many_at_once()
	declare_many_with_infer_type()
	fmt.Println(test_package_scope)
	utils.Test2()

}

func declare_many_with_infer_type() {
	var a, b, c = 1, "String", 3

	fmt.Println(a, b, c)
}

func declare_intialize_many_at_once() {
	var a, b, c int = 1, 2, 3

	fmt.Println(a, b, c)
}

func declare_intialize_variable() {
	var a int = 12
	fmt.Println("\n\n Declare variables")
	fmt.Println(a)
}

func zero_default_intialization() {
	var a  string
	var b int
	var c bool
	var d float64
	var e int8

	fmt.Printf("%v \n", a)
	fmt.Printf("%v \n", b)
	fmt.Printf("%v \n", c)
	fmt.Printf("%v \n", d)
	fmt.Printf("%v \n", e)

	fmt.Printf("%T \n", a)
	fmt.Printf("%T \n", b)
	fmt.Printf("%T \n", c)
	fmt.Printf("%T \n", d)
	fmt.Printf("%T \n", e)

}

func shorthand()  {
	a := "String"
	b := 12
	c := true
	d := `this is single quote "String"`
	e := `M`

	fmt.Printf("%v \n", a)
	fmt.Printf("%v \n", b)
	fmt.Printf("%v \n", c)
	fmt.Printf("%v \n", d)
	fmt.Printf("%v \n", e)

	fmt.Printf("%T \n", a)
	fmt.Printf("%T \n", b)
	fmt.Printf("%T \n", c)
	fmt.Printf("%T \n", d)
	fmt.Printf("%T \n", e)

}
