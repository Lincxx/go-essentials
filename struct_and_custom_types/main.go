package main

import (
	"example.com/structs/user"
	"fmt"
)

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	//var appUser user
	//
	//appUser = newUser(userFirstName, userLastName, userBirthdate)

	var appUser2 *user
	appUser2 := &user.User{
		FirstName: userFirstName,
		LastName:  userLastName,
		Birthdate: userBirthdate,
	}

	if err != nil {
		fmt.Println(err)
		return
	}

	// ... do something awesome with that gathered data!

	//we don't have to pass an arg, as one is not required anymore. The user struct will be pass automatically to the
	//method
	appUser2.outputUserDetails()
	appUser2.clearUserName()
	appUser2.outputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
