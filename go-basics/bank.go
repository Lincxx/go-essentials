package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to Go Bank!!!")
	fmt.Println("What would you like to do?")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")

	//get the user input
	var choice int
	fmt.Print("Your choice: ")
	fmt.Scan(&choice)

	//preform action based on the user selection

}