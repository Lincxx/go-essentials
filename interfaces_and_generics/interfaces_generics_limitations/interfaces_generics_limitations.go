package liminterfaces_generics_limitationsitations

import "fmt"

func main() {

	result := add(1, 2)

	fmt.Println(result)

}

// generics
func add[T int | float64 | string](a, b T) T {
	return a + b

	// aInt, aIsInt := a.(int)
	// baInt, bIsInt := b.(int)
	// if aIsInt, bIsInt {
	// 	return aInt + bInt
	// }

	// aFloat, aIsFloat := a.(float64)
	// bFloat, bIsFloat := b.(float64)
	// if aIsFloat, bIsFloat {
	// 	return aFloat + bFloat
	// }
	// aString, aIsString := a.(string)
	// bString, bIsSrting := b.(string)
	// if aIsString, bIsString {
	// 	return aString + bString
	// }
}
