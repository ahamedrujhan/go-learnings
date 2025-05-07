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

	err := os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)

	if err != nil {
		panic("Failed to write balance to file")
	}
}

func printBalance(balance float64) {
	balanceText := fmt.Sprint(balance)
	fmt.Println("Your account balance is ", balanceText)
}

func getAccountBalance(balance float64) {
	writeBalanceToFile(balance)
	printBalance(balance)
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
		presentOptions()
		//get the choice
		fmt.Println("Enter Your Choice: ")
		fmt.Scan(&choice)

		//switch statements
		switch choice {
		case 1:
			getAccountBalance(accountBalance)
			//continue
		case 2:
			fmt.Println("Enter Ammount to deposit")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Deposit amount should larger than 0")
				continue
			}

			accountBalance += depositAmount
			getAccountBalance(accountBalance)

		case 3:
			fmt.Println("Enter amount to withdraw money")
			var widthdrawAmount float64
			fmt.Scan(&widthdrawAmount)

			if widthdrawAmount <= 0 {
				fmt.Println("Withdraw amount should be greater than 0.")
				continue
			}

			if widthdrawAmount > accountBalance {
				fmt.Println("Can't withdraw. withdraw amount is larger than account balance")
				continue
			}
			accountBalance -= widthdrawAmount
			getAccountBalance(accountBalance)

		default:
			fmt.Println("GoodBye!")
			fmt.Println("Thanks for choosing our bank")
			return

		}
	}

}
