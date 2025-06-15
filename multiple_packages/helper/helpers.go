package helper

import "strings"

var MyVar = "GlobalVar" //can be exported as a global var. Remember global vars should also start with upper case

// multiple returns with return types
// fun name starting with lower case letter is treated as private function
// so to export the func, start func name with upper case
func ValidateUserInputs(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidTickets := userTickets > 0 && userTickets <= remainingTickets
	isValidEmail := strings.Contains(email, "@")
	return isValidName, isValidEmail, isValidTickets
}
