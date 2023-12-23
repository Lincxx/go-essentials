package main

import (
	"fmt"
)

func main() {
    var accountBalance = 1000.0

	fmt.Println("Welcome to Go Bank!!!")

	for {
		
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

		switch choice {
		case 1:
			fmt.Println("Your balance is:", accountBalance)
		case 2: 
			fmt.Println("Your deposit:")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater that 0")
				//return
				continue
			}

			accountBalance += depositAmount
			fmt.Println("New Balance:", accountBalance)
		case 3:
			fmt.Println("Withdrawal amount:")
			var withdrawalAmount float64
			fmt.Scan(&withdrawalAmount)

			if withdrawalAmount > accountBalance {
				fmt.Println("Invaild amount. You dont have that amount to withdraw")
				//return
				continue
			}

			if withdrawalAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater that 0")
				//return
				continue
			}

			accountBalance -= withdrawalAmount
			fmt.Println("New Balance:", accountBalance)
		default:
			fmt.Println("Have a nice day")
			fmt.Println("Thanks for choosing our bank")
			return
		}
	}
	
	
}