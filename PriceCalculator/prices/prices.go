package prices

import (
	"example.com/price-calculator/conversion"
	"example.com/price-calculator/ioManager"
	"fmt"
)

type TaxIncludedPriceJob struct {
	IOManager        ioManager.IOManger `json:"-"`
	TaxRate          float64            `json:"tax_rate"`
	InputPrices      []float64          `json:"input_prices"`
	TaxIncludedPrice map[string]string  `json:"tax_included_price"`
}

func NewTaxIncludedPriceJob(iom ioManager.IOManger, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IOManager:   iom,
		TaxRate:     taxRate,
		InputPrices: []float64{10, 20, 30},
	}
}

func (job TaxIncludedPriceJob) Process(doneChan chan bool, errorChan chan error) {

	// simulate the error
	//errorChan <- errors.New("An Error Occurred !!!")

	//loading the data from file
	err := job.LoadData()

	if err != nil {
		errorChan <- err
		return
	}

	result := make(map[string]string)

	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	job.TaxIncludedPrice = result

	// write to file
	//return job.IOManager.WriteData(job)

	job.IOManager.WriteData(job)

	doneChan <- true
}

func (job *TaxIncludedPriceJob) LoadData() error {

	lines, err := job.IOManager.ReadLines()

	if err != nil {
		return err
	}

	prices, err := conversion.StringsToFloat(lines)

	if err != nil {
		return err
	}

	// parse the convert prices to struct
	job.InputPrices = prices

	return nil
}
