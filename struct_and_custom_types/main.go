package main

import (
	"fmt"
	"time"
)

// Structs are usually outside of a function
type user struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser user

	//create and instance of the struct and we can omit values, just those values will be set to null example appUser2
	appUser = user{
		firstName: userFirstName,
		lastName:  userLastName,
		birthdate: userBirthdate,
		createdAt: time.Now(),
	}

	// last name and birthdate will be set to their null values
	//var appUser2 user
	//appUser2 = user{
	//	firstName: userFirstName,
	//	createdAt: time.Now(),
	//}

	// ... do something awesome with that gathered data!

	fmt.Println(userFirstName, userLastName, userBirthdate)
	fmt.Println(appUser)

	//this can be error prone. Missing value or wrong order. Might prefer a value type that groups item together (Struct)
	//outputUserDetails(firstName, lastName, birthdate)
}

func outputUserDetails(firstName, lastName, birthdate string) {
	fmt.Println(firstName, lastName, birthdate)
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
