package main

import "fmt"

func main() {

	var conferenceName = "Go Conference" //go internally determines the datatype based on the assigned value
	//this works
	// var conferenceName = "Go Conference" //go internally determines the datatype based on the assigned value

	//another way of var declaring (shorthand)
	conferenceRoom := "Stranager Things" //same as var conferenceRoom := "Stranager Things"

	const conferenceTickets = 50 //const val can't be changed. not allowed to change
	var remainingTickets = 50

	fmt.Println("Welcome to our", conferenceName, "booking application")
	fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "are still available")

	//another way of printing
	fmt.Printf("Welcome to our %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are are still available\n", conferenceTickets, remainingTickets)

	fmt.Println("Get your tickets to attned")

	/*
		below won't work here. If you use var then you should declare and assign atonce and later variables can be reassigned with other values
		var userName
		userName = "Tom"
		fmt.Print(userName)
	*/

	fmt.Printf("Conference Room is %v", conferenceRoom)

}
