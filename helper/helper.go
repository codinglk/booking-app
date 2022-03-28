package helper

import "strings"

// A go function can return multiple values, Capitalize the first letter of function if you create in other packages, which
// indicates this function will be used in other packages
func ValidateUserInput(firstName string, lastName string, userTickets uint, email string, remainingTickets uint) (bool, bool, bool) {
	// var isValidName bool = len(firstName) >=2 && len(lastName) >=2
	isValidName := len(firstName) >= 2 && len(lastName) >= 2 // It works same as above line
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber
}
