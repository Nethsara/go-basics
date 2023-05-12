package main

import "fmt" 

func main(){
	conferenceName  := "Go Conferene"
	const conferenceTickets int = 50
	var remainingTickets  uint = 50
	

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T\n", conferenceTickets, remainingTickets)
	
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v, We have available %v \n", conferenceTickets, remainingTickets)
	fmt.Println("Here you can start booking")

	var userName string
	var userTickets int

	userName = "Tom"
	userTickets = 2
	fmt.Printf("User %v booked %v tickets\n", userName, userTickets)

}
