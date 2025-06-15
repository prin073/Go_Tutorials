package main

import (
	"fmt"
	"multiple_packages/helper"
	"strings"
)

// package level variables(defined outside all functions)
// vars here can't be defined using :=
const conferenceTickets = 50         //const val can't be changed. not allowed to change
var conferenceName = "Go Conference" //go internally determines the datatype based on the assigned value
var remainingTickets uint = 50
var bookings = []string{}

func main() {

	//print types
	fmt.Printf("conferenceName dataType: %T,  tickets dataType: %T\n", conferenceName, conferenceTickets)

	greetUsers(conferenceTickets)

	//infinite loop
	for {
		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTickets := helper.ValidateUserInputs(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidTickets && isValidEmail {

			fmt.Printf("User %v %v booked %v tickets\n", firstName, lastName, userTickets)

			bookTickets(userTickets, firstName, lastName)

			fNames := getFirstNames()
			fmt.Printf("The first names of bookings are %v\n", fNames)

			// var noTicketsAvailable bool = remainingTickets == 0

			if remainingTickets == 0 {
				fmt.Println("Tickets are sold out. Please come next year")
				break
			} else if remainingTickets > 0 {
				fmt.Println("Tickets are currently available")
			} else {
				fmt.Println("Unknown ticket count")
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

func greetUsers(confTickets int) {
	fmt.Printf("Welcome to our %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are are still available\n", confTickets, remainingTickets)
	fmt.Println("Get your tickets to attend")
}

// return a val with return types
func getFirstNames() []string {
	fNames := []string{}
	//for-each
	for _, booking := range bookings {
		names := strings.Fields(booking) // s="Prince Kumar" to ["Prince", "Kumar"] if whitespace is present
		fNames = append(fNames, names[0])

	}
	return fNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var userTickets uint
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

	return firstName, lastName, email, userTickets
}

func bookTickets(userTickets uint, firstName string, lastName string) (uint, []string) {

	remainingTickets = remainingTickets - userTickets

	//slice methods
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email\n", firstName, lastName, userTickets)
	fmt.Printf("%v tickets remaining for %v\n", userTickets, conferenceName)
	return remainingTickets, bookings
}
