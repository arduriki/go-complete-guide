package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

type Admin struct {
	email string
	password string
	// add User
	User
}

// relate a method with the struct, starting with a receiver argument
func (u User) OutputUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthdate)
}

// change the values with a pointer
func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}

func NewAdmin(email, password, firstName string) Admin {
	return Admin{
		email: email,
		password: password,
		User: User{
			firstName: firstName,
			lastName: "ADMIN",
			birthdate: "---",
			createdAt: time.Now(),
		},
	}
}

// constructor function
func New(firstName, lastName, birthDate string) (*User, error) {
	// validate data
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("first name, last name and birthdate are required")
	}

	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthDate,
		createdAt: time.Now(),
	}, nil
}
