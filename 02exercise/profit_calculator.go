package main

import (
	"fmt"
	"os"
)

const profitInformationFile = "profitInformation.txt"

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

func getUserInput(infoText string) float64 {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)

	if userInput <= 0 {
		fmt.Println("Invalid amount. Must be greater that 0")
		return 0
	}

	return userInput
}

func profitOutput(revenue float64, expenses float64, taxRate float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit

	fmt.Println("EBT: ", ebt)
	fmt.Println("PRofit: ", profit)
	fmt.Println("Ratio:", ratio)

}

// write to a file
func writeInfoToFile(ebt float64, profit float64, taxRate float64) {
	//convert the balance into a string of bytes
	ebtText := fmt.Sprint(ebt)
	profitText := fmt.Sprint(profit)
	taxRateText := fmt.Sprint(taxRate)

	profitTextInfo := fmt.Sprint("EBT Tax:", ebtText, "\nProfit:", profitText, "\nTax Rate", taxRateText)
	fmt.Println("Writing to file....")
	os.WriteFile(profitInformationFile, []byte(profitTextInfo), 0644)
}
