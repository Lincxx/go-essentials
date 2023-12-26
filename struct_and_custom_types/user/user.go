package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	FirstName string
	LastName  string
	Birthdate string
	createdAt time.Time
}

// attach to struct - the (u user) is called a receiver
func (u *user) OutputUserDetails() {

	//in the pointer section we were taught to use derefence  *u, to use this we need this (*u).firstName
	//this is using ths address and not the value of the address, this work because this is a shortcut provided by Go
	fmt.Println(u.firstName, u.lastName, u.birthdate)
}

// mutation method - we're only changing the copy (u user) doing this will change the struct (u *user)
func (u *user) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}

// Creation / Constructor Functions (Pattern) - NOT A FEATURE BUILT IN GO
//func newUser(firstName string, lastName string, birthDate string) user {
//	return user{
//		firstName: firstName,
//		lastName:  lastName,
//		birthdate: birthDate,
//		createdAt: time.Now(),
//	}
//}

// Creation / Constructor Functions (Pattern) - NOT A FEATURE BUILT IN GO
// we can return this as a pointer - what this does is to prevent a copy
func NewUser2(firstName string, lastName string, birthDate string) (*user, error) {
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("First name, Last name, birthdate required")
	}

	return &user{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthDate,
		createdAt: time.Now(),
	}, nil
}