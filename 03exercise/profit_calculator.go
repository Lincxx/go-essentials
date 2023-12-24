package main

import (
	"errors"
	"fmt"
	"os"
)

const profitInformationFile = "profitInformation.txt"

func main() {
	//var revenue float64
	//var expenses float64
	//var taxRate float64

	// fmt.Print("What is your revenue: ")
	// fmt.Scan(&revenue)
	revenue, revErr := getUserInput("What is your revenue: ")

	if revErr != nil {
		fmt.Println(revErr)
		return
	}

	// fmt.Print("What is your expenses: ")
	// fmt.Scan(&expenses)
	expenses, expErr := getUserInput("What is your expenses: ")

	if expErr != nil {
		fmt.Println(expErr)
		return
	}

	// fmt.Print("What is your tax rate: ")
	// fmt.Scan(&taxRate)
	taxRate, taxErr := getUserInput("What is your tax rate: ")

	if taxErr != nil {
		fmt.Println(taxErr)
		return
	}

	// ebt := revenue - expenses
	// profit := ebt * (1-taxRate / 100)
	// ratio := ebt / profit

	// fmt.Println(ebt)
	// fmt.Println(profit)
	// fmt.Println(ratio)

	//can have a single err check
	//if revErr != nil || expErr != nil || taxErr != nil {
	//	fmt.Println(revErr, expErr, taxErr)
	//	return
	//}

	profitOutput(revenue, expenses, taxRate)
}

func getUserInput(infoText string) (float64, error) {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)

	if userInput <= 0 {
		return 0, errors.New("Invalid amount. Must be greater that 0")
	}

	return userInput, nil
}

func profitOutput(revenue float64, expenses float64, taxRate float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit

	fmt.Println(ebt)
	fmt.Println(profit)
	fmt.Println(ratio)
	//write to a file
	writeInfoToFile(ebt, profit, ratio)

}

// write to a file
func writeInfoToFile(ebt float64, profit float64, taxRate float64) {
	//convert the balance into a string of bytes
	//ebtText := fmt.Sprint(ebt)
	//profitText := fmt.Sprint(profit)
	//taxRateText := fmt.Sprint(taxRate)

	profitTextInfo := fmt.Sprint("EBT Tax: ", ebt, "\nProfit: ", profit, "\nTax Rate ", taxRate)
	fmt.Println("Writing to file....")
	os.WriteFile(profitInformationFile, []byte(profitTextInfo), 0644)
}
