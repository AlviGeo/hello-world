package main

import "fmt"

func main() {
	// Define map
	// emails := make(map[string]string)

	// Assign Key Value
	// emails["Alvi"] = "alvigeovan@gmail.com"
	// emails["Mike"] = "mike@gmail.com"
	// emails["Sharron"] = "sharron@gmail.com"

	// Declare Map and Add Key Value
	emails := map[string]string{"Alvi":"alvigeovan@gmail.com", "Mike": "mike@gmail.com"}
	
	emails["Sharron"] = "sharron@gmail.com"

	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["Mike"])

	// Delete From Map
	delete(emails, "Mike")
	fmt.Println(emails)
}