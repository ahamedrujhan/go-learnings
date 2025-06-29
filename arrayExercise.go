package main

import "fmt"

func main() {

	type Product struct {
		title string
		id    int
		price float64
	}

	hobbies := [3]string{"Reading", "Cricket", "Watching Movies"}

	// 1
	fmt.Println(hobbies)

	// 2
	fmt.Println(hobbies[:1])
	newHobbies := hobbies[1:3]
	fmt.Println(newHobbies)

	//6
	stocks := []Product{
		{id: 1, price: 150, title: "chocolate"},
		{id: 2, price: 100, title: "movies"},
	}

	newStocks := Product{
		id: 5, price: 20, title: "Pen",
	}
	stocks = append(stocks, newStocks)
	fmt.Println(stocks)
}
