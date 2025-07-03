package main

import (
	"example.com/price-calculator/fileManager"
	"example.com/price-calculator/prices"
	"fmt"
)

func main() {
	taxValue := []float64{0, 0.7, 0.1, 0.15}
	doneChans := make([]chan bool, len(taxValue))
	errorChans := make([]chan error, len(taxValue))

	for index, taxVal := range taxValue {
		doneChans[index] = make(chan bool)
		errorChans[index] = make(chan error)
		fm := fileManager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxVal*100))
		//cmd := cmdManager.New()
		//job := prices.NewTaxIncludedPriceJob(cmd, taxVal)
		job := prices.NewTaxIncludedPriceJob(fm, taxVal)
		//err := job.Process()

		// go routine setup
		go job.Process(doneChans[index], errorChans[index])

		//if err != nil {
		//	fmt.Println("Could not process job.")
		//	fmt.Println(err)
		//	return
		//}
	}

	// error handling for channels using select keyword

	for index := range taxValue {
		select {
		case err := <-errorChans[index]:
			if err != nil {
				fmt.Println(err)
			}
		case <-doneChans[index]:
			fmt.Println("Done!!!")
		}
	}

	// using the select keyword once the case value is executed it will move on not need to wait or care about other value
}
