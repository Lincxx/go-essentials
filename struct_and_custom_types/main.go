package main

import (
	"example.com/structs/user"
	"fmt"
)

type Str string

func (text Str) Log() {
	fmt.Println(text)
}

func main() {
	var name Str = "Max"
	name.Log()

	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	//var appUser user
	//
	//appUser = newUser(userFirstName, userLastName, userBirthdate)

	var appUser2 *user.User

	appUser2, err := user.New(userFirstName, userLastName, userBirthdate)

	if err != nil {
		fmt.Println(err)
		return
	}

	admin := user.NewAdmin("test@mail.com", "test123")
	admin.OutputUserDetails()
	admin.ClearUserName()

	// ... do something awesome with that gathered data!

	//we don't have to pass an arg, as one is not required anymore. The user struct will be pass automatically to the
	//method
	appUser2.OutputUserDetails()
	appUser2.ClearUserName()
	appUser2.OutputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
