package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/chethan0707/demo-go/helper"
)

var conferenceName string = "Go Conference"

const conferenceTickets = 50

var remainingTickets uint = 50

var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUser()
	for {
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber :=
			helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			bookTicket(userTickets, firstName, lastName, email)
			wg.Add(1)
			go sendTicket(userTickets, firstName, lastName, email)

			firstNames := getFirstNames()
			fmt.Printf("The first names of bookings: %v\n", firstNames)

			noTicketsRemaining := remainingTickets == 0
			if noTicketsRemaining {
				fmt.Println("Our Conference tickets are sold out!!!")
				break
			}
		} else {
			if !isValidEmail {
				fmt.Println("First name or Last name you entered is too short!!")
			}
			if !isValidEmail {
				fmt.Println("Email you entered does not have '@' sign!!")
			}
			if !isValidTicketNumber {
				fmt.Println("Ticket number you entered is invalid!!")
			}
		}
	}
	wg.Wait()
}

func greetUser() {
	fmt.Printf("Welcome to our %v booking application!\n", conferenceName)
	fmt.Printf("We have total of %d tickets and %v tickets are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets to attend")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	fmt.Print("Enter your first name: ")
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	fmt.Scan(&firstName)

	fmt.Print("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Print("Enter your email: ")
	fmt.Scan(&email)

	fmt.Print("Enter no. of tickets: ")
	fmt.Scan(&userTickets)
	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets -= userTickets
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets are remaining for %v\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("##################################")
	fmt.Printf("Sending ticket:\n  %v\n to  email address %v\n", ticket, email)
	fmt.Println("##################################")
	wg.Done()
}
