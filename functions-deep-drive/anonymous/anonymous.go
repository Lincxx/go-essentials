package anonymous

import (
	"fmt"
)

func main() {
	numbers := []int{1, 2, 3}

	double := createTransformer(2) //remember these are returning a function
	triple := createTransformer(3) //remember these are returning a function

	//----------- Introducing Anonymous Functions ------
	transformed := transformNumbers(&numbers, func(number int) int {
		return number * 2
	})

	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbers, triple)

	fmt.Println(transformed)
	fmt.Println(doubled)
	fmt.Println(tripled)
}

func transformNumbers(numbers *[]int, transform func(int) int) []int {
	dNumbers := []int{}

	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}

// ----- factory function --- this is a Pattern.
func createTransformer(factor int) func(int) int {
	//----here is the Closure -
	return func(number int) int {
		return number * factor
	}
}
