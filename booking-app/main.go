package main

import (
	"booking-app/helper"
	"fmt"
	"time"
)

var conferenceName = "Go conference"
const conferenceTickets = 100
var remainingTickets = conferenceTickets
var bookings = make([]UserData,0)

type UserData struct {
	firstName string
	LastName string
	email string
	ticketsToBuy int

}


func main() {

	greetUsers()

	for {

		firstName, LastName, email, ticketsToBuy := getUserInput()

		isvalidName, isvalidEmail, isvalidTicket := helper.UserInput(firstName, LastName, email, ticketsToBuy, remainingTickets)

		if isvalidEmail && isvalidName && isvalidTicket {

			bookTicket(ticketsToBuy, firstName, LastName, email)
			go sendTicket(ticketsToBuy, firstName, LastName, email)

			fmt.Printf("The first names of our bookings are: %v\n", firstName)

			firstNames := getFirstNames()
			fmt.Printf("The First names of our bookings are: %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("We are sold out, try next day")
				break
			}
		} else {
			if !isvalidName {
				fmt.Printf("FirstName or LastName you entered is too short\n")
			}
			if !isvalidEmail {
				fmt.Printf("email address you entered dosen't contain @ sign\n")
			}
			if !isvalidTicket {
				fmt.Printf("Please enter valid tickets to buy\n")
			}
		}
	}

}

func greetUsers() {

	fmt.Printf("Welcome to %v booking app\n", conferenceName)
	fmt.Printf("conferenceName is: %T, conferenceTickets is: %T, remainingTickets is: %T\n", conferenceName, conferenceTickets, remainingTickets)

}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, int) {
	var firstName string
	var LastName string
	var Email string
	var ticketsToBuy int

	fmt.Print("Please enter your first name: ")
	fmt.Scanln(&firstName)

	fmt.Print("Please enter your last name: ")
	fmt.Scanln(&LastName)

	fmt.Print("Please enter your email: ")
	fmt.Scanln(&Email)

	fmt.Print("Please enter the number of tickets you want to buy: ")
	fmt.Scanln(&ticketsToBuy)

	return firstName, LastName, Email, ticketsToBuy
}

func bookTicket(ticketsToBuy int, firstName string, LastName string, Email string) {
	remainingTickets = remainingTickets - ticketsToBuy
//create a map for a user


	var userData = UserData{
		firstName : firstName,
		LastName : LastName,
		email : Email,
		ticketsToBuy :	ticketsToBuy,
	}


	bookings = append(bookings, userData)
	fmt.Printf("List of bookings are : %v\n", bookings)

	fmt.Printf("You have bought %v tickets\n", ticketsToBuy)
	fmt.Printf("There are %v tickets left\n", remainingTickets)

}

func sendTicket(ticketsToBuy int, firstName string, LastName string, Email string) {
	time.Sleep(15 * time.Second)
	var ticket = fmt.Sprintf("%v, tickets for  %v %v ", ticketsToBuy, firstName, LastName)
	fmt.Println("\n -------------------------------------")
	fmt.Printf("Sending ticket : %v \n to email address : %v\n", Email, ticket)
	fmt.Println("-------------------------------------")
}
