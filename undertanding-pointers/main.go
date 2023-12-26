package main

import "fmt"

func main() {

	//intro to pointers
	age := 32 //reg var

	//just making things clear
	agePointer := &age

	//we can also declare this in two lines of code
	//var agePointer *int
	//agePointer := &age

	fmt.Println("Age", age)

	//fmt.Println(getAdultYears())
	//or create a helper variable
	adultYears := getAdultYears(age)
	fmt.Println(adultYears)

}

func getAdultYears(age int) int {
	return age - 18
}
