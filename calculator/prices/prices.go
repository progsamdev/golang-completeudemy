package prices

import (
	"fmt"
	"log"

	"calculator.samuel.com/conversion"
	"calculator.samuel.com/filemanager"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64                  `json:"TaxRate"`
	Prices            []float64                `json:"Prices"`
	TaxIncludedPrices map[string]string        `json:"TaxIncludedPrices"`
	FileManager       *filemanager.FileManager `json:"-"`
}

func NewTaxIncludedPriceJob(taxRate float64, fileManager *filemanager.FileManager) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		TaxRate:     taxRate,
		FileManager: fileManager,
	}
}

func (job *TaxIncludedPriceJob) Process() {

	job.LoadPrices()

	result := make(map[string]string)

	for _, price := range job.Prices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	job.TaxIncludedPrices = result
	err := job.FileManager.WriteJSON(job)
	if err != nil {
		log.Fatal(err)
	}
}

func (job *TaxIncludedPriceJob) LoadPrices() {

	lines, err := job.FileManager.ReadFile()

	if err != nil {
		log.Fatal(err)
	}

	prices, err := conversion.StringsToFloats(lines)

	if err != nil {
		log.Fatal(err)
	}

	job.Prices = prices

}
