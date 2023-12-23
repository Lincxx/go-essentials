package main

import (
	"fmt"
)

func main() {
    var accountBalance = 1000.0

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

	//another way to set a bool
	//wantsCheckBalance := choice == 1

	//preform action based on the user selection
	if choice == 1{
		fmt.Println("Your balance is:", accountBalance)
	} else if choice == 2 {
		fmt.Println("Your deposit:")
		var depositAmount float64
		fmt.Scan(&depositAmount)
		accountBalance += depositAmount
		fmt.Println("New Balance:", accountBalance)
	} else if choice == 3 {
		fmt.Println("Withdrawal amount:")
		var withdrawalAmount float64
		fmt.Scan(&withdrawalAmount)
		accountBalance -= withdrawalAmount
		fmt.Println("New Balance:", accountBalance)
	} else {
		fmt.Println("Have a nice day")
	}
	
}