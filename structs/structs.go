package main

import (
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

func main() {

	firstNameInput := getUserData("Please enter your first name: ")
	lastNameInput := getUserData("Please enter your last name: ")
	birthdateInput := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUserEmpty User

	//
	var appUser = User{
		firstName: firstNameInput,
		lastName:  lastNameInput,
		birthdate: birthdateInput,
		createdAt: time.Now(),
	}

	//
	var appUser2 = User{
		firstNameInput,
		lastNameInput,
		birthdateInput,
		time.Now(),
	}

	// ... do something awesome with that gathered data!

	fmt.Println(appUser, appUser2, appUserEmpty)
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
