package main

import "fmt"

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	fmt.Print("What is your revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("What is your expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("What is your tax rate: ")
	fmt.Scan(&taxRate)

	ebt := revenue - expenses
	profit := ebt * (1-taxRate / 100)
	ratio := ebt / profit

	fmt.Println(ebt)
	fmt.Println(profit)
	fmt.Println(ratio)
}
