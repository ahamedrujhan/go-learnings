package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func getBalanceFormFile() (float64, error) {
	data, err := os.ReadFile(accountBalanceFile)
	// error handling if the balance.txt not found
	if err != nil {
		return 1000, errors.New("Balance.txt not found!")
	}
	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)
	// error handling if the float conversion found
	if err != nil {
		return 1000, errors.New("Failed to parse balance value from balance.txt")
	}
	return balance, nil
}

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)

	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}

func main() {
	//account balance
	accountBalance, err := getBalanceFormFile()

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("-----------")
		//break the application if the error happens
		// return
		// braek the application and gives more context about the application crash
		// panic("Can't Continue Sorry....")
	}

	var choice int

	//welcome message
	fmt.Println("Welcome To The Bank!")

	for {
		//showing the options
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check Your Balance")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Widthdraw Money")
		fmt.Println("4 . Exit")

		//get the choice
		fmt.Println("Enter Your Choice: ")
		fmt.Scan(&choice)

		//switch statements
		switch choice {
		case 1:
			fmt.Println("Your account balance is: ", accountBalance)
			writeBalanceToFile(accountBalance)
			continue
		case 2:
			fmt.Println("Enter Ammount to deposit")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Deposit amount should larger than 0")
				continue
			}

			accountBalance += depositAmount
			fmt.Println("Your account balance is ", accountBalance)
			writeBalanceToFile(accountBalance)

		case 3:
			fmt.Println("Enter amount to widthdraw money")
			var widthdrawAmount float64
			fmt.Scan(&widthdrawAmount)

			if widthdrawAmount <= 0 {
				fmt.Println("Withdraw amount should be greater than 0.")
				continue
			}

			if widthdrawAmount > accountBalance {
				fmt.Println("Can't widthdraw. widthdraw ammount is larger than account balance")
				continue
			}
			accountBalance -= widthdrawAmount
			fmt.Println("Your new account balance is ", accountBalance)
			writeBalanceToFile(accountBalance)

		default:
			fmt.Println("GoodBye!")
			fmt.Println("Thanks for choosing our bank")
			return

		}
	}

}
