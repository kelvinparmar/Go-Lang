package helper

import "strings"

func UserInput(firstName string, LastName string, email string, ticketsToBuy int, remainingTickets int) (bool, bool, bool) {
	isvalidName := len(firstName) >= 2 && len(LastName) >= 2
	isvalidEmail := strings.Contains(email, "@")
	isvalidTicket := ticketsToBuy > 0 && ticketsToBuy <= remainingTickets
	return isvalidName, isvalidEmail, isvalidTicket
}


