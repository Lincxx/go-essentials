package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

//write to a file
func writeBalanceToFile(balance float64) {
	//convert the balance into a string of bytes
	balanceText := fmt.Sprint(balance)

	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}

func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(accountBalanceFile)

	//err will nil if we don't have an error
	if err != nil {
		return 1000, errors.New("Failed to find balance file")
	}

	//can only use string - float, int will not work with a byte array
	balanceText := string(data)
	//convert string to float
	balance, err := strconv.ParseFloat(balanceText, 64)
	
	if err != nil {
		return 1000, errors.New("Failed to parse stored balance value")
	}

	//we will return the balace and nil (meaning no error)
	return balance, nil
}


func main() {
    var accountBalance, err = getBalanceFromFile()

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Print("------")
	}

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
			writeBalanceToFile(accountBalance)
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
			writeBalanceToFile(accountBalance)
		default:
			fmt.Println("Have a nice day")
			fmt.Println("Thanks for choosing our bank")
			return
		}
	}
	
	
}