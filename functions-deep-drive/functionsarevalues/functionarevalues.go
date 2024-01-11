package functionsarevalues

import (
	"fmt"
)

// ---------Func types
type transformFn func(int) int

func main() {

	numbers := []int{1, 2, 3, 4, 5}
	moreNumbers := []int{1, 5, 1, 12}

	//double every number
	//doubled := transformNumbers(&numbers, double)
	//tripled := transformNumbers(&numbers, triple)
	// fmt.Println(doubled)
	// fmt.Println(tripled)

	transformFn1 := getTransformerFunction(&numbers)
	transformFn2 := getTransformerFunction(&moreNumbers)

	transformedNumbers := transformNumbers(&numbers, transformFn1)
	moreTransformedNumbers := transformNumbers(&moreNumbers, transformFn2)

	fmt.Println(transformedNumbers)
	fmt.Println(moreTransformedNumbers)

}

// passing a function as a parameter
// so the transform parameter is a func that takes ints and return ints
func transformNumbers(numbers *[]int, transform transformFn) []int {
	dNumebrs := []int{}

	for _, val := range *numbers {
		dNumebrs = append(dNumebrs, transform(val))
	}

	return dNumebrs
}

// -------Returning Functions As Values
func getTransformerFunction(numbers *[]int) transformFn {
	if (*numbers)[0] == 1 {
		return double
	} else {
		return triple
	}

}

// gen func that doubles
func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}
