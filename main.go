package main

import "fmt" 

func main(){
	var conferenceName = "Go Conferene"
	const conferenceTickets = 50
	var remainingTickets  = conferenceTickets

	
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v We have available %v", conferenceTickets, remainingTickets)
	fmt.Println("Here you can start booking")

	

}
