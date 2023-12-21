package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	
	var investmentAmount float64
	var years float64
	expectedReturnRate := 5.5
	

	//fmt.Print("Investment Amount: ")
	outputText("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	//fmt.Print("Expected Return Rate: ")
	outputText("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	//fmt.Print("How many years: ")
	outputText("How many years: ")
	fmt.Scan(&years)

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)
	
	//futureValue := investmentAmount * math.Pow(1 + expectedReturnRate / 100, years)
	//futureRealValue := futureValue / math.Pow(1 + inflationRate/100, years)



	formattedFV := fmt.Sprintf("Future Value: %.2f\n", futureValue)
	formattedRFV := fmt.Sprintf("Future Value (adjusted for inflation): %.2f\n", futureRealValue)
	fmt.Print(formattedFV, formattedRFV)
	fmt.Println("Future Value: ", futureValue)

	// the backtick will use the format we have set. So if there are newline or tabs they will be shown
	// fmt.Printf(`Future Value: %.2f\n 
	// 			Future Value (adjusted for inflation): %.2f`, futureValue, futureRealValue)

	// fmt.Printf("Future Value (adjusted for inflation): %.2f", futureRealValue)
}

func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValues(investmentAmount float64 , expectedReturnRate float64,  years float64) (fv float64, rfv float64) {
	//we can return multiple with a comma at the end of the first return
	fv = investmentAmount * math.Pow(1 + expectedReturnRate / 100, years)
	rfv = fv / math.Pow(1 + inflationRate/100, years)
	return fv, rfv
	
	//Go will return the return variable we set in the function return
	//return
}