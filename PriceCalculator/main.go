package main

import (
	"example.com/price-calculator/fileManager"
	"example.com/price-calculator/prices"
	"fmt"
)

func main() {
	taxValue := []float64{0, 0.7, 0.1, 0.15}

	for _, taxVal := range taxValue {
		fm := fileManager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxVal*100))
		//cmd := cmdManager.New()
		//job := prices.NewTaxIncludedPriceJob(cmd, taxVal)
		job := prices.NewTaxIncludedPriceJob(fm, taxVal)
		err := job.Process()

		if err != nil {
			fmt.Println("Could not process job.")
			fmt.Println(err)
			return
		}
	}
}
