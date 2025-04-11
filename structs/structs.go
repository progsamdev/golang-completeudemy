package main

import (
	"fmt"
	"gocomplete-structs/user"
)

func main() {
	firstName := getUserData("Please enter your first name v2: ")
	lastName := getUserData("Please enter your last name v2: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY) v2: ")

	appUser, err := user.New(firstName, lastName, birthdate)
	if err != nil {
		fmt.Println(err)
		return
	}
	appUser.OutputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
