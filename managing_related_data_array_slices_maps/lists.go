package main

import (
	"fmt"
)

func main() {
	//Array
	//opt 1
	var productNames [4]string = [4]string{"A book"}
	//opt 2
	prices := [4]float64{10.99, 9.99, 45.99, 20.0}

	productNames[2] = "Cool story"
	fmt.Println(prices[2])
	fmt.Println(productNames)

	// Slices
	featuredPrices := prices[1:3]
	fmt.Println(featuredPrices)
}
