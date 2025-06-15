package main

import (
	"fmt"
	"sync"
	"time"
)

// package level variables(defined outside all functions)
// vars here can't be defined using :=
const conferenceTickets = 50         //const val can't be changed. not allowed to change
var conferenceName = "Go Conference" //go internally determines the datatype based on the assigned value
var remainingTickets uint = 50

// waitGroup for sync goroutines
// sync package provides basic synchronization functionality
var wg = sync.WaitGroup{}

// struct(kind of lightweight class)
type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var bookings = make([]UserData, 0) //empty list of UserData struct of map

func main() {

	//print types
	fmt.Printf("conferenceName dataType: %T,  tickets dataType: %T\n", conferenceName, conferenceTickets)

	greetUsers(conferenceTickets)

	firstName, lastName, email, userTickets := getUserInput()

	isValidName, isValidEmail, isValidTickets := validateUserInputs(firstName, lastName, email, userTickets)

	if isValidName && isValidTickets && isValidEmail {

		fmt.Printf("User %v %v booked %v tickets\n", firstName, lastName, userTickets)

		bookTickets(userTickets, firstName, lastName, email)

		//sets the number of goroutines to wait for (increases the counter by the provided number)
		wg.Add(1) //we are just creating 1 separate thread other than main thread

		//takes some time to send the ticket. so to avoid obstruction for main thread execution
		//use concurrency and run sendTickets in separate thread by using go keyword
		go sendTickets(userTickets, firstName, lastName, email)

		fNames := getFirstNames()
		fmt.Printf("The first names of bookings are %v\n", fNames)

		// var noTicketsAvailable bool = remainingTickets == 0

		if remainingTickets == 0 {
			fmt.Println("Tickets are sold out. Please come next year")
		} else if remainingTickets > 0 {
			fmt.Println("Tickets are currently available")
		} else {
			fmt.Println("Unknown ticket count")
		}

		fmt.Printf("Out of %v tickets, % v is available now\n", conferenceTickets, remainingTickets)
		fmt.Printf("The whole struct %v\n", bookings)
		fmt.Printf("The first value %v\n", bookings[0])
		fmt.Printf("Struct Type %T\n", bookings)
		fmt.Printf("Struct Length %v\n", len(bookings))

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
	//Blocks until the WG counter is 0
	//waits for all the thread being done which added using Add(). Blocks the main thread from terminating the overall execution
	// if other threads are still in progress
	wg.Wait()
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
		names := booking.firstName //using from struct
		fNames = append(fNames, names)

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

func bookTickets(userTickets uint, firstName string, lastName string, email string) (uint, []UserData) {

	remainingTickets = remainingTickets - userTickets

	//create struct
	var user = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	//slice methods
	bookings = append(bookings, user)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email on %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", userTickets, conferenceName)
	return remainingTickets, bookings
}

func sendTickets(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	fmt.Println("#####################")
	var message = fmt.Sprintf(" %v tickets are sent to %v %v ", userTickets, firstName, lastName)
	fmt.Printf("Sending ticket\n to:%v\n%v\n", email, message)
	fmt.Println("#####################")

	//Decrements the WaitGroup counter by 1.
	//So this is called by the goroutines to indicate that it's finished
	wg.Done()
}
