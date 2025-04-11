package main

import (
	"fmt"

	"calculator.samuel.com/filemanager"
	"calculator.samuel.com/prices"
)

func main() {
	taxRates := []float64{0.0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxRates {
		fileManager := filemanager.NewFileManager("prices/prices.txt", fmt.Sprintf("prices/result-%v.json", taxRate))
		job := prices.NewTaxIncludedPriceJob(taxRate, fileManager)
		job.Process()
	}
}
