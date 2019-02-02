package main

import (
	"fmt"
	"go-samples/05_scope/01_package_scope/util"
)

func main() {
	util.PrintName()
	fmt.Println(util.Name)
}
