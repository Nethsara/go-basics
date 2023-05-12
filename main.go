package main

import "fmt" 

func main(){
	conferenceName  := "Go Conferene"
	const conferenceTickets int = 50
	var remainingTickets  uint = 50
		
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v, We have available %v \n", conferenceTickets, remainingTickets)
	fmt.Println("Here you can start booking")

	var firstName string
	var lastName string
	var email string
	var userTickets int
	fmt.Print("What is your first name : ")
	fmt.Scan(&firstName)

	fmt.Print("What is your last name : ")
	fmt.Scan(&lastName)

	fmt.Print("What is your email : ")
	fmt.Scan(&email)

	fmt.Print("Enter number of tickets : ")
	fmt.Scan(&userTickets)

	userTickets = 2
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName,lastName, userTickets, email)

}
