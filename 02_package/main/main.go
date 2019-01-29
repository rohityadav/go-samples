package main

import (
	"fmt"
	"go-samples/02_package/icomefromjaipur"
	"go-samples/02_package/stringutil"
)

func main() {
	fmt.Println(jaipur.FirstName)
	fmt.Println(jaipur.LastName())
	fmt.Println(stringutil.ReverseString("Rohit"))
	fmt.Println(stringutil.ReverseString("123456789"))
}
