package main

import "fmt" 

func main(){
	var conferenceName = "Go Conferene"
	const conferenceTickets = 50
	var remainingTickets  = conferenceTickets

	
	fmt.Println("Welcome to",conferenceName, "booking application")
	fmt.Println("We have total of", conferenceTickets, "We have available", remainingTickets)
	fmt.Println("Here you can start booking")

	

}
