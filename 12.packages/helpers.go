package main

import "strings"

// multiple returns with return types
func validateUserInputs(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidTickets := userTickets > 0 && userTickets <= remainingTickets
	isValidEmail := strings.Contains(email, "@")
	return isValidName, isValidEmail, isValidTickets
}
