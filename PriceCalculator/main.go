package main

import "fmt"

func main() {
	prices := []float64{10, 20, 30}
	taxValue := []float64{0, 0.7, 0.1, 0.15}

	result := make(map[float64][]float64)

	for _, taxVal := range taxValue {
		taxedAmount := make([]float64, len(prices))
		for priceIndex, price := range prices {
			taxedAmount[priceIndex] = price * (1 + taxVal)
		}
		result[taxVal] = taxedAmount
	}
	fmt.Println(result)

}
