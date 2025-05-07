package main

import (
	"com.example/fileOps"
	"fmt"
)

const accountBalanceFile = "balance.txt"

func printBalance(balance float64) {
	balanceText := fmt.Sprint(balance)
	fmt.Println("Your account balance is ", balanceText)
}

func getAccountBalance(balance float64, fileName string) {
	fileOps.WriteValueFromFile(balance, fileName)
	printBalance(balance)
}

func main() {
	//account balance
	accountBalance, err := fileOps.GetValueFromFile(accountBalanceFile)

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
			getAccountBalance(accountBalance, accountBalanceFile)
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
			getAccountBalance(accountBalance, accountBalanceFile)

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
			getAccountBalance(accountBalance, accountBalanceFile)

		default:
			fmt.Println("GoodBye!")
			fmt.Println("Thanks for choosing our bank")
			return

		}
	}

}
