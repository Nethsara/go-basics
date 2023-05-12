package main

import (
	"fmt"
	"strings"
) 

func main(){
	conferenceName  := "Go Conferene"
	const conferenceTickets int = 50
	var remainingTickets  uint = 50
	var bookings []string
		
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v, We have available %v \n", conferenceTickets, remainingTickets)
	fmt.Println("Here you can start booking")

	for{

		var firstName string
		var lastName string
		var email string
		var userTickets uint
		
		fmt.Print("What is your first name : ")
		fmt.Scan(&firstName)

		fmt.Print("What is your last name : ")
		fmt.Scan(&lastName)

		fmt.Print("What is your email : ")
		fmt.Scan(&email)

		fmt.Print("Enter number of tickets : ")
		fmt.Scan(&userTickets)
		
		if userTickets > remainingTickets{
			fmt.Printf("We only have %v tickets remaining.", remainingTickets)
			continue
		}

		remainingTickets -= userTickets
		bookings = append(bookings, firstName+" " + lastName)

		fmt.Printf("Whole slice : %v\n", bookings)
		fmt.Printf("First value : %v\n", bookings[0])
		fmt.Printf("Slice type : %T\n", bookings)
		fmt.Printf("Length slice : %v\n", len(bookings))

		fmt.Printf("%v tickets remaining\n", remainingTickets)

		fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName,lastName, userTickets, email)
		
		firstNames := []string{}
		for _, booking := range bookings{
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}
		
		fmt.Printf("First Names of the bookings %v\n", firstNames)

		if remainingTickets == 0 {
			fmt.Println("Our conference if already booked up")
			break
		}

	}

}
