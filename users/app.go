package main

import (
	"fmt"
	"time"
)

func getUserData(printText string) string {
	fmt.Println(printText)
	var inputText string
	fmt.Scan(&inputText)
	return inputText
}

type user struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

func main() {
	firstName := getUserData("please enter your first name")
	lastName := getUserData("please enter your last name")
	birthDate := getUserData("please enter your birth of date")

	var appUser user

	// struct literal notation
	// we can omit any value if we want
	//appUser = user{
	//	firstName: firstName,
	//	lastName:  lastName,
	//	birthDate: birthDate,
	//	createdAt: time.Now(),
	//}

	// in this method order is important
	appUser = user{
		firstName,
		lastName,
		birthDate,
		time.Now(),
	}
	//null value struct
	//appUser = user{}

	// do sth...

	fmt.Println(firstName, lastName, birthDate)
}
