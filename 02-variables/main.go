package main

import "fmt"

func main() {
	// Declare a variable with an initial value and var keyword
	// The type is inferred from the value Dynamic Intialization
	// If you declare a variable and do not use it, Go will throw a compile-time error.
	// This helps in maintaining cleaner code.
	var message = "Hello, Go!"

	// var lastName this will throw an error because it is declared but no value assigned
	//lastName = "Doe" // Dynamic Initialization

	// Declare a constant
	const pi = 3.14 // Constant declaration

	var lastName string // Static Initialization
	lastName = "Doe"    // Assigning value later

	middleName := "A." // Short variable declaration (only inside functions)

	// You can also declare multiple variables in a single line
	var a, b, c int = 1, 2, 3 // Multiple variable declaration with initialization

	// You can also use the short declaration for multiple variables
	x, y := 10, 20 // Short variable declaration for multiple variables

	// Print the multiple variables to avoid unused variable error

	// Print the values
	fmt.Println("Pi:", pi)

	fmt.Println("Message:", message)
	fmt.Println("Last Name:", lastName)

	fmt.Println("Middle Name:", middleName)
	fmt.Println(a, b, c, x, y)

}
