package main

import (
	"fmt"
)

//---------Func types
type transformFn func(int) int

func main() {
	
	numbers := []int{1, 2, 3, 4, 5}

	//double every number
	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbers, triple)
	fmt.Println(doubled)
	fmt.Println(tripled)

}

//passing a function as a parameter
//so the transform parameter is a func that takes ints and return ints
func transformNumbers(numbers *[]int, transform transformFn) []int {
	dNumebrs := []int{}

	for _, val := range *numbers {
		dNumebrs = append(dNumebrs, transform(val))
	}

	return dNumebrs
}

//gen func that doubles
func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}