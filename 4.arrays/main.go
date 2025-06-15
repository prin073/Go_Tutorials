package main

import "fmt"

func main() {

	var conferenceName = "Go Conference" //go internally determines the datatype based on the assigned value
	const conferenceTickets = 50         //const val can't be changed. not allowed to change
	var remainingTickets = 50

	//print types
	fmt.Printf("conferenceName dataType: %T,  tickets dataType: %T\n", conferenceName, conferenceTickets)

	fmt.Printf("Welcome to our %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets to attned")

	// arrays.

	//1.works
	// var bookings [50]string

	//2.works. var <var_name> = [size]dataType{"val1", "val2"}
	// var bookings = [50]string{}

	//3.works
	bookings := [50]string{}

	var firstName string
	var lastName string
	var userTickets int
	var email string

	//ask user for values. use Scan()
	fmt.Println("Enter firstName:")
	fmt.Scan(&firstName)

	fmt.Println("Enter lastName:")
	fmt.Scan(&lastName)

	fmt.Println("Enter tickets:")
	fmt.Scan(&userTickets)

	fmt.Println("Enter email:")
	fmt.Scan(&email)

	fmt.Printf("User %v %v booked %v tickets\n", firstName, lastName, userTickets)

	remainingTickets = remainingTickets - userTickets
	bookings[0] = firstName + "" + lastName

	fmt.Printf("Out of %v tickets, % v is available now\n", conferenceTickets, remainingTickets)
	fmt.Printf("The whole array %v\n", bookings)
	fmt.Printf("The first value %v\n", bookings[0])
	fmt.Printf("Array Type %T\n", bookings)
	fmt.Printf("Array Length %v\n", len(bookings))

}
