package main

import (
	"fmt"
	"math"
)

func main() {
	// since we want to use a float64, we can't use the alt declaration 
	//type comes last
	//var investmentAmount, years float64 = 1000, 10

	//we can use .0 to let go know that we want float values (it will infer this)
	// investmentAmount, years := 1000.0, 10.0
	// expectedReturnRate := 5.5

	//var years float64 = 10
	// - we can go one step further and declare this on one line - this can make it harder to read
	investmentAmount, years, expectedReturnRate := 1000.0, 10.0, 5.5

	futureValue := investmentAmount * math.Pow(1 + expectedReturnRate / 100, years)

	fmt.Println(futureValue)
}