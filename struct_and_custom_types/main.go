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

// attach to struct - the (u user) is called a receiver
func (u *user) outputUserDetails() {

	//in the pointer section we were taught to use derefence  *u, to use this we need this (*u).firstName
	//this is using ths address and not the value of the address, this work because this is a shortcut provided by Go
	fmt.Println(u.firstName, u.lastName, u.birthdate)
}

// mutation method - we're only changing the copy (u user) doing this will change the struct (u *user)
func (u *user) clearUserName() {
	u.firstName = ""
	u.lastName = ""
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

	// ... do something awesome with that gathered data!

	//we don't have to pass an arg, as one is not required anymore. The user struct will be pass automatically to the
	//method
	appUser.outputUserDetails()
	appUser.clearUserName()
	appUser.outputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
