package main

import "fmt"

func getUserData(printText string) string {
	fmt.Println(printText)
	var inputText string
	fmt.Scan(&inputText)
	return inputText
}

func main() {
	firstName := getUserData("please enter your first name")
	lastName := getUserData("please enter your last name")
	birthDate := getUserData("please enter your birth of date")

	// do sth...

	fmt.Println(firstName, lastName, birthDate)
}
