package main

import "fmt"

func main() {
	var revenue float64
	var expense float64
	var tex float64

	fmt.Print("Enter Revenue Amount: ")
	fmt.Scan(&revenue)

	fmt.Print("Enter Expense Amount: ")
	fmt.Scan(&expense)

	fmt.Print("Enter Tex Amount: ")
	fmt.Scan(&tex)

	ebt := revenue - expense

	profit := (1 - tex/100) * ebt

	ratio := profit / ebt

	fmt.Println("Earnings Before Tex: ", ebt)
	fmt.Println("Profit: ", profit)
	fmt.Println("Ratio: ", ratio)

}
