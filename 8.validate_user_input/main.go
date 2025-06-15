package main

import (
	"fmt"
	"strings"
)

func main() {

	conferenceName := "Go Conference" //go internally determines the datatype based on the assigned value
	const conferenceTickets = 50      //const val can't be changed. not allowed to change
	var remainingTickets = 50

	//print types
	fmt.Printf("conferenceName dataType: %T,  tickets dataType: %T\n", conferenceName, conferenceTickets)

	fmt.Printf("Welcome to our %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets to attned")

	bookings := []string{}

	//infinite loop
	for {
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

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidTickets := userTickets > 0 && userTickets <= remainingTickets
		isValidEmail := strings.Contains(email, "@")

		if isValidName && isValidTickets && isValidEmail {

			fmt.Printf("User %v %v booked %v tickets\n", firstName, lastName, userTickets)

			remainingTickets = remainingTickets - userTickets

			//slice methods
			bookings = append(bookings, firstName+" "+lastName)
			var fNames []string

			//for-each
			for _, booking := range bookings {
				names := strings.Fields(booking) // s="Prince Kumar" to ["Prince", "Kumar"] if whitespace is present
				fNames = append(fNames, names[0])

			}
			fmt.Printf("The first names of bookings are %v\n", fNames)

			// var noTicketsAvailable bool = remainingTickets == 0

			if remainingTickets == 0 {
				fmt.Println("Tickets are sold out. Please come next year")
				break
			} else if remainingTickets > 0 {
				fmt.Println("Tickets are remaining")
			} else {
				fmt.Println("Unknown tickets")
			}

			fmt.Printf("Out of %v tickets, % v is available now\n", conferenceTickets, remainingTickets)
			fmt.Printf("The whole slice %v\n", bookings)
			fmt.Printf("The first value %v\n", bookings[0])
			fmt.Printf("Slice Type %T\n", bookings)
			fmt.Printf("Slice Length %v\n", len(bookings))

		} else { //else shouldn't be in newline, should be immediately after closing brace } of if block

			if !isValidName {
				fmt.Println("firstName or lastName is too short")
			}
			if !isValidEmail {
				fmt.Println("email entered is invalid")
			}

			if !isValidTickets {
				fmt.Println("Number of tickets entered is invalid")
			}

		}
	}

}
