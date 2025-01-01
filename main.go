package main

import "fmt"

func main() {

	var storeName = "TREND TROVE STORE"
	const storeTickets = 50
	var remainingTickets = 50
	fmt.Printf("Welcome to %v our store \n", storeName)
	fmt.Printf("We have a total of %v tickets and %v are still available \n", storeTickets, remainingTickets)
	fmt.Println("Get your ticket here")

	var userName string
	var userTickets uint

	fmt.Println("Enter your name: ")
	fmt.Scan(&userName)
	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

}
