package main

// importing the fmt package for i/o
import "fmt"

// main method for the project
func main() {
	// creating a variable of type string using type inference to store the conference name
	conferenceName := "Go Conference"

	// creating a constant for getting the number of tickets
	const conferenceTickets = 50

	// creating a variable for getting the number of remaining tickets

	// the variable type of remainingTickets is uint since remainingTickets cannot be negative
	var remainingTickets uint = 50 // initially all the tickets are remaining

	// using the Printf() method from the fmt package for printing the formatted output
	fmt.Printf("Welcome to %v booking application\n", conferenceName) // here %v is a placeholder for conferenceName
	// and \n is used for new line

	// using Println() method for printing welcome message for users using string interpolation
	fmt.Println("We have", conferenceTickets, "tickets and", remainingTickets, "tickets are still available")
	fmt.Println("Get your tickets here to attend")

	// variable to get the user name
	var firstName string
	// variable to get lastName of user
	var lastName string
	// variable to get email address of user
	var email string
	// variable to get the number of tickets booked by the user
	var userTickets int

	fmt.Println("Enter your first name:")
	// getting the name of user as input using the Scan() method and the & keyword
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name:")
	// getting the name of user as input using the Scan() method and the & keyword
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address:")
	// getting the name of user as input using the Scan() method and the & keyword
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	// getting the name of user as input using the Scan() method and the & keyword
	fmt.Scan(&userTickets)

	// thank you message for the user
	fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation email at %v\n", firstName, lastName, userTickets, email)

	// initializing the value of userName and userTickets below
	// userName = "Tom"
	// userTickets = 2

	// fmt.Printf("User %v booked %v tickets \n", userName, userTickets)
}
