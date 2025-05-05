package main

import "fmt"

func main() {

	// creatin a variable and it's always best to state it above

	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	 remainingTickets := conferenceTickets

	fmt.Println("Welcome to the", conferenceName, "booking app")
	fmt.Println("Get your ticket here to attend")
	fmt.Println("we have a total of", conferenceTickets)

	// Go datatypes
	// when declaring or creating a variable, it must/ should be defined with a dtype

	var userName string
	var email string
	var userTickets int

	// creating booking array to store user data
	// var bookings [50]string

	//utilising slice instead of arrrays
	// Slices are better to use since it use when we have an infinite infinite/unknoen input from the user and so we neeed not always state the index position before asigning a value to it but we just APPEND automatically to the next position
	var bookings[]string;


	// creating a for loop for an automated system
for {
	
	// ask users for their name 

	fmt.Println("Welcome! could you please tell me your name");

	// we use pointers in the scan section to point to the register/ bin address to where the value is being stored(basicaly  it assigns the value of the address it is being stored in to)
	fmt.Scanln(&userName)

	fmt.Printf("Please enter your email\n")
	fmt.Scanln(&email)
	
	fmt.Printf("how many tickets will you like to book  %v\n", userName)
	fmt.Scanln(&userTickets)

	fmt.Printf("thank You %v  you booked %v tickets for the %v,\n", userName, userTickets, conferenceName )

	remainingTickets = remainingTickets - userTickets
	fmt.Printf("we have %v ticket remaining\n", remainingTickets)
	
	bookings = append(bookings, userName)
	fmt.Printf("the whole slice: %v\n", bookings)
	fmt.Printf("the first value: %v\n", bookings[0])
	fmt.Printf("the slice type: %T\n", bookings)
	fmt.Printf("the  slice length: %v\n", len(bookings))
	

}

}
