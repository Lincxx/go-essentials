package main

import "fmt"

func main() {

	//intro to pointers
	age := 32 //reg var

	fmt.Println(age)
	
	//fmt.Println(getAdultYears())
	//or create a helper variable
	adultYears := getAdultYears(age)
	fmt.Println(adultYears)

}

func getAdultYears(age int) int{
	return age - 18
}