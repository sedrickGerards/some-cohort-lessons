package main

import "fmt"

func main() {

	// creatin a variable and it's always best to state it above

	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = conferenceTickets - 1

	fmt.Println("Welcome to the", conferenceName, "booking app")
	fmt.Println("Get your ticket here to attend")
	fmt.Println("we have a total of", conferenceTickets, "and we are left with", remainingTickets)

	// Go datatypes
	// when declaring or creating a variable, it must/ should be defined with a dtype

	var userName string
	var email string
	var userTickets int
// ask users for their name 

	fmt.Println("Welcome! could you please tell me your name");

	// we use pointers in the scan section to point to the register/ bin address to where the value is being stored(basiccaly  it assigns the value of the address it is being stored in to)
	fmt.Scanln(&userName)

	fmt.Printf("Please enter your email\n")
	fmt.Scanln(&email)
	
	fmt.Printf("how many tickets will you like to book  %v\n", userName)
	fmt.Scanln(&userTickets)

	fmt.Printf("thank You %v  you booked %v tickets for the %v,\n", userName, userTickets, conferenceName )

}
