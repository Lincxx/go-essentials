package main

import "fmt"

func main() {

	//intro to pointers
	age := 32 //reg var

	//just making things clear - & creates the pointer
	//agePointer := &age

	//we can also declare this in two lines of code
	var agePointer *int
	agePointer = &age

	//prints out the address - might not be to useful here, maybe when we pass it to a function it will be
	//fmt.Println("Age", agePointer)

	//if we want the actual value we'll need to do something called dereferencing
	fmt.Println("Age", *agePointer)

	//fmt.Println(getAdultYears())
	//or create a helper variable
	//adultYears := getAdultYears(agePointer)
	//fmt.Println(adultYears)

	//editAdultYears(agePointer)
	fmt.Println(age)

}

// the name of the function states we are getting adult years (getAdultYears) and not editing them. Change the name
// to editAdultYears
func editAdultYears(age *int) {
	//return *age - 18
	//deference - then we are overriding that value in age
	*age = *age - 18
	//this could be a problem or unexpected behavior, that we wrote over a variable.
}
