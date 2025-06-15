package main

import "fmt"

func main() {

	//1. print
	fmt.Print("Hello World!")
	fmt.Println("How are you!")

	var userName = "Prince"                                               //go internally determines the datatype based on the assigned value
	const tickets = 50                                                    //const val can't be changed
	fmt.Println("My name is", userName, "and I have", tickets, "tickets") //while printing in console comma will be replace3d with a space
	fmt.Printf("My name is %v and I have %v tickets", userName, tickets)

}
