package main

import (
	"fmt"

	"example.com/structs/user"
)

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser *user.User
	// Instance of the struct, you can omit values to implement zero-values
	// err wasn't declared, that's why :=
	appUser, err := user.New(userFirstName, userLastName, userBirthdate)
	if err != nil {
		fmt.Println(err)
		return
	}

	admin := user.NewAdmin("test@example.com", "test123", "Jordi")
	// Using anonymous embedding:
	// accessing the User's methods and struct.
	admin.OutputUserDetails()
	admin.ClearUserName()
	admin.OutputUserDetails()

	// Call the struct's method.
	appUser.OutputUserDetails()
	appUser.ClearUserName()
	appUser.OutputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	// to introduce a value by pressing enter
	fmt.Scanln(&value)
	return value
}
