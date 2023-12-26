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
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	// ... do something awesome with that gathered data!

	fmt.Println(firstName, lastName, birthdate)

	//this can be error prone. Missing value or wrong order. Might prefer a value type that groups item together (Struct)
	outputUserDetails(firstName, lastName, birthdate)
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
