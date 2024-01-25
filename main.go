package main

import (
	"fmt"
	"strings"
)

func main() {
	var remainingTickets uint = 50
	var confrenceName = "Go Confrence"
	// alternative declaration
	bookings := []string{}
	//var bookings []string
	greetUsers(confrenceName)

	for {
		// Get user data
		userName, lastName, userTicket, email := getUserInputs()
		// User validation
		isValidEmail, isValidName, isValidTicket := validateUserInputs(userName, lastName, email, userTicket, remainingTickets)

		if isValidEmail && isValidTicket && isValidName {
			// Update tickets
			remainingTickets = remainingTickets - userTicket
			bookings = append(bookings, userName+" "+lastName)
			fmt.Printf("Hello %v, you have %v tickets \n", userName, userTicket)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, confrenceName)

			//call print first names
			var firstNames = printFirstNames(bookings)
			fmt.Printf("All the usernames who booked are %v \n", firstNames)

			noTicketsRemaining := remainingTickets == 0
			if noTicketsRemaining {
				//end the program
				fmt.Println("The confrence tickets are SOLD OUT")
				break
			}
		} else {
			if !isValidEmail {
				fmt.Println("Please enter a valid email address")
			}
			if !isValidName {
				fmt.Println("Please enter a valid name")
			}
			if !isValidTicket {
				fmt.Println("Please enter a valid number of tickets")
			}
			continue
		}

	}

}

func greetUsers(confName string) {
	const confrenceTickets = 50
	fmt.Printf("Welcome to %v application\n", confName)
	fmt.Println("We have a total of ",
		confrenceTickets, " tickets")
}

func printFirstNames(bookings []string) []string {
	userNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		userNames = append(userNames, names[0])
	}
	return userNames
}

func validateUserInputs(userName string, lastName string, email string, userTicket uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(userName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicket := userTicket > 0 && userTicket <= remainingTickets

	return isValidEmail, isValidName, isValidTicket
}

func getUserInputs() (string, string, uint, string) {
	var userName string
	var lastName string
	var userTicket uint
	var email string
	fmt.Println("Enter user name")
	fmt.Scan(&userName)
	fmt.Println("Enter last name")
	fmt.Scan(&lastName)
	fmt.Println("Enter email")
	fmt.Scan(&email)
	fmt.Println("Enter number of tickets")
	fmt.Scan(&userTicket)
	return userName, lastName, userTicket, email

}
