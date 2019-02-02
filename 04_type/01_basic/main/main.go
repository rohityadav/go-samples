package main

import "fmt"

func main() {
	var x int

	fmt.Printf("%v \n%T\n", x, x)
	//x="String" you cant do that as go is a statically type language every varible is intialized to hold certail type of value, you cant dynamicaly assingn some differnet value.
}
