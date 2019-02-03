package main

import "fmt"

func main() {
	a := 10
	fmt.Println(a) // this is the value of the variables.
	fmt.Println(&a) // this is the memoery addressese where the value actually stored.

	var b *int = &a
	fmt.Println(b) // this is the same memory addresse where value of a is stored.
	fmt.Println(&b) // this is the memory address where value of b is stored and what is the value of b..it is actually holding memory address of a.
	fmt.Println(*b) // this is the dereferencing operation which can invoke on pointer type, and it derefernce the value stored in the pointer which in this case is a.
	*b = 45
	// now i changed the value of the object whose memory addresse b is holding, so implicitly a is also changed.
	fmt.Println(a)
	fmt.Println(&a, b) // now if you print the value of memory location it is still going to be same becuase we didnt swap the memory addresses we directly put the value at the memory location.

	c := 20
	b = &c
	fmt.Println(&c, b) // so you can also change the value of b by swapping memory addresses.
	fmt.Println(*b) // now if you dereference b you will see different value as memeory addresses are already changed.
	c = 67
	fmt.Println(*b) // as b and c point to same memory addresses so deferencing b will give you latest value of c.




}
