package main

import (
	"example.com/price-calculator/prices"
)

func main() {
	taxValue := []float64{0, 0.7, 0.1, 0.15}

	for _, taxVal := range taxValue {
		job := prices.NewTaxIncludedPriceJob(taxVal)
		job.Process()
	}
}
