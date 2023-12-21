package main

import "fmt"

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	// fmt.Print("What is your revenue: ")
	// fmt.Scan(&revenue)
	revenue = getUserInput("What is your revenue: ")

	// fmt.Print("What is your expenses: ")
	// fmt.Scan(&expenses)
	expenses = getUserInput("What is your expenses: ")

	// fmt.Print("What is your tax rate: ")
	// fmt.Scan(&taxRate)
	taxRate = getUserInput("What is your tax rate: ")

	// ebt := revenue - expenses
	// profit := ebt * (1-taxRate / 100)
	// ratio := ebt / profit

	// fmt.Println(ebt)
	// fmt.Println(profit)
	// fmt.Println(ratio)
	profitOutput(revenue, expenses, taxRate)
}

func getUserInput(infoText string ) float64 {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)

	return userInput
}

func profitOutput(revenue float64, expenses float64, taxRate float64) {
	ebt := revenue - expenses
	profit := ebt * (1-taxRate / 100)
	ratio := ebt / profit

	fmt.Println(ebt)
	fmt.Println(profit)
	fmt.Println(ratio)

}