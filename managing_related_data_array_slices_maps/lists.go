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

	// Slices - are pretty much a window into the array. So it looks or modifies the org array.
	featuredPrices := prices[1:3]
	fmt.Println(featuredPrices)

	//omit the starting index
	featuredPrices2 := prices[:3]
	fmt.Println(featuredPrices2)

	featuredPrices3 := prices[1:]
	fmt.Println(featuredPrices3)

	//building slice off of an other slice
	hightlightedPrices := featuredPrices[:1]
	fmt.Println(hightlightedPrices)

	//Diving Deeper Into Slices
	featuredPrices4 := prices[1:]
	//this overrides the org array
	featuredPrices4[0] = 199.99
	hightlightedPrices1 := featuredPrices4[:1]
	fmt.Println(hightlightedPrices1)
	fmt.Println(prices)
	fmt.Println(len(prices))
	fmt.Println(len(featuredPrices4), cap(featuredPrices4))
	fmt.Println(len(hightlightedPrices1), cap(hightlightedPrices1))

}
