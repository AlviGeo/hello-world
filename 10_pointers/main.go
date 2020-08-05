package main

import "fmt"

// & is pointer, used for memory address
// when * is showed in term, it means it's a pointer
func main() {
	a := 5
	b := &a

	fmt.Println(a, b)
	fmt.Printf("%T\n", b)

	// Use * to Read Value from Address
	fmt.Println(*&a)

	// Change Value with Pointer
	*b = 10
	fmt.Println(a)
}