package main

import (
	"fmt"
	"strings"
)

var userName string

func main() {

	// creating a variable and it's always best to state it above

	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	remainingTickets := conferenceTickets

	greetUser(conferenceName, conferenceTickets)

	// creating booking array to store user data
	// var bookings [50]string

	//utilizing slice instead of arrays
	// Slices are better to use since it use when we have an infinite infinite/unknown input from the user and so we need not always state the index position before assigning a value to it but we just APPEND automatically to the next position
	var bookings []string

	// creating a for loop for an automated system
	for {

		// code to get out users first name
		firstName, secondName, email, userTickets := getUserInput()

		// handling user input errors

		isValidName, isValidEmail, isValidTicket := validataUserInput(firstName, secondName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicket {

			fmt.Printf("thank You %v  you booked %v tickets for the %v,\n", userName, userTickets, conferenceName)

			bookTicket( remainingTickets, userTickets, bookings)

			fmt.Printf("the whole slice: %v\n", bookings)
			fmt.Printf("the  slice length: %v\n", len(bookings))

			// code for first name selection
			// calling func to get user's first name

			firstNames := getFirstName(bookings)

			fmt.Printf("slice of all first names %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Printf("the %v is booked out, come again next session\n", conferenceName)
				break
			}
		} else {
			// we are not using else if because we want each of the codes to run and give corrections if the user made mistakes with all inputs
			if !isValidName {
				fmt.Println("Your first or last name must have more than 2 words")
			}
			if !isValidEmail {
				fmt.Println("the email you entered is not valid, it must contain '@'")
			}
			if !isValidTicket {
				fmt.Printf("sorry! we only have %v tickets left\n", remainingTickets)
			}
			continue
		}

	}

}

func greetUser(confName string, confTickets int) {

	fmt.Println("Welcome to the", confName, "booking app")
	fmt.Println("Get your ticket here to attend")
	fmt.Println("we have a total of", confTickets)
}

func getFirstName(bookings []string) []string {

	// doing the FOR EACh iteration of our array and after each iteration we receive two values which are the index and the value (we just choose booking as the value name to call it that )
	firstNameSlice := []string{}

	for _, booking := range bookings {

		var names = strings.Fields(booking)
		firstNameSlice = append(firstNameSlice, names[0])

	}

	return firstNameSlice

}

func validataUserInput(firstName string, secondName string, email string, userTickets int, remainingTickets int) (bool, bool, bool) {

	isValidName := len(firstName) > 2 && len(secondName) > 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicket := userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicket

}

// we don't need any input from the main here in the getUserINput we basically prompt the user to input
func getUserInput() (string, string, string, int) {

	// Go data types
	// when declaring or creating a variable, it must/ should be defined with a d type

	var firstName string
	var secondName string
	var email string
	var userTickets int

	fmt.Printf("Welcome! please enter your first name\n")

	// we use pointers in the scan section to point to the register/ bin address to where the value is being stored(basically  it assigns the value of the address it is being stored in to)
	fmt.Scanln(&firstName)

	fmt.Printf("Please enter your second name\n")
	fmt.Scanln(&secondName)

	fmt.Printf("Please enter your email name\n")
	fmt.Scanln(&email)

	userName = firstName + " " + secondName

	fmt.Printf("how many tickets will you like to book  %v\n", userName)
	fmt.Scanln(&userTickets)

	return firstName, secondName, email, userTickets

}

func bookTicket(remainingTickets int, userTickets int, bookings []string) ([]string) {

	remainingTickets = remainingTickets - userTickets
	fmt.Printf("we have %v ticket remaining\n", remainingTickets)

	bookings = append(bookings, userName)

	return  bookings
}
