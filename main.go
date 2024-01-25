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
}
