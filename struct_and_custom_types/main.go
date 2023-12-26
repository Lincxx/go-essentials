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

// Creation / Constructor Functions (Pattern) - NOT A FEATURE BUILT IN GO
func newUser(firstName string, lastName string, birthDate string) user {
	return user{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthDate,
		createdAt: time.Now(),
	}
}

// we can return this as a pointer - what this does is to prevent a copy
func newUser2(firstName string, lastName string, birthDate string) *user {
	return &user{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthDate,
		createdAt: time.Now(),
	}
}

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser user

	appUser = newUser(userFirstName, userLastName, userBirthdate)

	var appUser2 *user
	appUser2 = newUser2(userFirstName, userLastName, userBirthdate)

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
