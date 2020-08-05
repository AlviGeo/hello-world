package main

import "fmt"

func main() {
	// using var
	// var name = "Alvi"
	var age int32 = 19
	const isCool = true
	var size float32 = 2.3

	// Shorthand
	// name := "Alvi"
	// email := "alvigeovan29@gmail.com"

	name, email := "Alvi", "alvigeovan29@gmail.com"


	fmt.Println(name, age, isCool, email)
	fmt.Printf("%T\n", size)
}
