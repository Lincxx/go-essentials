package main

import "fmt"

func main() {

	//--------- Variadic Functions
	//--------- Splitting Slices into Parameter Value

	numbers := []int{1, 10, 15}
	sum := sumup(numbers...)
	fmt.Println(sum)

	sum2 := sumup(1, 10, 15, -5)
	fmt.Println(sum2)
}

// take a list of numbers sum them up
// ---Variadic Functions
func sumup(numbers ...int) int {
	sum := 0

	for _, value := range numbers {
		sum += value
	}

	return sum
}

//--- mixture variadic
func sumup2(startingVal int, numbers ...int) int {
	sum := 0

	for _, value := range numbers {
		sum += value
	}

	return sum
}
