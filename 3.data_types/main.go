package main

import "fmt"

func main() {

	//this works
	var conferenceName = "Go Conference" //go internally determines the datatype based on the assigned value

	//another way of var declaring (shorthand)
	conferenceRoom := "Stranager Things" //same as var conferenceRoom := "Stranager Things"

	const conferenceTickets = 50 //const val can't be changed. not allowed to change. Constnats can't be declared like :=
	var remainingTickets = 50

	//print types
	fmt.Printf("conferenceName dataType: %T,  tickets dataType: %T\n", conferenceName, conferenceTickets)

	fmt.Printf("Welcome to our %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are are still available\n", conferenceTickets, remainingTickets)

	fmt.Println("Get your tickets to attned")

	/*
		below won't work here. If you use var then you should declare and assign atonce and later variables can be reassigned with other values
		var userName
		userName = "Tom"
		fmt.Print(userName)
	*/

	var userName string
	var userTickets int
	var email string

	/* this works
	userName = "Prince"
	userTickets = 2
	*/

	//ask user for values. use Scan()
	fmt.Println("Enter UserName:")
	fmt.Scan(&userName) //pointers, passes address of userName and when user enters value, assigns value to it

	fmt.Println("Enter tickets:")
	fmt.Scan(&userTickets)

	fmt.Println("Enter email:")
	fmt.Scan(&email)

	fmt.Printf("User %v booked %v tickets", userName, userTickets)

	remainingTickets = remainingTickets - userTickets

	fmt.Printf("Out of %v tickets, % v is available now\n", conferenceTickets, remainingTickets)
	fmt.Printf("Conference Room is %v", conferenceRoom)

}
