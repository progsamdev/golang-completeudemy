package main

import (
	"fmt"
	"log"

	"calculator.samuel.com/filemanager"
	"calculator.samuel.com/prices"
)

func main() {
	taxRates := []float64{0.0, 0.07, 0.1, 0.15}
	doneChans := make([]chan bool, len(taxRates))
	errChans := make([]chan error, len(taxRates))
	for index, taxRate := range taxRates {
		fileManager := filemanager.NewFileManager("prices/prices.txt", fmt.Sprintf("prices/result-%v.json", taxRate))
		job := prices.NewTaxIncludedPriceJob(taxRate, fileManager)
		doneChans[index] = make(chan bool)
		errChans[index] = make(chan error)
		go job.Process(doneChans[index], errChans[index])
	}

	for index := range taxRates {
		select {
		case err := <-errChans[index]:
			if err != nil {
				fmt.Printf("Job %v failed\n", err)
				log.Fatal(err)
			}
		case <-doneChans[index]:
			fmt.Printf("Job %v done\n", index)
		}
	}

}
