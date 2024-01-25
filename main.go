package main

import "fmt"

func main() {
	var confrenceName = "Go Conference"
	const confrenceTickets = 50
	var remainingTickets = 50

	fmt.Printf("Welcome to %v application\n", confrenceName)
	fmt.Println("We have a total of ",
		confrenceTickets, " tickets and ", remainingTickets,
		" tickets remaining")

	var userName string
	var userTicket int
	fmt.Println("Enter user name")
	fmt.Scan(&userName)
	fmt.Println("Enter number of tickets")
	fmt.Scan(&userTicket)

	// Update tickets
	remainingTickets = remainingTickets - userTicket

	fmt.Printf("Hello %v, you have %v tickets \n", userName, userTicket)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, confrenceName)

}
