package recursion

import "fmt"

func main() {
	//---------- Making Sense Of Recursion ------
	fact := factorial(5)

	fmt.Println(fact)
}

// factorial of 5 = 5 * 4 * 3 * 2 * 1  => 120
func factorial(number int) int {
	//---we need an exit condition
	//----recursion approach
	if number == 0 {
		return 1
	}
	return number * factorial(number-1)

	// result := 1

	// for i := 1; i <= number; i++ {
	// 	result = result * i
	// }

	// return result
}
