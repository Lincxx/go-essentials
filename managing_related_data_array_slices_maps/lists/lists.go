package lists

import "fmt"

// func main() {
//-------------------Array--------------
// 	//opt 1
// 	var productNames [4]string = [4]string{"A book"}
// 	//opt 2
// 	prices := [4]float64{10.99, 9.99, 45.99, 20.0}

// 	productNames[2] = "Cool story"
// 	fmt.Println(prices[2])
// 	fmt.Println(productNames)

//----------------Slices - are pretty much a window into the array. So it looks or modifies the org array.---------
// 	featuredPrices := prices[1:3]
// 	fmt.Println(featuredPrices)

// 	//omit the starting index
// 	featuredPrices2 := prices[:3]
// 	fmt.Println(featuredPrices2)

// 	featuredPrices3 := prices[1:]
// 	fmt.Println(featuredPrices3)

// 	//building slice off of an other slice
// 	hightlightedPrices := featuredPrices[:1]
// 	fmt.Println(hightlightedPrices)

//------------------Diving Deeper Into Slices------------------
// 	featuredPrices4 := prices[1:]
// 	//this overrides the org array
// 	featuredPrices4[0] = 199.99
// 	hightlightedPrices1 := featuredPrices4[:1]
// 	fmt.Println(hightlightedPrices1)
// 	fmt.Println(prices)
// 	fmt.Println(len(prices))
// 	fmt.Println(len(featuredPrices4), cap(featuredPrices4))
// 	fmt.Println(len(hightlightedPrices1), cap(hightlightedPrices1))

// }

// func main() {
// 	// more common to use slices
//--------------------Building Dynamic Lists With Slices--------------------
// 	prices := []float64{10.99, 8.99}
// 	fmt.Println(prices[0:1])
// 	prices[1] = 9.99
// 	//can not add to the slice this way
// 	//prices[2] = 11.99

// 	//this does not add to the slice, but rather creates a new slice or we can reassign the var
// 	prices = append(prices, 5.99)
// 	prices = prices[1:]

// 	fmt.Println(prices)
// }

//-------------------Unpacking list values----------------
func main() {
	prices := []float64{10.99, 8.99}
	fmt.Println(prices[0:1])
	prices[1] = 9.99

	//can not add to the slice this way
	//prices[2] = 11.99
	//this does not add to the slice, but rather creates a new slice or we can reassign the var
	prices = append(prices, 5.99)
	prices = prices[1:]
	fmt.Println(prices)

	discountPrices := []float64{101.99, 80.99, 20.59}
	//merge discountPrice into prices by unpacking the list
	prices = append(prices, discountPrices...)
	fmt.Println(prices)

}
