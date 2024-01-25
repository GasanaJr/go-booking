package main

import (
	"fmt"
	"strings"
)

func main() {
	var confrenceName = "Go Conference"
	const confrenceTickets = 50
	var remainingTickets uint = 50

	fmt.Printf("Welcome to %v application\n", confrenceName)
	fmt.Println("We have a total of ",
		confrenceTickets, " tickets and ", remainingTickets,
		" tickets remaining")

	// alternative declaration
	bookings := []string{}
	//var bookings []string

	for {
		var userName string
		var lastName string
		var userTicket uint
		fmt.Println("Enter user name")
		fmt.Scan(&userName)
		fmt.Println("Enter last name")
		fmt.Scan(&lastName)
		fmt.Println("Enter number of tickets")
		fmt.Scan(&userTicket)

		if userTicket > remainingTickets {
			fmt.Println("Invalid number of tickets")
			continue
		}

		// Update tickets
		remainingTickets = remainingTickets - userTicket
		bookings = append(bookings, userName)
		fmt.Printf("Hello %v, you have %v tickets \n", userName, userTicket)
		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, confrenceName)

		//create username slice
		userNames := []string{}
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			userNames = append(userNames, names[0])
		}
		fmt.Printf("All the usernames who booked are %v \n", userNames)

		noTicketsRemaining := remainingTickets == 0
		if noTicketsRemaining {
			//end the program
			fmt.Println("The confrence tickets are SOLD OUT")
			break
		}
	}

}
