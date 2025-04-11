package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func GetFloatFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)

	if err != nil {
		return 1000, errors.New("Failed to find balance file.")
	}

	valueS := string(data)
	value, err := strconv.ParseFloat(valueS, 64)

	if err != nil {
		return 1000, errors.New("Failed to parse stored balance value.")
	}

	return value, nil
}

func WriteFloatToFile(fileName string, value float64) {
	balanceText := fmt.Sprint(value)
	os.WriteFile(fileName, []byte(balanceText), 0644)
}
