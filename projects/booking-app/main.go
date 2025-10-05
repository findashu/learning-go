package main

import "fmt"

func main() {

	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50

	fmt.Printf("Welcome to %v Booking Application.\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)

	//var bookings = [50]string{} // empty array of strings with a capacity of 50

	for {
		var bookings []string // Slice is a dynamically sized array. Array doesn't have append method but slice has append method

		var userName string
		var userTickets uint
		// var userTickets int

		fmt.Println("Enter your name:")

		// reading from standard input
		// Have to use & (Pointer) to get the memory address of the variable
		// Basically we are passing the memory address of the variable to the Scan function
		// &userName means we are passing the memory address of userName variable to the Scan function
		// so that Scan function can store the input value at that memory address
		// If we don't use & then Scan function will not be able to store the input value in userName variable
		// and userName variable will remain empty
		fmt.Scan(&userName)

		// Pointer is a variable that stores the memory address of another variable, memory address of any variable can be obtained using & operator
		// fmt.Printf("Address of variable userName is: %v\n", &userName)

		fmt.Printf("Enter number of tickets:")
		fmt.Scan(&userTickets)

		remainingTickets = remainingTickets - userTickets

		bookings = append(bookings, userName) // appending userName to the bookings slice

		fmt.Printf("The whole slice: %v\n", bookings)
		// fmt.Printf("The first value: %v\n", bookings[0])
		// fmt.Printf("Slice type: %T\n", bookings)
		// fmt.Printf("Slice length: %v\n", len(bookings))

		fmt.Printf("Thank you %v for booking %v tickets. You will receive a confirmation email shortly.\n", userName, userTickets)
		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
	}
}
